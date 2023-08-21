package main

import "HeidiTask/internal"

func init() {
	// mysql initialization
	internal.MysqlMigrate()

	// kafka topic initialization
	internal.CreateKafkaTopic()
}

func main() {
	app := App{}
	app.run()
}
