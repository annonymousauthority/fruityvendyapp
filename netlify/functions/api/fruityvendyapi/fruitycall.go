package main

import (
	"fruityvendy.com/datagenerator"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/fruits", datagenerator.GetFruits)
	r.Run("localhost:3000")
}
