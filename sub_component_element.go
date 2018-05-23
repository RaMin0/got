package got

import "github.com/ramin0/messenger"

// SubComponentElement type
type SubComponentElement struct {
	Title    string
	Subtitle string
	ImageURL string
	Buttons  []*SubComponentButton
}

// Execute func
func (c *SubComponentElement) Execute(ctx *BotContext) *messenger.ElmElement {
	buttons := []*messenger.ElmButton{}
	for _, b := range c.Buttons {
		buttons = append(buttons, b.Execute(ctx))
	}

	elm := &messenger.ElmElement{
		Title:    ctx.Interpolate(c.Title),
		Subtitle: ctx.Interpolate(c.Subtitle),
		ImageURL: ctx.Interpolate(c.ImageURL),
	}
	elm.Buttons = buttons
	return elm
}
