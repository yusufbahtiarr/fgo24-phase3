package routers

import (
	"go-test/controllers"
	"go-test/middleware"

	"github.com/gin-gonic/gin"
)

func inventoryGroup(r *gin.RouterGroup) {
	r.GET("/transactions", middleware.AuthMiddlware(), controllers.GetTransactionHistory)
	r.POST("/transactions", middleware.AuthMiddlware(), controllers.CreateTransaction)
}
