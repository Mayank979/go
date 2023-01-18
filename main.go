package main

import (
	"fmt"
	"log"

	"github.com/Mayank979/go/router"
	"github.com/Mayank979/go/storage"
)

func main() {

	fmt.Println("Go Project Server")

	store, err := storage.NewPostgresStore()

	if err != nil {
		log.Fatal("error while setting up postgres")
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := router.RunAPIServer(":3000", store)

	server.Run()

}
