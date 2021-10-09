package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/shreyatr03/Instagram/helper"
	"github.com/shreyatr03/Instagram/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// for creating a new post
func PostsHandler(response http.ResponseWriter, request *http.Request) {

	switch request.Method {

	// if method is post
	case http.MethodPost:
		var post models.Post
		var user models.User

		// decode request body as post document
		json.NewDecoder(request.Body).Decode(&post)

		// reading UserId from post body
		uid := post.UserId
		log.Println("UserId of post to be created: ", uid)

		// fetching user collection
		var instaDatabase = helper.ConnectDB().Database("instadb")
		var usersCollection = instaDatabase.Collection("user")

		// filtering users where userID matches
		err := usersCollection.FindOne(context.TODO(), bson.M{"_id": uid}).Decode(&user)

		// if error (userId not found)
		if err != nil {
			log.Fatal("User does not exist: ", err)
		}

		// fetching collection of post
		postsCollection := instaDatabase.Collection("post")

		// add post time as now
		post.TimeStamp = primitive.Timestamp{T: uint32(time.Now().Unix())}

		// create post
		postResult, err1 := postsCollection.InsertOne(context.TODO(), post)

		// if error in creating post
		if err1 != nil {
			log.Fatal("Could not create post: ", err)
		}

		fmt.Println(postResult.InsertedID)

		// send back the created post Id as reponse
		json.NewEncoder(response).Encode(postResult)

	default:
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// for fetching post by id
func PostHandler(response http.ResponseWriter, request *http.Request) {

	switch request.Method {

	// if method is get
	case http.MethodGet:
		var post models.Post
		//var user models.User

		// extracting post id from request url
		id := path.Base(request.RequestURI)

		// fetching post collection
		var instaDatabase = helper.ConnectDB().Database("instadb")
		var postsCollection = instaDatabase.Collection("post")

		// converting postid to object if format
		docID, err := primitive.ObjectIDFromHex(id)
		fmt.Println("Request post ObjectId: ", docID)

		// if id is not valid
		if err != nil {
			log.Fatal("Invalid id:", err)
		}

		// fetch post with post id
		err = postsCollection.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&post)

		// if error in fetching post
		if err != nil {
			log.Fatal("Post does not exist: ", err)
		}

		// send back the post as reponse
		json.NewEncoder(response).Encode(post)

	default:
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// for fetching all the posts of an user
func UserPostsHandler(response http.ResponseWriter, request *http.Request) {

	switch request.Method {

	// if method is get
	case http.MethodGet:

		var posts models.Posts

		// extract userId from url
		uid := path.Base(request.RequestURI)

		// fetching post collection
		var instaDatabase = helper.ConnectDB().Database("instadb")
		var postsCollection = instaDatabase.Collection("post")

		// converting post id to post objectId type
		docID, err := primitive.ObjectIDFromHex(uid)
		fmt.Println("UserId whose posts are to be fetched: ", docID)

		// if id not valid
		if err != nil {
			log.Fatal("Invalid id: ", err)
		}

		// getting cursor to filtered posts based on userId
		cursor, err := postsCollection.Find(context.TODO(), bson.D{{"userId", docID}})

		// if error in filtering
		if err != nil {
			log.Println(err)
		} else {

			var post models.Post

			// if cursor next value exists, continue loop
			for cursor.Next(context.TODO()) {

				//decode every element in array as post
				err := cursor.Decode(&post)

				if err != nil {
					fmt.Print(err)
				}

				//append post to the posts array
				posts = append(posts, post)
			}
		}

		//send back the posts array as reponse
		json.NewEncoder(response).Encode(posts)
	default:
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
