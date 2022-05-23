package controllers

import (
	"example/models"
	"example/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
	T           http.RoundTripper
}

func New(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.Login(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	tokenerr := services.JWTAuthService().GenerateToken(user.Username, true)
	if tokenerr != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": tokenerr.Error()})
		return
	}
	// uc.RoundTrip(ctx.Request, user.Username)
	// fmt.Println("#$$$$$$$$$$$$$$    " + ctx.FullPath())
	fmt.Println("token from redis    " + services.JWTAuthService().GetJWTAuthFromRedis(user.Username))
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/create", uc.CreateUser)
	userroute.POST("/login", uc.Login)
}

func (uc *UserController) RoundTrip(req *http.Request, username string) (*http.Response, error) {
	req.Header.Add("username", username)
	return uc.T.RoundTrip(req)
}
