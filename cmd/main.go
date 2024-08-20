package main

import(
	"93HW/config"
	"93HW/routes"
	"log"
)

func main(){
	router := routes.SetupRouter()

	config.ConnectRedis()

	log.Println("Server is running on port :8080")
	router.Run(":8080")
}