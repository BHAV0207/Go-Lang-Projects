package main

import (
	"context"
	"log"
	"time"

	"github.com/thedevsaddam/renderer"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var rnd *renderer.Render //That package is a helper to easily send responses (JSON, HTML, XML) in Go web apps.
// rnd = your waiter → takes your data and serves it to the customer in the right format (JSON/HTML).
var db *mgo.Database //That package is a MongoDB driver for Go.
// db = your database → connects to the database and allows you to perform operations on it.

const (
	MongoURI       string = "mongodb://localhost:27017"
	dbName         string = "todo"
	collectionName string = "task"
	port           string = ":8000"
)

type todo struct {
	Id        bson.ObjectId `bson:"_id , omitempty" json:"id"`
	Title     string        `bson:"title" json:"title"`
	completed bool          `bson:"completed" json:"completed"`
	createdAt time.Time     `bson:"created_at" json:"created_at"`
}

func init() {
	rnd = renderer.New()
	// renderer.New() creates a new Render object from the renderer package.
	// It’s assigned to the global rnd variable (which we declared earlier as var rnd *renderer.Render).
	// This rnd will now be your tool for sending responses in HTTP handlers.

	// Create client options
	clientOptions := options.Client().ApplyURI(MongoURI)

	// Connect with a 10-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error: ", err)
	}

	// Ping the database to check connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping error: ", err)
	}

	// Select the database
	db = client.Database(dbName)
	log.Println("Connected to MongoDB Atlas and using DB:", dbName)
}
