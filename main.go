package main

import (
	"back/routes"
	"back/services"
	"back/utility"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Read environment variable to set the hearing port
	port, presentPort := os.LookupEnv("APIPORT")

	// Read environment variable to set the postgres container IP
	postgresIp, presentIp := os.LookupEnv("DATABASEIP")

	if presentPort == false {
		port = "8000"
	}

	if presentIp == false {
		fmt.Println("Trying to connect to default database ip: 127.0.0.1")
		postgresIp = "127.0.0.1"
	}

	var db = utility.GetConnection(postgresIp)
	services.SetDB(db)
	var appRouter = routes.CreateRouter()

	log.Println("Listening on Port " + port)
	log.Fatal(http.ListenAndServe(":"+port, appRouter))
}
