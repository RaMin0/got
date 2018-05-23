package got

import (
	"fmt"
)

const (
	// PluginSessionModeRead const
	PluginSessionModeRead = "read"

	// PluginSessionModeWrite const
	PluginSessionModeWrite = "write"

	// PluginSessionModeClear const
	PluginSessionModeClear = "clear"
)

// PluginSession type
type PluginSession struct {
	ComponentBase
	Mode string
	Src  string
	Dst  string
}

// Execute func
func (c *PluginSession) Execute(ctx *BotContext) error {
	switch c.Mode {
	case PluginSessionModeRead:
		dst := c.Dst
		if dst == "" {
			dst = c.Src
		}
		src := ctx.Session()[c.Src]
		if src == nil {
			src = ""
		}
		ctx.Interpolations()[fmt.Sprintf("var:%s", dst)] = fmt.Sprintf("%+v", src)
	case PluginSessionModeWrite:
		ctx.Session()[c.Dst] = ctx.Interpolate(c.Src)
	case PluginSessionModeClear:
		ctx.Session().Clear()
	}

	return nil
}
