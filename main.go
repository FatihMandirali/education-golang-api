package main

import (
	"education.api/dbconnect"
	"education.api/handlers"
)

// @title User API
// @description User microservice server.
// @schemes http https
// @securityDefinitions.basic BasicAuth
func main() {
	dbconnect.DbInit()
	dbconnect.InitialMigration()
	handlers.Run()
}
