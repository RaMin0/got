package got

// BotUserSession type
type BotUserSession map[string]interface{}

// Clear func
func (s BotUserSession) Clear() {
	for k := range s {
		delete(s, k)
	}
}
