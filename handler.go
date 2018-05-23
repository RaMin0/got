package got

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"github.com/ramin0/messenger"
)

// Handler func
func (bot BotData) Handler() http.Handler {
	baseURL, err := url.Parse(bot.BaseURL)
	if err != nil {
		panic(fmt.Errorf("%q is not a valid value for BaseURL: %v", bot.BaseURL, err))
	}

	mux := http.NewServeMux()
	mux.Handle(fmt.Sprintf("%s/", baseURL.Path), bot.Client.
		HandleMessage(bot.messageHandler).
		HandlePostback(bot.postbackHandler).
		Handler())
	mux.HandleFunc(fmt.Sprintf("%s/callbacks/", baseURL.Path),
		bot.callbackHandler)
	return mux
}

func (bot BotData) messageHandler(ms *messenger.Messenger, m *messenger.Message) {
	if p, ok := bot.IsPending(m.Sender.ID, BotUserStateStateText); ok {
		postback := map[string]interface{}{"value": m.Text}
		payload := map[string]interface{}{"postback": postback}
		bot.StartWithBlock(m.Sender.ID, p.BlockID, payload)
		return
	}

	bot.Start(m.Sender.ID)
}

func (bot BotData) postbackHandler(ms *messenger.Messenger, p *messenger.Postback) {
	botPayload := &BotPayload{}
	botPayload.Unmarshal(p.Payload)

	payload := map[string]interface{}{"postback": botPayload.Content}
	bot.StartWithBlock(p.Sender.ID, botPayload.BlockID, payload)
}

func (bot BotData) callbackHandler(w http.ResponseWriter, r *http.Request) {
	callbackID := regexp.MustCompile("([^\\/]+)/?$").FindStringSubmatch(r.URL.Path)[1]

	postback := map[string]interface{}{}
	r.ParseForm()
	for k, v := range r.Form {
		postback[k[9:len(k)-1]] = v[0]
	}

	payload := map[string]interface{}{"postback": postback}
	bot.StartWithCallback(callbackID, payload)
}
