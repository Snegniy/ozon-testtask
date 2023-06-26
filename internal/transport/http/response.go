package http

import (
	"encoding/json"
	"github.com/Snegniy/ozon-testtask/internal/model"
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	"go.uber.org/zap"
	"net/http"
)

func writeJSON(w http.ResponseWriter, data model.UrlStorage) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		logger.Error("json Encoder failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return err
}
