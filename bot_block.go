package got

// BotBlock type
type BotBlock struct {
	bot        *BotData
	components []BotComponent
}

// Component func
func (b *BotBlock) Component(component BotComponent) BotComponent {
	component.setBlock(b)
	b.components = append(b.components, component)
	return component
}

func (b *BotBlock) execute(ctx *BotContext) error {
	for _, c := range b.components {
		if err := c.Execute(ctx); err != nil {
			if err == ErrBreak {
				break
			}
			ctx.Interpolations()["postback:error"] = err.Error()
			return b.Component(&PluginGoToBlock{TargetBlockID: "_error"}).Execute(ctx)
		}
	}
	return nil
}

func (b *BotBlock) executeLater(ctx *BotContext) {
	go b.execute(ctx)
}
