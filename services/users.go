package services

import (
	"net/http"

	dto "study-api/dto"
	models "study-api/model"
	mongo "study-api/mongo"

	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUsers(c *gin.Context) {
	cursor, err := mongo.MongoClient.Database("users").Collection("users").Find(context.TODO(), bson.D{{}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var users []bson.M
	if err = cursor.All(context.TODO(), &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
func GetUserById(c *gin.Context) {

	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	collection := mongo.MongoClient.Database("users").Collection("users")

	var user models.User

	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&user)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {

	var body dto.CreateUser

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email is required",
		})
		return
	}

	collection := mongo.MongoClient.Database("users").Collection("users")

	filter := bson.M{
		"$or": []bson.M{
			{"email": body.Email},
			{"user": body.User},
		},
	}

	var existing models.User
	err := collection.FindOne(context.TODO(), filter).Decode(&existing)

	if err == nil {

		if existing.Email == body.Email {
			c.JSON(http.StatusConflict, gin.H{
				"error": "email already exists",
			})
			return
		}

		if existing.UserName == body.User {
			c.JSON(http.StatusConflict, gin.H{
				"error": "username already exists",
			})
			return
		}
	}

	newUser := models.User{
		UserName: body.User,
		Email:    body.Email,
		Password: body.Password,
	}

	result, err := collection.InsertOne(context.TODO(), newUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser.ID = result.InsertedID.(primitive.ObjectID)

	c.JSON(http.StatusCreated, newUser)
}
