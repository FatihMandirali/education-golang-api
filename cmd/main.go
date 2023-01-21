package main

import (
	"education.api/dbconnect"
	"education.api/handlers"
)

func main() {
	dbconnect.DbInit()
	dbconnect.InitialMigration()
	handlers.Run()
}
