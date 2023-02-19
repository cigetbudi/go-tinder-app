package main

import (
	"go-tinder-app/api/routes"
	"go-tinder-app/db"
)

func main() {
	db.InitDB()
	r := routes.InitRoutes()

	r.Run(":2929")
}
