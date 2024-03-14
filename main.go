package main

import (
	"assignment2-golang/database"
	"assignment2-golang/routes"
	"fmt"
)

var PORT = ":8080"

func main() {
	database.StartDB()
	fmt.Println("Listening Port .....")
	routes.StartServer().Run(PORT)
}