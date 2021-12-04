package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/umutsevdi/go-simple-transaction/server/cmd/config"
	"github.com/umutsevdi/go-simple-transaction/server/db/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAccounts(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/accounts" {
			app.NotFound(w)
			return
		}
		ctx := context.TODO()
		accounts := []model.Account{}
		instances, err := (*app.Collections)[0].Find(ctx, bson.D{})
		if err != nil {
			app.NotFound(w)
			return
		}
		err = instances.All(ctx, &accounts)
		if err != nil {
			app.NotFound(w)
			return
		}
		fmt.Fprintln(w, accounts)
	}
}

func GetAccountById(app *config.Application, id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p, err := primitive.ObjectIDFromHex(id)
		if r.URL.Path != "/accounts/find" {
			app.NotFound(w)
			return
		}
		accCollection := (*app.Collections)[0]
		ctx := context.TODO()
		account := model.Account{}
		instance := accCollection.FindOne(ctx, bson.M{"_id": p})
		if err != nil {
			app.NotFound(w)
			return
		}
		err = instance.Decode(&account)
		if err != nil {
			app.NotFound(w)
			return
		}
		fmt.Fprintln(w, account)
	}
}
