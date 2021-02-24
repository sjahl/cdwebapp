package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "0.0.1"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()

		responseText := `<h1>Hello</h1>
<p>My hostname is: %s</p>
<p>My version is: %s</p>
`
		fmt.Fprintf(w, responseText, hostname, version)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
