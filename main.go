package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ken2020n/gologu-server-mvp/database"
	"github.com/ken2020n/gologu-server-mvp/model"
	"github.com/ken2020n/gologu-server-mvp/model/base"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"strconv"
)

const host = "localhost:8000"

func main() {
	setupWebServer()
}

func setupWebServer() {

	// create a router
	router := gin.New()

	// global middleware
	router.Use(gin.Logger())

	// test endpoint
	router.GET("/test", HandleTest())
	router.GET("/dbtest", HandleDbTest())

	// insertion endpoints
	router.POST("/client", HandleClient())
	router.POST("/error", HandleError())
	router.POST("/http", HandleHttp())

	// query endpoints
	router.GET("/errors/:max", HandleErrorsQuery())
	router.GET("/error/:id", HandleTest())

	// run server
	err := router.Run(host)
	if err != nil {
		log.Fatalln("Server Error: ", err)
	}
}

func HandleErrorResponse(c *gin.Context) {
	if r := recover(); r != nil {
		err := fmt.Errorf("%v", r)
		c.JSON(http.StatusBadRequest, base.ErrorResponse{Message: err.Error()})
	}
}

func HandleTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer HandleErrorResponse(c)
		c.JSON(200, gin.H{"message": "API works"})
	}
}

func HandleDbTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer HandleErrorResponse(c)

		defer database.CloseClient()
		database.GetClient()

		c.JSON(200, gin.H{
			"message": "database connected",
		})
	}
}

func HandleClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer HandleErrorResponse(c)

		var request model.Client
		err := c.Bind(&request)
		if err != nil {
			c.Status(http.StatusBadRequest)
		}

		defer database.CloseClient()
		client := database.GetClient()
		coll := client.Database("log").Collection("client")
		doc := request
		result, err := coll.InsertOne(context.TODO(), doc)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"Id": result.InsertedID,
		})
	}
}

func HandleError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer HandleErrorResponse(c)

		var request model.Error
		err := c.Bind(&request)
		if err != nil {
			c.Status(http.StatusBadRequest)
		}

		defer database.CloseClient()
		client := database.GetClient()
		coll := client.Database("log").Collection("error")
		doc := request
		result, err := coll.InsertOne(context.TODO(), doc)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"Id": result.InsertedID,
		})
	}
}

func HandleErrorsQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer HandleErrorResponse(c)

		strMax := c.Param("max")
		intMax, err := strconv.ParseInt(strMax, 0, 8)
		if err != nil {
			panic(err)
		}

		fmt.Println(intMax)

		defer database.CloseClient()
		client := database.GetClient()
		coll := client.Database("log").Collection("error")

		findOptions := options.Find()
		findOptions.SetLimit(intMax)
		filter := bson.D{}

		cursor, err := coll.Find(context.TODO(), filter, findOptions)
		if err != nil {
			panic(err)
		}

		var result []model.Error
		var data model.Error
		for cursor.Next(context.TODO()) {
			if err := cursor.Decode(&data); err != nil {
				panic(err)
			}
			fmt.Println(data)
			result = append(result, data)
		}
		if err := cursor.Err(); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, result)
	}
}

func HandleHttp() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer HandleErrorResponse(c)

		var request model.Http
		err := c.Bind(&request)
		if err != nil {
			c.Status(http.StatusBadRequest)
		}

		defer database.CloseClient()
		client := database.GetClient()
		coll := client.Database("log").Collection("http")
		doc := request
		result, err := coll.InsertOne(context.TODO(), doc)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"Id": result.InsertedID,
		})
	}
}
