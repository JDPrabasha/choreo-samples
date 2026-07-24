package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {

	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/timeout-test/sleep", sleep)

	serverPort := 9090
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", serverPort),
		Handler: serverMux,
	}
	go func() {
		log.Printf("Starting HTTP Timeout Test service on port %d\n", serverPort)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP ListenAndServe error: %v", err)
		}
		log.Println("HTTP server stopped serving new requests.")
	}()

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	<-stopCh // Wait for shutdown signal

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Shutting down the server...")
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Shutdown complete.")
}

// sleep blocks for `seconds` (default 90) before responding, so a caller can
// observe whether the gateway/enforcer cuts the request off before the
// response is actually sent.
func sleep(w http.ResponseWriter, r *http.Request) {
	seconds := 90
	if v := r.URL.Query().Get("seconds"); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil && parsed > 0 {
			seconds = parsed
		}
	}

	start := time.Now()
	log.Printf("Sleep request received, sleeping for %ds\n", seconds)
	time.Sleep(time.Duration(seconds) * time.Second)
	elapsed := time.Since(start)

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Slept for %s (requested %ds), responded at %s\n", elapsed, seconds, time.Now().Format(time.RFC3339))
}
