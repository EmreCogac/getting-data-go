package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/fakeData", fakeData)

	router.Run()
}

type words struct {
	ID       string `json:"id"`
	TitleEng string `json:"eng"`
	TitleTr  string `json:"tr`
}

var word = []words{
	{ID: "1", TitleEng: "hello", TitleTr: "merhaba"},
	{ID: "2", TitleEng: "cat", TitleTr: "kedi"},
	{ID: "3", TitleEng: "car", TitleTr: "araba"},
	{ID: "4", TitleEng: "door", TitleTr: "kapı"},
}

func fakeData(c *gin.Context) {

	c.JSON(http.StatusOK, word)

}
