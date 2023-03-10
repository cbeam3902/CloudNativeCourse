package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongodbEndpoint = "mongodb://db:27017" // Find this from the Mongo container
)

type Post struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	Body      string             `bson:"body"`
	Tags      []string           `bson:"tags"`
	Comments  uint64             `bson:"comments"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func main() {
	var db database // Database is a struct that contains a map and a mutex
	var err error
	db.client, err = mongo.NewClient(
		options.Client().ApplyURI(mongodbEndpoint),
	)
	checkError(err)
	db.ctx = context.Background()
	err = db.client.Connect(db.ctx)
	db.col = db.client.Database("blog").Collection("posts")
	mux := http.NewServeMux()                    // Function mux
	mux.HandleFunc("/list", db.list)             // Mux handle for list
	mux.HandleFunc("/read", db.read)             // Mux handle for read
	mux.HandleFunc("/create", db.create)         // Mux handle for create
	mux.HandleFunc("/update", db.update)         // Mux handle for update
	mux.HandleFunc("/delete", db.delete)         // Mux handle for delete
	log.Fatal(http.ListenAndServe(":8000", mux)) // Begin and log the server
}

type database struct { // Database struct
	ctx    context.Context
	client *mongo.Client
	col    *mongo.Collection
}

func (db *database) list(w http.ResponseWriter, req *http.Request) { // List every item and it's price in the database
	var p []Post
	curr, err := db.col.Find(db.ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if err = curr.All(db.ctx, &p); err != nil {
		log.Fatal(err)
	}
	for _, posts := range p {
		layout := fmt.Sprintf("Title: %s\nTags:\n%s\nBody:\n%s\nComment count: %d\nCreated on %s\nUpdated on %s\n", posts.Title, strings.Join(posts.Tags, ", "), posts.Body, posts.Comments, posts.CreatedAt, posts.UpdatedAt)
		fmt.Fprintf(w, "%s\n", layout)
	}
	log.Printf("posts: %+v", p)
}

func (db *database) create(w http.ResponseWriter, req *http.Request) { // Place a new item and price in the database
	title := req.URL.Query().Get("title") // Collect the requested item
	body := req.URL.Query().Get("body")   // Collect the requested price
	tags := req.URL.Query().Get("tags")
	comments := req.URL.Query().Get("comments")

	tempComments, e := strconv.ParseUint(comments, 10, 64) // Parse the price
	if e != nil {                                          // Check if the price can become a float
		w.WriteHeader(http.StatusNotFound) // Price can't become a float
		fmt.Fprintf(w, "Comment count not supported: %s\n", comments)
	} else {
		var test Post
		filter := bson.M{"title": bson.M{"$eq": title}}
		if err := db.col.FindOne(db.ctx, filter).Decode(&test); err == nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "That post is already created, please update it instead!\n")
		} else {
			res, err := db.col.InsertOne(db.ctx, &Post{
				ID:        primitive.NewObjectID(),
				Title:     title,
				Body:      body,
				Tags:      strings.Split(tags, ","),
				Comments:  tempComments,
				CreatedAt: time.Now(),
			})
			checkError(err)
			log.Printf("inserted id: %s\n", res.InsertedID.(primitive.ObjectID).Hex())
			fmt.Fprintf(w, "Posted %s into blog\n", title)
		}
	}
}

func (db *database) read(w http.ResponseWriter, req *http.Request) { // Place a new item and price in the database
	title := req.URL.Query().Get("title") // Collect the requested item

	var posts Post
	filter := bson.M{"title": bson.M{"$eq": title}}
	if err := db.col.FindOne(db.ctx, filter).Decode(&posts); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Cannot find post for that title\n")
	} else {
		layout := fmt.Sprintf("Title: %s\nTags:\n%s\nBody:\n%s\nComment count: %d\nCreated on %s\nUpdated on %s\n", posts.Title, strings.Join(posts.Tags, ", "), posts.Body, posts.Comments, posts.CreatedAt, posts.UpdatedAt)
		fmt.Fprintf(w, "%s\n", layout)
	}
}

func (db *database) update(w http.ResponseWriter, req *http.Request) { // Place a new item and price in the database
	title := req.URL.Query().Get("title") // Collect the requested item
	body := req.URL.Query().Get("body")   // Collect the requested price
	tags := req.URL.Query().Get("tags")
	comments := req.URL.Query().Get("comments")

	tempComments, e := strconv.ParseUint(comments, 10, 64) // Parse the price
	if e != nil {                                          // Check if the price can become a float
		w.WriteHeader(http.StatusNotFound) // Price can't become a float
		fmt.Fprintf(w, "Comment count not supported: %s\n", comments)
	} else {
		var test Post
		filter := bson.M{"title": bson.M{"$eq": title}}
		if err := db.col.FindOne(db.ctx, filter).Decode(&test); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Cannot find blog post for that title!\n")
		} else {
			update := bson.M{"$set": bson.M{"body": body, "tags": strings.Split(tags, ","), "comments": tempComments}, "$currentDate": bson.M{"updated_at": true}}
			_, err := db.col.UpdateOne(db.ctx, filter, update)
			checkError(err)
			log.Printf("updated post %s\n", title)
			fmt.Fprintf(w, "Updated post %s\n", title)
		}
	}
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) { // Place a new item and price in the database
	title := req.URL.Query().Get("title") // Collect the requested item

	filter := bson.M{"title": bson.M{"$eq": title}}
	var posts Post
	if err := db.col.FindOne(db.ctx, filter).Decode(&posts); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Cannot find post for that title\n")
	} else {
		if _, err := db.col.DeleteOne(db.ctx, filter); err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, "Deleted post %s!\n", title)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
