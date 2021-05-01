package main

import (
	"log"
	"net/http"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

var pageSize string = "1"

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

func setPageSize(c *gin.Context) {
	pageSize = c.DefaultQuery("pageSize", "1")
	totalPage, err := models.GetTotalCount(pageSize)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err})
		log.Panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"totalpage": totalPage})
}

func itemsGetHandler(c *gin.Context) {
	itemList, err := models.GetAllItems()
	if err != nil {
		log.Panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"itemlist": itemList})
}

func itemsGetPartial(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	itemList, err := models.GetPageOfItems(page, pageSize)
	if err != nil {
		log.Panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"itemlist": itemList})
}

func itemsPostHandler(c *gin.Context) {
	itemName := c.PostForm("itemName")
	err := models.AddItem("0", itemName)
	if err != nil {
		log.Panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func itemsPatchHandler(c *gin.Context) {
	var rsp struct {
		IsDone string `json:"isDone"`
	}
	itemId := c.Param("id")
	if err := c.ShouldBindJSON(&rsp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
