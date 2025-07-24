package routers

import (
	"go-test/controllers"

	"github.com/gin-gonic/gin"
)

func authRouter(r *gin.RouterGroup) {
	r.POST("/login", controllers.Login)
}
