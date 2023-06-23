package transport

import (
	"encoding/json"
	"github.com/Snegniy/ozon-testtask/internal/model"
	"net/http"
)

type HttpHandlers struct {
	srv Services
}

func NewHttpHandlers(srv Services) *HttpHandlers {
	return &HttpHandlers{srv: srv}
}

type Services interface {
	GetShortLink(url string) (model.UrlBaseStorage, error)
	GetBaseLink(url string) (model.UrlShortStorage, error)
}

func (h *HttpHandlers) PostLink(w http.ResponseWriter, r *http.Request) {
	var body map[string]string
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rawUrl, ok := body["url"]
	if !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := h.srv.GetShortLink(rawUrl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HttpHandlers) GetLink(w http.ResponseWriter, r *http.Request) {
	var body map[string]string
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rawShortUrl, ok := body["url"]
	if !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := h.srv.GetBaseLink(rawShortUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
