package main

import (
	"cart_api/internal/config"
	"cart_api/internal/db"
	"cart_api/internal/handlers"
	"cart_api/internal/repository"
	"cart_api/internal/service"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	dbConn, err := db.NewPostgres(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	repo := repository.NewRepository(dbConn)
	svc := service.NewService(repo)
	srv := handlers.NewServer(svc)

	mux := srv.RegisterRoutes()
	httpServer := &http.Server{
		Addr:    ":" + cfg.HTTP.Port,
		Handler: mux,
	}

	go func() {
		fmt.Printf("Cart API веб-сервер успешно запущен на порту %s...\n", cfg.HTTP.Port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка HTTP-сервера: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	fmt.Println(" Получен сигнал остановки! Начинаем плавное завершение...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Ошибка при плавном выключении сервера: %v", err)
	}

	fmt.Println("Сервер успешно и безопасно остановлен.")
}
