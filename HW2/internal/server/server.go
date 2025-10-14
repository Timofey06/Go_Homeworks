package server

import (
	"context"
	"hw2/internal/server/handlers"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func StartServer(port string) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()
	mux.HandleFunc("/version", handlers.VersionHandler)
	mux.HandleFunc("/decode", handlers.DecodeHandler)
	mux.HandleFunc("/hard-op", handlers.HardOpHandler)

	serv := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	go func() {
		slog.Info("Сервер запущен", "addr", serv.Addr)
		if err := serv.ListenAndServe(); err != nil {
			slog.Error("Ошибка сервера", "error", err)
		}
	}()

	<-ctx.Done()
	slog.Info("graceful shotdown...")
	if err := serv.Shutdown(context.Background()); err != nil {
		slog.Error("Ошибка при завершении сервера", "error", err)
	}
}
