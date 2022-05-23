package controllers

import (
	"example/models"
	"example/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService services.ProductService
}

func NewProduct(productservice services.ProductService) ProductController {
	return ProductController{
		ProductService: productservice,
	}
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product
	token := ctx.Request.Header.Get("token")
	// token := services.JWTAuthService().GetJWTAuthFromRedis(username)
	product.Token = token
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := pc.ProductService.CreateProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (pc *ProductController) GetAll(ctx *gin.Context) {
	users, err := pc.ProductService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (pc *ProductController) UpdateProduct(ctx *gin.Context) {
	var product models.Product
	token := ctx.Request.Header.Get("token")
	// token := services.JWTAuthService().GetJWTAuthFromRedis(username)
	product.Token = token
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := pc.ProductService.UpdateProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (pc *ProductController) DeleteProduct(ctx *gin.Context) {
	var itemname string = ctx.Param("itemname")
	// token := ctx.Request.Header.Get("token")
	// services.JWTAuthService().GetJWTAuthFromRedis(username)
	err := pc.ProductService.DeleteProduct(&itemname)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (pc *ProductController) RegisterProductRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/product")
	userroute.POST("/create", pc.CreateProduct)
	userroute.GET("/listproduct", pc.GetAll)
	userroute.PATCH("/update", pc.UpdateProduct)
	userroute.DELETE("/delete/:itemname", pc.DeleteProduct)
}
