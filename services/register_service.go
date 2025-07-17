package services

import (
	"context"
	"fmt"
	"time"

	"github.com/reawphongsaphak/slOwOrk-Backend/config"
	"github.com/reawphongsaphak/slOwOrk-Backend/database"
	"github.com/reawphongsaphak/slOwOrk-Backend/models"
	"github.com/reawphongsaphak/slOwOrk-Backend/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func RegisterNewUser(user models.User) (*mongo.InsertOneResult, error){
	client := database.ConnectDB()
	cfg := config.LoadConfig()
	coll := client.Database(cfg.DB_NAME).Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	existingUser := coll.FindOne(ctx, bson.M{"email": user.Email})
	if existingUser.Err() == nil {
		return nil, fmt.Errorf("this email already register")
	} else if existingUser.Err() != mongo.ErrNoDocuments {
		return nil, existingUser.Err()
	}

	HashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	HashString := string(HashPassword)

	result, err := coll.InsertOne(
		ctx,
		bson.M{
			"user_id" 	: 1,
			"email"		: user.Email,
			"password"	: HashString,
			"resume"	: user.Resume,
			"files"		: user.Files,
			"detail"	: bson.M{
				"name"	: user.Detail.Name,
				"role"	: user.Detail.Role,
				"phone_number"	: user.Detail.PhoneNumber,
				"university"	: user.Detail.University,
			},
		},
	)

	return result, err
}