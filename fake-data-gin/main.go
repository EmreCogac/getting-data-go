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
	Eng string `json:"eng"`
	Tr  string `json:"tr`
}

var dictionary = []words{
	{ID: "1", Eng: "hello", Tr: "merhaba"},
	{ID: "2", Eng: "cat", Tr: "kedi"},
	{ID: "3", Eng: "car", Tr: "araba"},
	{ID: "4", Eng: "door", Tr: "kapı"},
}

func fakeData(c *gin.Context) {

	c.JSON(http.StatusOK, dictionary)

}
