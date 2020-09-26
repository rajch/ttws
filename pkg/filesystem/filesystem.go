// Package filesystem registers a handler that emits directory listings from the web server host filesystem.
// The path for this handler needs a trailing slash. This submits all requests beginning with that
// path to the handler, which parses the url after the path as the filesystem path whose directory
// listing it shows. The listing is three directory levels deep by default. This depth can be changed
// by including a query string parameter called 'depth'.
package filesystem

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/rajch/ttws/pkg/webserver"
)

// DefaultPath is the default path to invoke the filesystem module - /files
const DefaultPath = "/files/"

func getFiles(startpath string, depth int, currentdepth int) string {
	var result string

	entries, err := ioutil.ReadDir(startpath)
	if err != nil {
		result = err.Error()
		return result
	}

	for _, entry := range entries {
		if currentdepth > 0 {
			result += strings.Repeat("   ", currentdepth)

			if entry.IsDir() {
				result += "\\"
			} else {
				result += "|"
			}

			result += "--"
		}

		result += entry.Name() + "\n"

		if entry.IsDir() && currentdepth < depth {
			result += getFiles(path.Join(startpath, entry.Name()), depth, currentdepth+1)
		}
	}

	return result
}

func filesystemhandler(w http.ResponseWriter, r *http.Request) {
	startpath := strings.TrimPrefix(r.URL.Path, "/files")
	if startpath == "" {
		startpath = "/"
	}

	depthqs := r.URL.Query()["depth"]
	depth := 1
	if len(depthqs) > 0 {
		depth, _ = strconv.Atoi(depthqs[0])
	}

	fmt.Fprintln(w, getFiles(startpath, depth, 0))
}

func init() {
	webserver.AddHandler(DefaultPath, filesystemhandler)
}
