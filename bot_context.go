package got

import (
	"fmt"
	"regexp"
	"strings"
)

// BotContext type
type BotContext struct {
	bot     *BotData
	userID  string
	payload map[string]interface{}
}

// Session func
func (ctx *BotContext) Session() BotUserSession {
	s, ok := ctx.bot.userSessions[ctx.userID]
	if !ok {
		s = &BotUserSession{}
		ctx.bot.userSessions[ctx.userID] = s
	}
	return *s
}

func (ctx *BotContext) state() *botUserState {
	s, ok := ctx.bot.userStates[ctx.userID]
	if !ok {
		s = &botUserState{}
		ctx.bot.userStates[ctx.userID] = s
	}
	return s
}

// Callback func
func (ctx *BotContext) Callback(blockID string) *BotCallback {
	callback := &BotCallback{
		id:      uniqueID(),
		bot:     ctx.bot,
		blockID: blockID,
		userID:  ctx.userID,
	}
	ctx.bot.callbacks[callback.id] = callback
	return callback
}

// Interpolations func
func (ctx *BotContext) Interpolations() map[string]interface{} {
	_, ok := ctx.payload["interpolations"]
	if !ok {
		i := map[string]interface{}{}
		ctx.payload["interpolations"] = i

		for k, v := range ctx.Postback() {
			i[fmt.Sprintf("postback:%s", k)] = v
		}
	}
	return ctx.payload["interpolations"].(map[string]interface{})
}

// Interpolate func
func (ctx *BotContext) Interpolate(text string) string {
	is := ctx.Interpolations()
	if len(is) == 0 {
		return text
	}

	interpolate := func(s string, is map[string]interface{}) string {
		return regexp.MustCompile("(%\\{[^\\{\\}]+\\})").
			ReplaceAllStringFunc(s, func(match string) string {
				key := strings.TrimSuffix(match[2:], "}")
				value, ok := is[key]
				if !ok {
					return match
				}
				return fmt.Sprintf("%+v", value)
			})
	}

	prevText := ""
	for prevText != text {
		prevText = text
		text = interpolate(text, is)
	}

	return text
}

// InterpolateMap func
func (ctx *BotContext) InterpolateMap(dict map[string]interface{}) map[string]interface{} {
	interpolatedDict := make(map[string]interface{}, len(dict))
	for k, v := range dict {
		if vd, ok := v.(map[string]interface{}); ok {
			interpolatedDict[k] = ctx.InterpolateMap(vd)
			continue
		}

		interpolatedDict[k] = ctx.Interpolate(fmt.Sprintf("%v", v))
	}
	return interpolatedDict
}

// Postback func
func (ctx *BotContext) Postback() map[string]interface{} {
	_, ok := ctx.payload["postback"]
	if !ok {
		p := map[string]interface{}{}
		ctx.payload["postback"] = p
	}
	return ctx.payload["postback"].(map[string]interface{})
}
