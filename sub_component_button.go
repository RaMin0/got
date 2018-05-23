package got

import "github.com/ramin0/messenger"

const (
	// SubComponentButtonTypeWebURL const
	SubComponentButtonTypeWebURL = "web_url"
	// SubComponentButtonTypePostback const
	SubComponentButtonTypePostback = "postback"

	// SubComponentButtonWebviewHeightRatioCompact const
	SubComponentButtonWebviewHeightRatioCompact = "compact"
	// SubComponentButtonWebviewHeightRatioTall const
	SubComponentButtonWebviewHeightRatioTall = "tall"
	// SubComponentButtonWebviewHeightRatioFull const
	SubComponentButtonWebviewHeightRatioFull = "full"

	// SubComponentButtonWebviewShareButtonHide const
	SubComponentButtonWebviewShareButtonHide = "hide"

	// SubComponentButtonMessengerExtensionsTrue const
	SubComponentButtonMessengerExtensionsTrue = "true"
)

// SubComponentButton type
type SubComponentButton struct {
	Title               string
	Type                string
	Payload             *BotPayload
	URL                 string
	WebviewHeightRatio  string
	WebviewShareButton  string
	MessengerExtensions string
}

// Execute func
func (c *SubComponentButton) Execute(ctx *BotContext) *messenger.ElmButton {
	switch c.Type {
	case SubComponentButtonTypePostback:
		return c.executePostback(ctx)
	case SubComponentButtonTypeWebURL:
		return c.executeWebURL(ctx)
	}
	return nil
}

func (c *SubComponentButton) executePostback(ctx *BotContext) *messenger.ElmButton {
	payload := BotPayload{
		BlockID: c.Payload.BlockID,
		Content: ctx.InterpolateMap(c.Payload.Content),
	}

	return &messenger.ElmButton{
		Type:    c.Type,
		Title:   ctx.Interpolate(c.Title),
		Payload: payload.Marshal(),
	}
}

func (c *SubComponentButton) executeWebURL(ctx *BotContext) *messenger.ElmButton {
	return &messenger.ElmButton{
		Type:                c.Type,
		Title:               ctx.Interpolate(c.Title),
		URL:                 ctx.Interpolate(c.URL),
		WebviewHeightRatio:  ctx.Interpolate(c.WebviewHeightRatio),
		WebviewShareButton:  ctx.Interpolate(c.WebviewShareButton),
		MessengerExtensions: ctx.Interpolate(c.MessengerExtensions),
	}
}
