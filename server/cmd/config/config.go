package config

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Collections *[]*mongo.Collection
	Ctx *context.Context
}

func (app *Application) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}
