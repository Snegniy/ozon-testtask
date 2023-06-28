package rest

import (
	"context"
	"encoding/json"
	"github.com/Snegniy/ozon-testtask/internal/model"
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"
)

type Handlers struct {
	srv Services
}

func NewHttpHandlers(srv Services) *Handlers {
	//logger.Debug("new rest handlers")
	return &Handlers{srv: srv}
}

func (h *Handlers) Register(r *chi.Mux) {
	r.Post("/", h.PostLink)
	r.Get("/", h.GetLink)
}

type Handler interface {
	PostLink(w http.ResponseWriter, r *http.Request)
	GetLink(w http.ResponseWriter, r *http.Request)
}

type Services interface {
	GetShortLink(ctx context.Context, url string) (model.UrlStorage, error)
	GetBaseLink(ctx context.Context, url string) (model.UrlStorage, error)
}

func (h *Handlers) PostLink(w http.ResponseWriter, r *http.Request) {
	logger.Debug("post rest link")
	body := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		logger.Error("json Decoder failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rawUrl, ok := body["url"]
	if !ok {
		logger.Warn("request - incorrect format", zap.Error(err))
		http.Error(w, "incorrect format", http.StatusBadRequest)
		return
	}
	res, err := h.srv.GetShortLink(r.Context(), rawUrl)

	if err != nil {
		logger.Warn("response not found", zap.Error(err))
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = writeJSON(w, res)
}

func (h *Handlers) GetLink(w http.ResponseWriter, r *http.Request) {
	logger.Debug("get rest link")
	body := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		logger.Error("json Decoder failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rawShortUrl, ok := body["url"]
	if !ok {
		logger.Warn("request - incorrect format", zap.Error(err))
		http.Error(w, "incorrect format", http.StatusBadRequest)
		return
	}
	res, err := h.srv.GetBaseLink(r.Context(), rawShortUrl)
	if err != nil {
		logger.Warn("response not found", zap.Error(err))
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_ = writeJSON(w, res)
}
