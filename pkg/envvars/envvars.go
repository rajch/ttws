package envvars

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rajch/ttws/pkg/webserver"
)

// Route is the default path to invoke the envvars module - /env
const Route = "/env"

func envvarshandler(w http.ResponseWriter, r *http.Request) {
	env := os.Environ()

	for _, variable := range env {
		fmt.Fprintln(w, variable)
	}

}

func init() {
	webserver.AddRoute(Route, envvarshandler)
}
