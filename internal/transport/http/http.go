package http

import (
	"encoding/json"
	"github.com/Snegniy/ozon-testtask/internal/model"
	"net/http"
)

type Handlers struct {
	srv Services
}

func NewHttpHandlers(srv Services) *Handlers {
	return &Handlers{srv: srv}
}

type Services interface {
	GetShortLink(url string) (model.UrlStorage, error)
	GetBaseLink(url string) (model.UrlStorage, error)
}

func (h *Handlers) PostLink(w http.ResponseWriter, r *http.Request) {
	body := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rawUrl, ok := body["url"]
	if !ok {
		http.Error(w, "incorrect format", http.StatusBadRequest)
		return
	}
	res, err := h.srv.GetShortLink(rawUrl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = writeJSON(w, res)
}

func (h *Handlers) GetLink(w http.ResponseWriter, r *http.Request) {
	body := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rawShortUrl, ok := body["url"]
	if !ok {
		http.Error(w, "incorrect format", http.StatusBadRequest)
		return
	}
	res, err := h.srv.GetBaseLink(rawShortUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_ = writeJSON(w, res)
}
