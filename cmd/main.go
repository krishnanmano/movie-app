package main

import (
	"fmt"
	"log"
	"movie_app/config"
	"movie_app/router"
	"net/http"
)

func main() {
	configData := config.LoadAppConfiguration("./dev.json")
	if configData == nil {
		panic("Configuration not found")
	}
	ginRouter := router.SetUpRouter(*configData)
	fmt.Println("Application loaded successfully ")
	log.Fatal(http.ListenAndServe(configData.Server.Address, ginRouter))
}
