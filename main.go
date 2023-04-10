package main

import (
	"os"
	"outdoorsy/app"
)

func main() {
	var apiPort, dbConnectionUrl string
	if apiPort = os.Getenv("API_PORT"); apiPort == "" {
		apiPort = "1212"
	}
	if dbConnectionUrl = os.Getenv("DB_CONNECTION_URL"); dbConnectionUrl == "" {
		dbConnectionUrl = "postgresql://postgres:root@localhost/postgres?sslmode=disable"
	}
	app.Initialize(apiPort, dbConnectionUrl)
}
