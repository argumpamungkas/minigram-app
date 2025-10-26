package main

import (
	"minigram-app-backend/config"
	"minigram-app-backend/routers"
)

func main() {
	// call startDB
	config.ConnectDatabase()

	// start routing
	routers.StartServer().Run(":8080")
}
