package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NikenCarolina/warehouse-be/internal/config"
)

type Server struct {
	http.Server
	GracePeriod time.Duration
}

func NewServer(config *config.Config, handler http.Handler) Server {
	return Server{
		Server: http.Server{
			Addr:    config.App.ServerAddress,
			Handler: handler,
		},
		GracePeriod: config.App.ServerGracePeriod,
	}
}

func (s *Server) Run() {
	go func() {
		log.Println("listening on port ", s.Addr)
		err := s.ListenAndServe()
		if err != nil {
			log.Fatalf("failed to run server on %s", s.Addr)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	timeout := time.Duration(s.GracePeriod) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	<-stop

	log.Println("shutting down server...")
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown server: %v", err)
	}

	log.Println("server shutdown gracefully")
}
