package main

import (
	"github.com/pvtalbot/money/app"
	"github.com/pvtalbot/money/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app, cleanup := app.NewApplication()
	defer cleanup()

	server.RunHttpServer(app)
}
