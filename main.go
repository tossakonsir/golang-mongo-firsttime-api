package main

import (
	"context"
	"example/controllers"
	"example/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	us          services.UserService
	ps          services.ProductService
	uc          controllers.UserController
	pc          controllers.ProductController
	ctx         context.Context
	userc       *mongo.Collection
	productc    *mongo.Collection
	mongoclient *mongo.Client
	err         error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb+srv://sunmongo:passw0rd@cluster0.r2cfi.mongodb.net/?retryWrites=true&w=majority")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	userc = mongoclient.Database("store").Collection("users")
	us = services.NewUserService(userc, ctx)
	uc = controllers.New(us)
	productc = mongoclient.Database("store").Collection("products")
	ps = services.NewProductService(productc, ctx)
	pc = controllers.NewProduct(ps)
	server = gin.Default()

}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	uc.RegisterUserRoutes(basepath)
	basepath2 := server.Group("/v2")
	pc.RegisterProductRoutes(basepath2)

	log.Fatal(server.Run(":8080"))

}
