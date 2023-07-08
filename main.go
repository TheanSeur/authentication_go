package main

import (
	"authentication/api/routes"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	r := gin.Default()
	r.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Next()
	})
	client, err := loadMongodbConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	db := client.Database("userLogin")
	routes.UserRoutes(r, db, client)

	r.Run(":7070")
}

func loadMongodbConfig() (*mongo.Client, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	return client, nil
}
