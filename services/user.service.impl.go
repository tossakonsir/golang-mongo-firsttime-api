package services

import (
	"context"
	"errors"
	"example/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userconllection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: userconllection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err

}

func (u *UserServiceImpl) Login(user *models.User) error {
	filter := bson.D{primitive.E{Key: "username", Value: user.Username}, primitive.E{Key: "password", Value: user.Password}}
	err := u.usercollection.FindOne(u.ctx, filter).Decode(&user)
	if err != nil {
		return errors.New("no matched document found for Login")
	}

	return nil

}
