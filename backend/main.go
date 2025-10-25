package main

import (
	"minigram-app-backend/repository"
	"minigram-app-backend/routers"
)

func main() {
	// call startDB
	repository.StartDB()

	// start routing
	routers.StartServer().Run(":8080")
}
