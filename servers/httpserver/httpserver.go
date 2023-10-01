package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	httpServer *http.Server
	ctx        context.Context
)

func StartUpHTTP(mainCtx context.Context, portNum int) {
	ctx = mainCtx
	port := fmt.Sprintf(":%d", portNum)
	setUpEndpoints()

	httpServer = &http.Server{
		Addr: port,
	}

	fmt.Printf("HTTP Server starting on %s\n", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("HTTP Error starting server:", err)
		return
	}
}

func setUpEndpoints() {
	// http.HandleFunc("/healthCheck", healthCheck)
	// http.HandleFunc("/addValue/", addValue)
	// http.HandleFunc("/deleteKey/", deleteKey)
	// http.HandleFunc("/getValue/", getValue)
}

func ShutdownHTTPServer(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Shutting down http server on port: %s\n", httpServer.Addr)
	httpCTX, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(httpCTX); err != nil {
		fmt.Println("Error shutting down server:", err)
		return
	}
	fmt.Println("Server gracefully shut down.")
}
