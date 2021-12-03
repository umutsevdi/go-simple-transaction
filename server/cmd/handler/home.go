package handler

import (
	"fmt"
	"net/http"

	"github.com/umutsevdi/go-simple-transaction/server/cmd/config"
)

func Home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			app.NotFound(w)
			return
		}
		fmt.Fprint(w, "Home Page")
	}
}
