package main

import (
	"mytesting/restapi/internal/database"
	"mytesting/restapi/internal/transport"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	app := transport.Setup()
	database.InitDatabase()
	app.Listen(3000)
}
