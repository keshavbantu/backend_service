package controller

import (
	repo "backend_service/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func GetWord(c *gin.Context) {
	res, err := repo.GetWords()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, res)
}
