package main

import (
	ctrl "backend_service/controller"
	repo "backend_service/repository"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	var err error
	connStr := "user=wordracer_db_user dbname=wordracer_db password=lHbel8aPE98WavAfYErTaor3orpNC7Dh port=5432 host=ohio-postgres.render.com sslmode=require"
	repo.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer repo.DB.Close()
	r.GET("/getword", ctrl.GetWord)
	r.Run()
}
