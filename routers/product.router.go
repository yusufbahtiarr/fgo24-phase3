package routers

import (
	"go-test/controllers"

	"github.com/gin-gonic/gin"
)

func productRouter(r *gin.RouterGroup) {
	r.GET("", controllers.GetProducts)
	r.GET("/:id", controllers.GetProductByID)
	r.GET("/category", controllers.GetCategoryProducts)
}
