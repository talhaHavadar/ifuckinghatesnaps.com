// Command server serves ifuckinghatesnaps.com.
package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ifuckinghatesnaps.com/internal/site"
	"ifuckinghatesnaps.com/quotes"
)

func main() {
	addr := flag.String("addr", ":8080", "address to listen on")
	dev := flag.String("dev", "", "serve templates and static files from this directory (internal/site) and re-read them on every request, for development")
	flag.Parse()

	var opts []site.Option
	if *dev != "" {
		opts = append(opts, site.FromDisk(*dev))
		slog.Info("serving templates and static files from disk", "dir", *dev)
	}
	if err := run(*addr, opts); err != nil {
		slog.Error("server stopped", "err", err)
		os.Exit(1)
	}
}

func run(addr string, opts []site.Option) error {
	handler, err := site.New(quotes.FS, opts...)
	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	shutdownErr := make(chan error, 1)
	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		shutdownErr <- srv.Shutdown(shutdownCtx)
	}()

	slog.Info("listening", "addr", addr)
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return <-shutdownErr
}
