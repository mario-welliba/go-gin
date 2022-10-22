package models

import (
	"context"
	"gin/forms"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// User defines user object structure
type User struct {
	ID         bson.D `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string `json:"name" bson:"name"`
	Email      string `json:"email" bson:"email"`
	Password   string `json:"password" bson:"password"`
	IsVerified bool   `json:"is_verified" bson:"is_verified"`
}

// UserModel defines the model structure
type UserModel struct{}

// Signup handles registering a user
func (u *UserModel) Signup(data forms.SignupUserCommand) (*mongo.InsertOneResult, error) {
	// Connect to the user collection
	collection := dbConnect.Use(databaseName, "user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// Assign result to error object while saving user
	userResult, err := collection.InsertOne(ctx, bson.M{
		"name":     data.Name,
		"email":    data.Email,
		"password": data.Password,
		// This will come later when adding verification
		"is_verified": false,
	})

	return userResult, err
}
