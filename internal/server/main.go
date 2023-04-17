package main

import (
	"log"
	"meeting/internal/models"
	"meeting/internal/server/router"
)

func main() {
	models.NewDB()
	e := router.Router()
	err := e.Run()
	if err != nil {
		log.Fatalln("run err.", err)
		return
	}
}
