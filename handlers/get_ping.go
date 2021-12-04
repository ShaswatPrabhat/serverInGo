package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPing(context *gin.Context) {
	context.String(http.StatusOK, "pong")
}
