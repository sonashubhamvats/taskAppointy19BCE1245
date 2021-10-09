package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	getUserRe   = regexp.MustCompile(`^\/users\/([a-zA-Z0-9]+)$`)
	getPostRe   = regexp.MustCompile(`^\/posts\/([a-zA-Z0-9]+)$`)
	getUserPost = regexp.MustCompile(`^\/posts\/users\/([a-zA-Z0-9]+)$`)
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
}
type Post struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User_id    string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Caption    string             `json:"caption,omitempty" bson:"caption,omitempty"`
	Image_url  string             `json:"image_url,omitempty" bson:"image_url,omitempty"`
	Time_stamp string             `json:"time_stamp,omitempty" bson:"time_stamp,omitempty"`
}

var collection *mongo.Collection
var client *mongo.Client
var ctx = context.TODO()

func CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	if client == nil {
		fmt.Println("Client nil")
	} else {
		collection := client.Database("instagram").Collection("users")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		result, _ := collection.InsertOne(ctx, user)
		json.NewEncoder(w).Encode(result)
	}
}
func CreatePostsEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	if client == nil {
		fmt.Println("Client nil")
	} else {
		collection := client.Database("instagram").Collection("posts")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		result, _ := collection.InsertOne(ctx, post)
		json.NewEncoder(w).Encode(result)
	}
}
func GetUsersEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	matches := getUserRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	var users []User
	if client == nil {
		fmt.Println("Client nil")
	} else {
		collection := client.Database("instagram").Collection("users")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			internalServerError(w, r)
			return
		}
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var user User
			cursor.Decode(&user)
			if user.ID.Hex() == matches[1] {
				users = append(users, user)
			}
		}
		if err := cursor.Err(); err != nil {
			internalServerError(w, r)
			return
		}
		json.NewEncoder(w).Encode(users)

	}

}
func GetPostsEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	matches := getPostRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	var posts []Post
	if client == nil {
		fmt.Println("Client nil")
	} else {
		collection := client.Database("instagram").Collection("posts")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			internalServerError(w, r)
			return
		}
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var post Post
			cursor.Decode(&post)
			if post.ID.Hex() == matches[1] {
				posts = append(posts, post)
			}
		}
		if err := cursor.Err(); err != nil {
			internalServerError(w, r)
			return
		}
		json.NewEncoder(w).Encode(posts)

	}

}
func GetUserPostsEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	matches := getUserPost.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	var posts []Post
	if client == nil {
		fmt.Println("Client nil")
	} else {
		collection := client.Database("instagram").Collection("posts")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			internalServerError(w, r)
			return
		}
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var post Post
			cursor.Decode(&post)
			if post.User_id == matches[1] {
				posts = append(posts, post)
			}
		}
		if err := cursor.Err(); err != nil {
			internalServerError(w, r)
			return
		}
		json.NewEncoder(w).Encode(posts)

	}

}
func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error":"not found"}`))
}
func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"error":"internal server error"}`))
}
func main() {
	fmt.Println("Starting the connection")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	if client == nil {
		fmt.Println("Null")
	} else {
		fmt.Println("Connected")
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/users", CreateUserEndpoint)
	mux.HandleFunc("/users/", GetUsersEndPoint)
	mux.HandleFunc("/posts", CreatePostsEndpoint)
	mux.HandleFunc("/posts/", GetPostsEndPoint)
	mux.HandleFunc("/posts/users/", GetUserPostsEndPoint)
	http.ListenAndServe("localhost:8080", mux)
}
