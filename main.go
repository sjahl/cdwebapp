package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "0.0.9"

// Sysinfo accepts a hostname and version as strings and returns a formatted string
func Sysinfo(hostname, version string) string {
	return fmt.Sprintf("<h1>Hello</h1>\n<p>My hostname is: %s</p>\n<p>My version is: %s</p>\n", hostname, version)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()

		responseText := Sysinfo(hostname, version)
		fmt.Fprintf(w, responseText)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
