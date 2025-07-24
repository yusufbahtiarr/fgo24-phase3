package main

import (
	"go-test/routers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	godotenv.Load()
	routers.CombineGroup(r)

	r.Run(":" + os.Getenv("APP_PORT"))
}
