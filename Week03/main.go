package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	listenAddr string
)

func main() {
	var g errgroup.Group

	g, ctx := errgroup.WithContext(context.Background())

	port := 5000
	for i := port; i <= 10000; i++ {

		g.Go(test(ctx, i))
	}

}

func test(ctx context.Context, port int) {
	select {
	case <-ctx.Done():
	default:

		if port%4 == 0 {
			ctx.cancel()
		}
		server := newWebserver()
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Errorf("Could not listen on %s: %v\n", listenAddr, err)
		}
	}

}

func newWebserver() *http.Server {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return &http.Server{
		Addr:         listenAddr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
	}
}
