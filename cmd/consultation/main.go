package main

import "HeidiTask/internal"

func init() {
	// kafka topic initialization
	internal.CreateKafkaTopic()
}

func main() {
	app := App{}
	app.run()
}
