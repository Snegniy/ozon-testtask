package main

import (
	"github.com/Snegniy/ozon-testtask/internal/config"
	"github.com/Snegniy/ozon-testtask/internal/repository/memdb"
	"github.com/Snegniy/ozon-testtask/internal/repository/postgre"
	"github.com/Snegniy/ozon-testtask/internal/service"
	grpchandlers "github.com/Snegniy/ozon-testtask/internal/transport/grpc"
	"github.com/Snegniy/ozon-testtask/internal/transport/http"
	"github.com/Snegniy/ozon-testtask/pkg/graceful"
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	cfg := config.NewConfig()
	logger.Init(cfg.DebugMode)

	logger.Debug("Create router...")
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	var r service.Repository
	if cfg.StorageType == "postgre" {
		var err error
		r, err = postgre.NewRepository(cfg)
		if err != nil {
			logger.Fatal("database not open")
		}
	} else {
		r = memdb.NewRepository()
	}

	s := service.NewService(r)
	h := http.NewHttpHandlers(s)
	g := grpchandlers.NewGrpcServer(s)

	Register(router, h)
	go graceful.StartServer(router, cfg.HTTPServerHostPort)
	g.StartServer(cfg.GRPCServerHostPort)

}
func Register(r *chi.Mux, h *http.Handlers) {
	r.Post("/", h.PostLink)
	r.Get("/", h.GetLink)
}
