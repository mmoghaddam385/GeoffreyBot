package endpoints

import (
	"fmt"
	"net/http"
)

func fallbackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request handled by fallback handler; path - %v\n", r.URL.Path)

	fmt.Fprintln(w, "You've reached Geoffrey, please leave a message after the beep...")
}