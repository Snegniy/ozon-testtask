package http

import (
	"encoding/json"
	"github.com/Snegniy/ozon-testtask/internal/model"
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"net/http"
)

type Handlers struct {
	srv Services
}

func NewHttpHandlers(srv Services) *Handlers {
	logger.Debug("new http handlers")
	return &Handlers{srv: srv}
}

type Services interface {
	GetShortLink(ctx context.Context, url string) (model.UrlStorage, error)
	GetBaseLink(ctx context.Context, url string) (model.UrlStorage, error)
}

func (h *Handlers) PostLink(w http.ResponseWriter, r *http.Request) {
	logger.Debug("post http link")
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
	logger.Debug("get http link")
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
