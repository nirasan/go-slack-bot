package app

import (
	"net/http"
)

type AppHandler struct{}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

func (h *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	payload, err := DecodeJSON(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch payload.Type() {
	case "url_verification":
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(payload.String("challenge")))
		return
	default:
		w.Write([]byte("Hello world!"))
		break
	}
}
