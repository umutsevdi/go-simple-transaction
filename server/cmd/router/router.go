package router

import (
	"net/http"

	"github.com/umutsevdi/go-simple-transaction/server/cmd/config"
	"github.com/umutsevdi/go-simple-transaction/server/cmd/handler"
)

func SetRoutes(app *config.Application) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Home(app))
	mux.HandleFunc("/accounts", handler.GetAccounts(app))

	return mux
}
