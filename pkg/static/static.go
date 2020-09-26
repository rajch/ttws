// Package static allows directories on the web server host to be served statically.
package static

import (
	"net/http"

	"github.com/rajch/ttws/pkg/webserver"
)

// DefaultPath is the default path to invoke the static module - /
const DefaultPath = "/"

// AddDefaultHandler if true, will serve the web server working directory
// on the path /.
var AddDefaultHandler = false

// ServeDirectory adds a handler that serves the specified directory on the specified path.
func ServeDirectory(path string, directorypath string) {
	fs := http.FileServer(http.Dir(directorypath))
	webserver.AddHandler(path, fs.ServeHTTP)
}

func init() {
	webserver.AddInitFunc(func() {
		if AddDefaultHandler {
			fs := http.FileServer(http.Dir("."))
			webserver.AddHandler("/", fs.ServeHTTP)
		}
	})
}
