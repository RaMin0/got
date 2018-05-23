package got

const (
	// BotUserStateStateText const
	BotUserStateStateText = "text"
)

type botUserState struct {
	state   string
	payload *BotPayload
}
