package app

import (
	"context"
	"crud/internal/handler"
	"crud/internal/pkg/authclient"
	"crud/internal/pkg/server"
	"crud/internal/repository/cache"
	"crud/internal/service"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
)

func Run() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	var wg sync.WaitGroup

	// initialize dbs
	DB, err := cache.RecipeCacheInit(ctx, &wg)
	if err != nil {
		log.Fatalf("ERROR failed to initialize user database: %v", err)
	}

	authclient.Init("localhost:8000")

	// initialize service
	service.Init(DB)

	go func() {
		err := server.Run("localhost:8080", handler.ServerHandler)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("ERROR server run ", err)
		}
	}()

	log.Println("INFO CRUD service is running")

	<-ctx.Done()

	if err = server.Stop(); err != nil {
		log.Fatal("ERROR server was not gracefully shutdown", err)
	}
	wg.Wait()

	log.Println("INFO CRUD service was gracefully shutdown")
}
