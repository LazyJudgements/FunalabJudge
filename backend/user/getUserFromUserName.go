package user

import (
	"context"
	"go-test/db"
	"go-test/types"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserFromUserName(c *gin.Context, userName string) *types.User {
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not available."})
		return &types.User{}
	}
	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("USERS_COLLECTION")
	collection := (client.(*mongo.Client)).Database(dbName).Collection(usrCol)

	filter := bson.M{"userName": userName}

	var user types.User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return &types.User{}
	}
	return &user
}

func GetUserFromUserId(c *gin.Context, userId int32) *types.User {
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not available."})
		return &types.User{}
	}
	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("USERS_COLLECTION")
	collection := (client.(*mongo.Client)).Database(dbName).Collection(usrCol)

	filter := bson.M{"userId": userId}

	var user types.User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return &types.User{}
	}
	return &user
}

// MakeFilterWithNonnilField関数が未完成なので動かない
func SearchUser(c *gin.Context, userInfo types.User) *types.User {
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not available."})
		return &types.User{}
	}
	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("USERS_COLLECTION")
	collection := (client.(*mongo.Client)).Database(dbName).Collection(usrCol)

	filter := db.MakeFilterWithNonnilField(userInfo)

	var user types.User
	// 基本的にusers collectionからは1つのdocumentを問い合わせる想定
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return &types.User{}
	}
	return &user
}
