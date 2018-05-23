package got

import (
	"fmt"
)

// BotCallback type
type BotCallback struct {
	bot     *BotData
	id      string
	blockID string
	userID  string
}

// URL func
func (c *BotCallback) URL() string {
	return fmt.Sprintf("%s/callbacks/%s", c.bot.BaseURL, c.id)
}
