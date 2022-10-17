package router

import (
	"fmt"
	itemController "microservicio/controllers/item"

	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine) {
	// Products Mapping
	//router.GET("/items/:id", bookController.Get)
	router.POST("/items", itemController.InsertItem)

	fmt.Println("Finishing mappings configurations")
}
