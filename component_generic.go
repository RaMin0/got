package got

import "github.com/ramin0/messenger"

// ComponentGeneric type
type ComponentGeneric struct {
	ComponentBase
	Dynamic  bool
	Prefix   string
	Actions  []*ComponentGenericAction
	Elements []*SubComponentElement
}

// ComponentGenericAction type
type ComponentGenericAction struct {
	Name    string
	BlockID string
}

// Execute func
func (c *ComponentGeneric) Execute(ctx *BotContext) error {
	elements := c.Elements
	if c.Dynamic {
		elements = c.buildElements(ctx)
	}

	elms := []*messenger.ElmElement{}
	for _, e := range elements {
		elms = append(elms, e.Execute(ctx))
	}

	return c.client().SendGenericTemplate(ctx.userID,
		elms)
}

func (c *ComponentGeneric) buildElements(ctx *BotContext) []*SubComponentElement {
	var elements []*SubComponentElement
	elementsData := ctx.Interpolations()[c.Prefix].([]map[string]interface{})
	for _, data := range elementsData {
		e := &SubComponentElement{}
		elements = append(elements, e)

		if title, ok := data["title"].(string); ok {
			e.Title = title
		}
		if subtitle, ok := data["subtitle"].(string); ok {
			e.Subtitle = subtitle
		}
		if imageURL, ok := data["image_url"].(string); ok {
			e.ImageURL = imageURL
		}
		if buttons, ok := data["buttons"].([]map[string]interface{}); ok {
			for _, b := range buttons {
				button := &SubComponentButton{}
				e.Buttons = append(e.Buttons, button)

				if t, ok := b["type"].(string); ok {
					button.Type = t

					switch t {
					case SubComponentButtonTypeWebURL:
						button.URL = b["url"].(string)
					case SubComponentButtonTypePostback:
						if actionData, ok := b["action"].(map[string]interface{}); ok {
							var action *ComponentGenericAction
							for _, a := range c.Actions {
								if actionData["name"] == a.Name {
									action = a
									break
								}
							}
							if action == nil {
								continue
							}

							button.Payload = &BotPayload{
								BlockID: action.BlockID,
								Content: actionData["params"].(map[string]interface{}),
							}
						}
					}
				}

				if title, ok := b["title"].(string); ok {
					button.Title = title
				}
			}
		}
	}
	return elements
}
