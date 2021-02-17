package commands

import (
  "geoffrey/endpoints"

  "log"
  "fmt"
  "net/http"
  "os"
)

type ServerCommand struct{}

func (*ServerCommand) Name() string { return "server" }
func (*ServerCommand) Usage() string { return "server - Start the Geoffrey server on port $PORT"}

func (*ServerCommand) Execute(args []string) int {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	registerHandlers()

	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
	
	return 0
}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	
	return ":" + port, nil
}

// registerHandlers loops over all endpoints in the registry and tells 
// the server to handle them
func registerHandlers() {
	fmt.Println("Registering Endpoints...")

	for path, handler := range endpoints.GetAllEndpoints() {
		fmt.Printf("Registering Handler for %v\n", path)
		http.HandleFunc(path, handler)
	}
}
