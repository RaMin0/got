package got

import (
	"strings"

	"github.com/ramin0/messenger"
)

const (
	// ComponentTextTextMaxLength const
	ComponentTextTextMaxLength = 640
)

// ComponentText type
type ComponentText struct {
	ComponentBase
	Text    string
	Buttons []*SubComponentButton
}

var x = 0

// Execute func
func (c *ComponentText) Execute(ctx *BotContext) error {
	text := ctx.Interpolate(c.Text)

	chars := strings.Split(text, "")

	textParts := []string{""}
	for len(chars) > 0 {
		lastTextPart := textParts[len(textParts)-1]

		if len(lastTextPart)+len(chars[0]) <= ComponentTextTextMaxLength {
			lastTextPart += chars[0]
			chars = chars[1:]

			if len(lastTextPart) == ComponentTextTextMaxLength && len(chars) > 0 {
				lastChars := strings.Split(lastTextPart[len(lastTextPart)-3:], "")
				lastTextPart = lastTextPart[:len(lastTextPart)-3] + "..."
				chars = append(append(strings.Split("...", ""), lastChars...), chars...)
			}

			textParts[len(textParts)-1] = lastTextPart
		} else {
			textParts = append(textParts, "")
		}
	}

	for i, textPart := range textParts {
		if i == len(textParts)-1 && len(c.Buttons) > 0 {
			buttons := []*messenger.ElmButton{}

			for _, b := range c.Buttons {
				buttons = append(buttons, b.Execute(ctx))
			}

			return c.client().SendButtons(ctx.userID,
				textPart,
				buttons)
		}

		if err := c.client().SendText(ctx.userID,
			textPart); err != nil {
			return err
		}
	}

	return nil
}
