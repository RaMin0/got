package got

// PluginGoToBlock type
type PluginGoToBlock struct {
	ComponentBase
	TargetBlockID string
}

// Execute func
func (c *PluginGoToBlock) Execute(ctx *BotContext) error {
	if block, ok := c.bot().blocks[c.TargetBlockID]; ok {
		if err := block.execute(ctx); err != nil {
			return err
		}
	}
	return ErrBreak
}
