# Got

## Guide
[![Facebook Bots](https://img.youtube.com/vi/kAxKSK0Xle0/0.jpg)](https://www.youtube.com/watch?v=kAxKSK0Xle0)

## Example
```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ramin0/got"
)

func main() {
	http.Handle("/messenger/", got.New(&got.BotData{
		BaseURL: "https://mybot.com/messenger",
		Page: &got.BotPage{
			ID:          "...",
			AccessToken: "...",
			VerifyToken: "...",
		},
	}, func(bot *got.BotData) {
		bot.StartBlock(func(b *got.BotBlock) {
			b.Component(&got.PluginFacebookProfile{})
			b.Component(&got.ComponentText{
				Text: "Hello, %{fb:first_name}!",
			})
		})
	}).Handler())

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("Listening on port %s...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
```
