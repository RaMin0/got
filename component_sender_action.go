package got

const (
	// ComponentSenderActionActionTypingOn const
	ComponentSenderActionActionTypingOn = "typing_on"
)

// ComponentSenderAction type
type ComponentSenderAction struct {
	ComponentBase
	Action string
}

// Execute func
func (c *ComponentSenderAction) Execute(ctx *BotContext) error {
	return c.client().SendSenderAction(ctx.userID,
		c.Action)
}
