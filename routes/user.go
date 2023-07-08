package routes

import (
	"authentication/api/controller"
	"authentication/api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoutes(routes *gin.Engine, db *mongo.Database, client *mongo.Client) {
	UserCollectionImpl := controller.UserControllerImpl{
		UserModelImpl: models.NewUserModelImpl(db, client),
	}

	r := routes.Group("/user")
	r.POST("/register", UserCollectionImpl.Register)
	r.POST("/login", UserCollectionImpl.LoginUser)
}
