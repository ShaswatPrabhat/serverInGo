package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"serverInGo/handlers"
)

func validateUserAndValue(fl validator.FieldLevel) bool {
	if len(fl.Field().String()) < 3 {
		return false
	}
	return true
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", handlers.GetPing)
	r.GET("/getUser/:name", handlers.GetUserName)
	r.POST("/setUser", handlers.SetUserName)
	r.GET("/getUserFromServiceB", handlers.HitServiceB)

	return r
}

func main() {
	validatorEngine := binding.Validator.Engine()
	v, ok := validatorEngine.(*validator.Validate)
	if ok {
		v.RegisterValidation("validateUserAndValue", validateUserAndValue)
	}
	r := setupRouter()
	r.Run(":8080")
}
