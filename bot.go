package got

import (
	"fmt"

	"github.com/ramin0/messenger"
)

const (
	// StartBlock const
	StartBlock = "_start"
	// ErrorBlock const
	ErrorBlock = "_error"
	// FallbackBlock const
	FallbackBlock = "_fallback"
)

var (
	// ErrBreak var
	ErrBreak = fmt.Errorf("break")
)

// BotData type
type BotData struct {
	BaseURL      string
	Meta         map[string]interface{}
	Page         *BotPage
	Client       botClient
	contexts     map[string]*BotContext
	callbacks    map[string]*BotCallback
	userSessions map[string]*BotUserSession
	userStates   map[string]*botUserState
	blocks       map[string]*BotBlock
}

// New func
func New(bot *BotData, init func(*BotData)) *BotData {
	bot.Client = botClient{messenger.New(messenger.Options{
		AccessToken: bot.Page.AccessToken,
		VerifyToken: bot.Page.VerifyToken,
	})}
	bot.contexts = map[string]*BotContext{}
	bot.callbacks = map[string]*BotCallback{}
	bot.userSessions = map[string]*BotUserSession{}
	bot.userStates = map[string]*botUserState{}
	bot.blocks = map[string]*BotBlock{}

	init(bot)

	return bot
}

func (bot *BotData) context(userID string) *BotContext {
	c, ok := bot.contexts[userID]
	if !ok {
		c = &BotContext{
			bot:     bot,
			userID:  userID,
			payload: map[string]interface{}{},
		}
		bot.contexts[userID] = c
	}
	return c
}

// IsPending func
func (bot *BotData) IsPending(userID, state string) (*BotPayload, bool) {
	s := bot.context(userID).state()
	if s.state == state {
		delete(bot.userStates, userID)
		return s.payload, true
	}
	return nil, false
}

// Block func
func (bot *BotData) Block(name string, init func(*BotBlock)) {
	block := &BotBlock{
		bot: bot,
	}

	init(block)

	bot.blocks[name] = block
}

// StartBlock func
func (bot BotData) StartBlock(init func(*BotBlock)) {
	bot.Block(StartBlock, init)
}

// ErrorBlock func
func (bot BotData) ErrorBlock(init func(*BotBlock)) {
	bot.Block(ErrorBlock, init)
}

// FallbackBlock func
func (bot BotData) FallbackBlock(init func(*BotBlock)) {
	bot.Block(FallbackBlock, init)
}

// Start func
func (bot *BotData) Start(userID string) {
	bot.StartWithBlock(userID, "_start")
}

// StartWithCallback func
func (bot *BotData) StartWithCallback(callbackID string, payload ...map[string]interface{}) {
	if callback, ok := bot.callbacks[callbackID]; ok {
		bot.StartWithBlock(callback.userID, callback.blockID, payload...)
	}
}

// StartWithBlock func
func (bot *BotData) StartWithBlock(userID, blockID string, payload ...map[string]interface{}) {
	ctx := bot.context(userID)
	if len(payload) > 0 {
		ctx.payload = payload[0]
	}

	if block, ok := bot.blocks[blockID]; ok {
		block.executeLater(ctx)
	}
}
