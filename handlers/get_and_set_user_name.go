package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var db = make(map[string]string)

type UserValue struct {
	User  string `json:"user" binding:"required,validateUserAndValue"`
	Value string `json:"value" binding:"required,validateUserAndValue"`
}

func GetUserName(context *gin.Context) {
	user := context.Params.ByName("name")
	value, ok := db[user]
	if ok {
		context.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	} else {
		context.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}

func SetUserName(context *gin.Context) {
	var userValue UserValue
	log.Println(context.Request.Body)

	err := context.BindJSON(&userValue)
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
	} else {
		db[userValue.User] = userValue.Value
		context.JSON(http.StatusOK, gin.H{"user": userValue.User, "value": userValue.Value})
	}
}
