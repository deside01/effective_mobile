package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/deside01/effective_mobile/internal/config"
	"github.com/deside01/effective_mobile/internal/database/db"
	"github.com/deside01/effective_mobile/internal/handlers"
	"github.com/deside01/effective_mobile/internal/repos"
	"github.com/deside01/effective_mobile/internal/router"
	"github.com/deside01/effective_mobile/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	dbpool, err := pgxpool.New(context.Background(), cfg.DB.GetDSN())
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer dbpool.Close()

	repo := repos.NewSubscriptionRepo(db.New(dbpool))
	svc := services.NewSubscriptionService(repo)
	h := handlers.NewSubscriptionHandler(svc)

	r := router.NewRouter(h)

	srv := http.Server{
		Addr:         cfg.HTTP.Addr,
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.HTTP.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.HTTP.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.HTTP.IdleTimeout) * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
