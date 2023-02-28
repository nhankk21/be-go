package main

import (
	"godk/core"
	"godk/router"
)

func main() {
	server := core.NewServer()
	if server == nil {
		panic("Error while start new server!")
	}
	router.SetupRouter(*server)
	err := server.Expose(":8080")
	if err != nil {
		panic(err)
	}

}
