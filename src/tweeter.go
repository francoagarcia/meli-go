package main

import (
	"fmt"
	"os"

	"github.com/francoagarcia/meli-go/src/endpoints"
	"github.com/francoagarcia/meli-go/src/service"
)

func main() {

	mode := os.Args[1]

	tweetManager := service.NewTweetManager()
	manager := endpoints.NewManager(tweetManager)

	if mode == "rest" {
		restEndpoint := endpoints.NewRest(manager)
		restEndpoint.Execute()
	} else if mode == "cmd" {
		cmdEndpoint := endpoints.NewCmd(manager)
		cmdEndpoint.Execute()
	} else {
		fmt.Printf("Opcion de ejecución no soportada %s", mode)
	}
}
