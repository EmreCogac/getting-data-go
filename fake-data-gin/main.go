package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/fakeData", fakeData) // http://localhost:9090/fakeData

	router.Run(":9090") // localhost:9090
}

type words struct {
	ID  string `json:"id"`
	ENG string `json:"eng"`
	TR  string `json:"tr"`
}

var dictionary = []words{
	{ID: "1", ENG: "hello", TR: "merhaba"},
	{ID: "2", ENG: "cat", TR: "kedi"},
	{ID: "3", ENG: "car", TR: "araba"},
	{ID: "4", ENG: "door", TR: "kapı"},
}

func fakeData(c *gin.Context) {

	c.JSON(http.StatusOK, dictionary)

}
