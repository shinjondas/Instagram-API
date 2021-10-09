package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//creating structure for User
type User struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string `json:"name,omitempty bson:"name,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

func main() {
	fmt.Println("Server started...")
	connect()
	handleRequest()
}

var client *mongo.Client

//Establish MongoDB connection
func connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.NewClient(clientOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected to MondoDB Server")
	}

}

//Create function to handle all GET/POST Requests and routing
func handleRequest() {
	http.HandleFunc("/users", createUser)
	http.HandleFunc("/userGet", getAllUsers)
	http.HandleFunc("/users/", searchUser)
	/*http.HandleFunc("/posts", createPost)
	http.HandleFunc("/postGet", getAllPosts)
	http.HandleFunc("/posts/", searchPost)*/
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

//Getting all users
func getAllUsers(response http.ResponseWriter, request *http.Request) {
	var users []User
	collection := client.Database("test").Collection("User_Test")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}
	if err = cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(users)
}

//Creating user for platform
func createUser(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	decoder := json.NewDecoder(request.Body)
	var newUser User
	log.Println("Created user at:", time.Now())
	err := decoder.Decode(&newUser)
	if err != nil {
		panic(err)
	}
	log.Println(newUser.ID)
	collection := client.Database("test").Collection("User_Test")
	insertResult, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted post with ID:", insertResult.InsertedID)
}

//Search user from URL

func searchUser(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var id string = request.URL.Path
	id = strings.TrimPrefix(id, "/users/")
	var user User
	collection := client.Database("test").Collection("User_Test")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	log.Println(user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	fmt.Println("Returned User ID NO : ", user.ID)
	json.NewEncoder(response).Encode(user)
}
