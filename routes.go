// routes.go

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func initializeRoutes() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	router.LoadHTMLGlob(fmt.Sprintf("website/*"))

	router.GET("/", showMainPage)

	router.GET("/website/:name", requireResource)

	// router.GET("/users", usersGetHandler)

	// router.GET("/users/:name", usersGetHandler)

	// router.POST("/users", usersPostHandler)

	// router.PATCH("/users/:name", usersPatchHandler)

	// router.DELETE("/users", usersDeleteHandler)

	// router.DELETE("/users/:name", usersDeleteHandler)

	router.GET("/items", itemsGetHandler)

	router.POST("/items", itemsPostHandler)

	router.PATCH("/items/:id", itemsPatchHandler)

	router.DELETE("/items/", itemDeleteAllHandler)

	router.DELETE("/items/:id", itemsDeleteHandler)

	router.Run(":80")
}
