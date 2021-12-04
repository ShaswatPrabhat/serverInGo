package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

var db = make(map[string]string)

type UserValue struct {
	User  string `json:"user" binding:"required,validateUserAndValue"`
	Value string `json:"value" binding:"required,validateUserAndValue"`
}

func validateUserAndValue(fl validator.FieldLevel) bool {
	if len(fl.Field().String()) < 3 {
		return false
	}
	return true
}

func getPing(context *gin.Context) {
	context.String(http.StatusOK, "pong")
}
func getUserName(context *gin.Context) {
	user := context.Params.ByName("name")
	value, ok := db[user]
	if ok {
		context.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	} else {
		context.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}
func setUserName(context *gin.Context) {
	var userValue UserValue
	log.Println(context.Request.Body)

	err := context.BindJSON(&userValue)
	if err != nil {
		context.String(http.StatusBadRequest, "Bad Request MF!!")
	} else {
		db[userValue.User] = userValue.Value
		context.JSON(http.StatusOK, gin.H{"user": userValue.User, "value": userValue.Value})
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", getPing)
	r.GET("/getUser/:name", getUserName)
	r.POST("/setUser", setUserName)
	return r
}

func main() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validateUserAndValue", validateUserAndValue)
	}
	r := setupRouter()
	r.Run(":8080")
}
