package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var db = make(map[string]string)

type UserValue struct {
	User  string `json:"user" binding:"required"`
	Value string `json:"value" binding:"required"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/getUser/:name", func(context *gin.Context) {
		user := context.Params.ByName("name")
		log.Println(db)
		value, ok := db[user]
		if ok {
			context.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			context.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})
	r.POST("/setUser", func(context *gin.Context) {
		var userValue UserValue
		log.Println(context.Request.Body)

		error := context.BindJSON(&userValue)
		if error != nil {
			context.String(http.StatusBadRequest, "Bad Request MF!!")
		} else {
			db[userValue.User] = userValue.Value
			context.JSON(http.StatusOK, gin.H{"user": userValue.User, "value": userValue.Value})
		}
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
