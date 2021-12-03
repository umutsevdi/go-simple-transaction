package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/umutsevdi/go-simple-transaction/server/cmd/config"

	"github.com/umutsevdi/go-simple-transaction/server/cmd/router"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Configuring server attributes
	port := flag.String("port", ":3000", "Port address for the application")
	uri := flag.String("uri", "mongodb://root:password@localhost:27017", "MongoDB database access URI")
	flag.Parse()
	fmt.Println("flag.port:\t" + *port + "\nflag.uri:\t" + *uri)

	// Connecting to the MongoDB via MongoDB drive using the uri
	client, err := mongo.NewClient(options.Client().ApplyURI(*uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &config.Application{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	srv := &http.Server{
		Addr:    *port,
		Handler: router.SetRoutes(app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
