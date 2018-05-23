package got

// PluginSwitch type
type PluginSwitch struct {
	ComponentBase
	Input          string
	Cases          map[string]string
	DefaultBlockID string
}

// Execute func
func (c *PluginSwitch) Execute(ctx *BotContext) error {
	input := ctx.Interpolate(c.Input)
	targetBlockID := ctx.Interpolate(c.DefaultBlockID)

	if blockID, ok := c.Cases[input]; ok {
		targetBlockID = ctx.Interpolate(blockID)
	}

	if targetBlockID == "" {
		return nil
	}

	return c.block.Component(&PluginGoToBlock{
		TargetBlockID: targetBlockID,
	}).Execute(ctx)
}
