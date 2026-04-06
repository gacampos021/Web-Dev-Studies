package services

import (
	"net/http"

	"study-api/dto"
	"study-api/models"
	"study-api/mongo"

	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUsers(c *gin.Context) {
	cursor, err := mongo.MongoClient.Database("users").Collection("users").Find(context.TODO(), bson.D{{}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
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
			"status": http.StatusBadRequest,
			"error":  "invalid id",
		})
		return
	}

	users := mongo.MongoClient.Database("users").Collection("users")

	var user models.User

	err = users.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&user)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {

	var body dto.CreateUser

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	if body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "email is required",
		})
		return
	}

	users := mongo.MongoClient.Database("users").Collection("users")

	filter := bson.M{
		"$or": []bson.M{
			{"email": body.Email},
			{"user": body.User},
		},
	}

	var existing models.User
	err := users.FindOne(context.TODO(), filter).Decode(&existing)

	if err == nil {

		if existing.Email == body.Email {
			c.JSON(http.StatusConflict, gin.H{
				"status": http.StatusConflict,
				"error":  "email already exists",
			})
			return
		}

		if existing.UserName == body.User {
			c.JSON(http.StatusConflict, gin.H{
				"status": http.StatusConflict,
				"error":  "username already exists",
			})
			return
		}
	}

	newUser := models.User{
		UserName: body.User,
		Email:    body.Email,
		Password: body.Password,
		Name:     body.Name,
		Age:      body.Age,
	}

	result, err := users.InsertOne(context.TODO(), newUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser.ID = result.InsertedID.(primitive.ObjectID)

	c.JSON(http.StatusCreated, newUser)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "invalid id",
		})
		return
	}

	var body dto.UpdateUser

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	users := mongo.MongoClient.Database("users").Collection("users")

	var updatedUser models.User

	errFind := users.FindOneAndUpdate(
		context.TODO(),
		bson.M{"_id": objectID},
		bson.M{"$set": body},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedUser)

	if errFind != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errFind.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "invalid id",
		})
		return
	}

	users := mongo.MongoClient.Database("users").Collection("users")

	errFind := users.FindOneAndDelete(
		context.TODO(),
		bson.M{"_id": objectID},
	).Err()

	if errFind != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errFind.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
