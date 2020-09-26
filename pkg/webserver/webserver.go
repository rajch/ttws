package webserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"flag"
)

var (
	handlers        map[string]func(http.ResponseWriter, *http.Request)
	roothandlerpath string

	initfuncs []func()

	// Flags
	portflag = flag.String("p", "8080", "Port on which to run the server.")
)

func parseflags() {
	if !flag.Parsed() {
		flag.Parse()
	}
}

// GetOption gets an the value for an option. The value can be, in descending order
// of preference:
// - provided as a command-line option
// - provided as an environment variable
// - the default value
func GetOption(optflag *string, envvarname string, defaultvalue string) string {
	parseflags()

	result := *optflag

	if result == "" || result == defaultvalue {
		result = os.Getenv(envvarname)
	}

	if result == "" {
		result = defaultvalue
	}

	return result
}

// AddHandler adds a handler to the webserver.
// If a handler already exists on the specified path, it will be replaced.
func AddHandler(path string, handler func(http.ResponseWriter, *http.Request)) {
	if handlers == nil {
		handlers = make(map[string]func(http.ResponseWriter, *http.Request))
	}

	handlers[path] = handler
}

// RemoveHandler removes the handler on the specified path.
// If a handler does not exist on the specified path, nothing happens.
func RemoveHandler(path string) {
	if handlers != nil {
		delete(handlers, path)
	}
}

// SetRootHandler sets a root handler, which will respond to /
// The handler has to be added using AddHandler before calling SetRootHandler.
// A nonexistant path will be ignored.
func SetRootHandler(path string) {
	roothandlerpath = path
}

// AddInitFunc adds a function to be called before starting the server.
// Handlers can use this to, for example, parse flags.
func AddInitFunc(f func()) {
	if initfuncs == nil {
		initfuncs = []func(){}
	}
	initfuncs = append(initfuncs, f)
}

// ListenAndServe starts the web server.
// It will stop on receiving SIGINT or SIGTERM.
func ListenAndServe() {
	parseflags()

	port := ":" + GetOption(portflag, "PORT", "8080")

	serverMux := http.NewServeMux()
	server := http.Server{Addr: port, Handler: serverMux}

	// Use a channel to signal server closure
	serverClosed := make(chan struct{})

	// Set up handlers
	for path, handler := range handlers {
		serverMux.HandleFunc(path, handler)
	}

	// Set up home route, if specified and valid
	if roothandlerpath != "" {
		roothandler, ok := handlers[roothandlerpath]
		if ok {
			serverMux.HandleFunc("/", roothandler)
		}
	}

	// Call init functions
	for _, initfunction := range initfuncs {
		initfunction()
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
			log.Fatalf("Error during HTTP server shutdown: %v.", err)
		}

		close(serverClosed)
	}()

	// Start listening using the server
	log.Printf("Server starting on port %v...\n", port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("The server failed with the following error: %v.\n", err)
	}

	<-serverClosed

	log.Println("Server shut down.")
}
