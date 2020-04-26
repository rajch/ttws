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

// Route is the default path to invoke the filesystem module - /files
const Route = "/files/"

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
	webserver.AddRoute(Route, filesystemhandler)
}
