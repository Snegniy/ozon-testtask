package graceful

import (
	"context"
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func StartServer(r *chi.Mux, host string) {
	logger.Debug("Start app server")

	srv := &http.Server{
		Addr:    host,
		Handler: r,
	}
	go func() {
		logger.Info("Server started", zap.String("host:port", host))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen:", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown", zap.Error(err))
	}
	logger.Info("Server exiting")
}
