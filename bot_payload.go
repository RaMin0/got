package got

import (
	"encoding/base64"
	"encoding/json"
)

// BotPayload type
type BotPayload struct {
	BlockID string                 `json:"block_id"`
	Content map[string]interface{} `json:"content"`
}

// Marshal func
func (s *BotPayload) Marshal() string {
	b, _ := json.Marshal(s)
	return base64.URLEncoding.EncodeToString(b)
}

// Unmarshal func
func (s *BotPayload) Unmarshal(str string) {
	b, _ := base64.URLEncoding.DecodeString(str)
	json.Unmarshal(b, s)
}
