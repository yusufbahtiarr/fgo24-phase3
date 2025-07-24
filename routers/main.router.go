package routers

import "github.com/gin-gonic/gin"

func CombineGroup(r *gin.Engine) {
	authRouter(r.Group("/auth"))
	productRouter(r.Group("/products"))
	inventoryGroup(r.Group("/inventory"))
}
