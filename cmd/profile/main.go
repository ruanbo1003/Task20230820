package main

import (
	"HeidiTask/internal"
	"fmt"
	"log"
)

func init() {
	// mysql initialization
	internal.MysqlMigrate()
}

func main() {
	fmt.Println("profile server")

	restfulApi()
}

func restfulApi() {
	router := initRouter()

	err := router.Run("0.0.0.0:10001")
	if err != nil {
		log.Fatal("gin run failed:", err)
	}
}
