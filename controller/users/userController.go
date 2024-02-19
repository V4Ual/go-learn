package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vishalsharma/api/config"
	"github.com/vishalsharma/api/model"
	"github.com/vishalsharma/api/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// var user model.User
	u := model.User{}
	_ = json.NewDecoder(r.Body).Decode(&u)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	u.ID = primitive.NewObjectID()
	u.Password = string(hashedPassword)
	_, err = config.UserCollection.InsertOne(context.Background(), &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	r.Body.Close()

	// var response responses.APIResponse
	// response.Status = true
	// response.Data = u
	// response.StatusCode = 200

  responses.SuccessResponse(200,"Create successFully",u)

}

func Login(w http.ResponseWriter, r *http.Request) {
	u := model.User{}
	_ = json.NewDecoder(r.Body).Decode(&u)
	filter := bson.D{{"name", u.Name}}
	password := u.Password
	err := config.UserCollection.FindOne(context.TODO(), filter).Decode(&u)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Println([]byte(u.Password), []byte(password))
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	} else {

		json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
		return
	}

}
