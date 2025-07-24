package routers

import (
	"go-test/controllers"
	"go-test/middleware"

	"github.com/gin-gonic/gin"
)

func productRouter(r *gin.RouterGroup) {
	r.GET("", middleware.AuthMiddlware(), controllers.GetProducts)
	r.GET("/:id", middleware.AuthMiddlware(), controllers.GetProductByID)
	r.GET("/category", middleware.AuthMiddlware(), controllers.GetCategoryProducts)
}
