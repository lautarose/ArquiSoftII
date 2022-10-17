package router

import (
	bookController "microservicio/controllers/book"
	"github.com/gin-gonic/gin"
	"fmt"
)

func MapUrls(router *gin.Engine) {
	// Products Mapping
	router.GET("/items/:id", bookController.Get)
	router.POST("/items", bookController.Insert)

	fmt.Println("Finishing mappings configurations")
}
