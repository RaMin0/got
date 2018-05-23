package got

// BotComponent type
type BotComponent interface {
	setBlock(*BotBlock)
	Execute(*BotContext) error
}

// ComponentBase type
type ComponentBase struct {
	block *BotBlock
}

func (c *ComponentBase) bot() *BotData {
	return c.block.bot
}

func (c *ComponentBase) client() botClient {
	return c.bot().Client
}

func (c *ComponentBase) setBlock(b *BotBlock) {
	c.block = b
}
