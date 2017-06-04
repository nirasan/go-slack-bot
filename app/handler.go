package app

import (
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
	"github.com/nlopes/slack"
	"strings"
	"os"
)


type AppHandler struct{}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

func (h *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	p, err := DecodeJSON(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch p.Type() {
	case "url_verification":
		// イベント通知先URLの認証用アクセスへのレスポンス
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(p.String("challenge")))
		return
	case "event_callback":
		// イベント通知へのレスポンス
		data, ok := p["event"].(map[string]interface{})
		if !ok {
			http.Error(w, "failed to cast", http.StatusBadRequest)
			return
		}
		pp := Payload(data)
		// ボット宛にメッセージを送信されたらレスポンスを返す
		if pp.String("type") == "message" && strings.Index(pp.String("text"), "<@U5L8SFV5W>") != -1 {
			// GAE から http リクエストを送信する場合は urlfetch ライブラリを利用する必要がある
			slack.SetHTTPClient(urlfetch.Client(ctx))
			// 環境変数からアクセストークンを取得. SLACK_TOKEN は appengine/secret.yaml に定義されている.
			token := os.Getenv("SLACK_TOKEN")
			api := slack.New(token)
			_, _, err = api.PostMessage(pp.String("channel"), "こんにちは", slack.PostMessageParameters{})
			if err != nil {
				log.Debugf(ctx, "failed to post message: %+v", err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			log.Debugf(ctx, "Success")
		}
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusOK)
		break
	}
}
