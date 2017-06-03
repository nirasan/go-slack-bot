package gae

import (
	"github.com/nirasan/go-slack-bot/app"
	"net/http"
)

func init() {
	http.Handle("/", app.NewAppHandler())
}
