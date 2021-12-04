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
	databaseName := flag.String("db", "app", "Name of the database")
	flag.Parse()
	fmt.Println("\nflag.uri:\t" + *uri + "flag.port:\t" + *port + "\ndb:\t" + *databaseName)

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

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &config.Application{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		Collections: InitDB(client,&ctx,*databaseName),
		Ctx:    &ctx,
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

func InitDB(client *mongo.Client,ctx *context.Context, dbName string)(*[]*mongo.Collection){
	databases, err := client.ListDatabaseNames(*ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	db := client.Database(dbName)

	collections, err := client.Database(dbName).ListCollectionNames(*ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(collections)

	// Checking collections using prime number method
	flag := 1
	for i := 0; i < len(collections); i++ {
		if collections[i] == "account" {
			flag *= 2
			fmt.Println("skipping account ")
		}else if collections[i] == "card"{
			flag *= 3
			fmt.Println("skipping card ")
		}else if collections[i] == "transaction"{
			flag *= 5
			fmt.Println("skipping transaction ")
		}
	}
	fmt.Println(flag)

	if flag % 2 != 0{
		db.CreateCollection(*ctx, "account")
		fmt.Println("creating account")
	}

	if flag % 3 != 0{
		db.CreateCollection(*ctx, "card")
		fmt.Println("creating card")
	}

	if flag % 5 != 0{
		db.CreateCollection(*ctx, "transaction")
		fmt.Println("creating transaction")
	}
	
	collectionArray := make([]*mongo.Collection,3)
	collectionArray[0] = db.Collection("account")
	collectionArray[1] = db.Collection("card")
	collectionArray[2] = db.Collection("transaction")
	return 	&collectionArray
}