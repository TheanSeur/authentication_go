package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	Id       string `json:"_id,omitempty" bson:"_id"`
	Gmail    string `json:"gmail" bson:"gmail"`
	UserName string `json:"username" bson:"username"`
	FullName string `json:"fullname" bson:"fullname"`
	Password string `json:"password" bson:"password"`
}

type UserModelImpl struct {
	UserCollection *mongo.Collection
}

func NewUserModelImpl(db *mongo.Database, client *mongo.Client) *UserModelImpl {
	return &UserModelImpl{
		UserCollection: db.Collection("UserProfile"),
	}
}

func (u *UserModelImpl) RegisterUser(user *UserModel) (*mongo.InsertOneResult, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		return nil, err
	}

	user.Password = string(hash)

	res, err := u.UserCollection.InsertOne(context.Background(), &user)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (u *UserModelImpl) Login(req *UserModel) *UserModel {

	var userModel *UserModel

	selector := bson.M{"$or": []bson.M{{"username": req.UserName}, {"gmail": req.Gmail}}}

	err := u.UserCollection.FindOne(context.TODO(), selector).Decode(&userModel)

	if err != nil {
		return nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password))

	if err != nil {
		fmt.Println("Password incorrect")
		return nil
	}

	return userModel
}
