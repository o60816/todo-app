package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

func showMainPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{},
	)
}

func requireResource(c *gin.Context) {
	resourceName := c.Param("name")
	c.HTML(
		http.StatusOK,
		resourceName,
		gin.H{},
	)
}

// func usersGetHandler(c *gin.Context) {

// }

// func usersPostHandler(c *gin.Context) {

// }

// func usersPatchHandler(c *gin.Context) {

// }

// func usersDeleteHandler(c *gin.Context) {

// }

func itemsGetHandler(c *gin.Context) {
	itemList, err := models.GetAllItems()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err})
		log.Panic(err)
		return
	}
	c.JSON(http.StatusOK, itemList)
}

func itemsPostHandler(c *gin.Context) {
	itemName := c.PostForm("itemName")
	err := models.AddItem("0", itemName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err})
		log.Panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Add successfully"})
}

func itemsPatchHandler(c *gin.Context) {
	type Response struct {
		IsDone string `json:"isDone"`
	}

	var rsp Response
	itemId := c.Param("id")
	if err := c.ShouldBindJSON(&rsp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("abcd:" + rsp.IsDone)
	err := models.UpdateItem(rsp.IsDone, itemId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err})
		log.Panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update successfully"})
}

func itemDeleteAllHandler(c *gin.Context) {
	err := models.DeleteAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err})
		log.Panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete successfully"})
}

func itemsDeleteHandler(c *gin.Context) {
	err := models.DeleteItem(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err})
		log.Panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete successfully"})
}
