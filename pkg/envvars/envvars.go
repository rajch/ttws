package envvars

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rajch/ttws/pkg/webserver"
)

// DefaultPath is the default path to invoke the envvars module - /env
const DefaultPath = "/env"

func envvarshandler(w http.ResponseWriter, r *http.Request) {
	env := os.Environ()

	for _, variable := range env {
		fmt.Fprintln(w, variable)
	}

}

func init() {
	webserver.AddHandler(DefaultPath, envvarshandler)
}
