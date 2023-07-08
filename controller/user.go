package controller

import (
	"authentication/api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserControllerImpl struct {
	UserModelImpl *models.UserModelImpl
}

func NewUserController(db *mongo.Database, client *mongo.Client) *UserControllerImpl {
	userModel := models.NewUserModelImpl(db, client)
	return &UserControllerImpl{
		UserModelImpl: userModel,
	}
}

func (c *UserControllerImpl) Register(ctx *gin.Context) {
	var user models.UserModel
	ctx.BindJSON(&user)
	user.Id = primitive.NewObjectID().Hex()
	res, err := c.UserModelImpl.RegisterUser(&user)

	if err != nil {
		ctx.JSON(400, bson.M{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, bson.M{
		"data":    res.InsertedID,
		"message": "success",
	})
}

func (c *UserControllerImpl) LoginUser(ctx *gin.Context) {
	var user models.UserModel
	ctx.BindJSON(&user)
	res := c.UserModelImpl.Login(&user)

	if res == nil {
		ctx.JSON(400, bson.M{
			"message": "incorrect",
		})
		return
	}
	ctx.JSON(200, bson.M{
		"data":    "Login successful",
		"message": "success",
	})
}
