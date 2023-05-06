package main

import (
	"fmt"
	"net"
	"net/http"

	"os"

	Utils "app/utils"

	"github.com/gorilla/mux"
	Hermes "github.com/realTristan/Hermes"
	"github.com/realTristan/Hermes/server/listener"
	"github.com/realTristan/Hermes/server/routes"
)

// Main function
func main() {
	// Update hermes
	Utils.UpdateHermes()

	// Verify that the user is trying to serve the cache
	if os.Args[1] != "serve" {
		panic("incorrect usage. example: ./hermes serve")
	}

	// Get the arg data
	var args, err = Utils.GetArgData(os.Args)
	if err != nil {
		panic(err)
	}

	// Get the port and json file
	var cache *Hermes.Cache = Hermes.InitCache()

	// Initialize a new listener
	if l, err := listener.New(args.Port()); err != nil {
		panic(err)
	} else {
		// Establish a new gorilla mux router
		var router *mux.Router = mux.NewRouter()

		// Set and handle router endpoints
		routes.Set(router, cache)
		http.Handle("/", router)

		// Print the serving port
		var port int = l.Addr().(*net.TCPAddr).Port
		Utils.PrintLogoWithPort(port)

		// Serve the listener
		if err := http.Serve(l, nil); err != nil {
			fmt.Println(err)
		}
	}
}