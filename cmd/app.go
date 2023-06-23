package cmd

import (
	"github.com/Snegniy/ozon-testtask/internal/config"
	"github.com/Snegniy/ozon-testtask/internal/repository/local"
	"github.com/Snegniy/ozon-testtask/internal/service"
	"github.com/Snegniy/ozon-testtask/internal/transport"
	"github.com/Snegniy/ozon-testtask/pkg/graceful"
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Run() {
	cfg := config.NewConfig()
	logger.Init(cfg.DebugMode)

	logger.Debug("Create router...")
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	r := local.NewRepository()
	s := service.NewService(r)
	h := transport.NewHttpHandlers(s)

	Register(router, h)
	graceful.StartServer(router, cfg.ServerHostPort)

}
func Register(r *chi.Mux, h *transport.HttpHandlers) {
	r.Post("/", h.PostLink)
	r.Get("/", h.GetLink)
}
