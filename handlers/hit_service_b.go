package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type DownStreamResponse struct {
	DownStreamData string `json:"downstreamData" binding:"required"`
}

func HitServiceB(context *gin.Context) {
	resp, err := http.Get("http://localhost:8090/ping")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Downstream system down"})
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			context.JSON(http.StatusInternalServerError,
				gin.H{"error": "Downstream system response reading error"})
		}
		downStreamResponse := DownStreamResponse{}
		jsonErr := json.Unmarshal(body, &downStreamResponse)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		context.JSON(http.StatusOK, gin.H{
			"responseFromDownStream": downStreamResponse,
		})
	}
	defer resp.Body.Close()
}
