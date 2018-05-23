package got

// PluginUserInput type
type PluginUserInput struct {
	ComponentBase
	SuccessBlockID string
}

// Execute func
func (c *PluginUserInput) Execute(ctx *BotContext) error {
	ctx.state().state = BotUserStateStateText
	ctx.state().payload = &BotPayload{BlockID: c.SuccessBlockID}
	return nil
}
