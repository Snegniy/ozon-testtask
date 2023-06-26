package http

import (
	"encoding/json"
	"github.com/Snegniy/ozon-testtask/internal/model"
	"net/http"
)

func writeJSON(w http.ResponseWriter, data model.UrlStorage) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return err
}
