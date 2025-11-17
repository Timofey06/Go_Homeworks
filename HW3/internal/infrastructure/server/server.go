package server

import (
	"context"
	"hw3/internal/application/interfaces"
	"hw3/internal/application/service"
	"hw3/internal/infrastructure/server/handlers"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func StartServer(port string, repo interfaces.Repository) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	s := service.NewAccountService(repo)
	h := handlers.NewHandler(s)
	mux := http.NewServeMux()
	mux.HandleFunc("/balance", h.Balance)
	mux.HandleFunc("/add", h.Add)
	mux.HandleFunc("/transfer", h.Transfer)

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
