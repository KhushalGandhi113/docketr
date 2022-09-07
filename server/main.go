package main

import (
	"docker/server/models"
	"docker/server/routers"
	"fmt"
)

func main() {
	r := routers.RegisterRoutes()
	models.ConnectDatabase()

	fmt.Println("Successfully connected")
	r.Run("localhost:8080")
}
