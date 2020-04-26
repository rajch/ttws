package webserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	//server       http.Server
	//serverMux    *http.ServeMux
	//serverClosed chan struct{}
	routes    map[string]func(http.ResponseWriter, *http.Request)
	homeroute string
)

// AddRoute adds a handler to the webserver's routes.
func AddRoute(route string, routeHandler func(http.ResponseWriter, *http.Request)) {
	if routes == nil {
		routes = make(map[string]func(http.ResponseWriter, *http.Request))
	}

	routes[route] = routeHandler
}

// SetHome sets a route as the home route, which will respond to /
func SetHome(route string) {
	homeroute = route
}

// Serve starts the tiny test web server.
// A handler is set up for SIGINT and SIGTERM
func Serve() {
	serverMux := http.NewServeMux()
	server := http.Server{Addr: ":8080", Handler: serverMux}

	// Use a channel to signal server closure
	serverClosed := make(chan struct{})

	// Set up routes
	for route, routeHandler := range routes {
		serverMux.HandleFunc(route, routeHandler)
	}

	// Set up home route, if specified and valid
	if homeroute != "" {
		roothandler, ok := routes[homeroute]
		if ok {
			serverMux.HandleFunc("/", roothandler)
		}
	}

	go func() {
		signalReceived := make(chan os.Signal, 1)

		// Handle SIGINT
		signal.Notify(signalReceived, os.Interrupt)
		// Handle SIGTERM
		signal.Notify(signalReceived, syscall.SIGTERM)

		// Wait for signal
		<-signalReceived

		log.Println("Server shutting down...")
		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Fatalf("Error during HTTP server Shutdown: %v", err)
		}

		close(serverClosed)
	}()

	// Start listening using the server
	log.Println("Server starting...")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("The server failed with the following error:%v\n", err)
	}

	<-serverClosed

	log.Println("Server shut down.")
}
