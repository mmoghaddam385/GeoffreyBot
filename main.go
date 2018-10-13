package main

import (
	"encoding/json"
	"io/ioutil"
  "log"
  "fmt"
  "net/http"
  "os"
)

func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request recieved!\n")
  fmt.Fprintln(w, "Hello World")
}

func recieveMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Message Recieved!!\n")
	body, err := ioutil.ReadAll(r.Body)

	if (err != nil) {
		fmt.Printf("Error reading request body: %v", err)
		return
	}

	bodyMap := make(map[string] string)

	json.Unmarshal(body, &bodyMap)

	fmt.Printf("Body: %v", bodyMap)
}

func main() {
  addr, err := determineListenAddress()
  if err != nil {
    log.Fatal(err)
  }

  http.HandleFunc("/", hello)
  http.HandleFunc("/message", recieveMessage)
  log.Printf("Listening on %s...\n", addr)
  if err := http.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}