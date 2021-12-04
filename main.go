package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"os"
	"serverInGo/handlers"
)

type Configuration struct {
	Port string `json:"port"'`
}

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

func LoadConfig() (port string) {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error in while opening config file", err, " defaulting to default port 8080")
		port = "8080"
		return
	}
	port = configuration.Port
	return
}
func main() {
	validatorEngine := binding.Validator.Engine()
	v, ok := validatorEngine.(*validator.Validate)
	if ok {
		v.RegisterValidation("validateUserAndValue", validateUserAndValue)
	}
	r := setupRouter()

	r.Run(":" + LoadConfig())
}
