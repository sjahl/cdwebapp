package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const version = "0.0.26"

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

// Sysinfo accepts a hostname and version as strings and returns a formatted string
func Sysinfo(hostname, version string) string {
	return fmt.Sprintf("<h1>Hello</h1>\n<p>My hostname is: %s</p>\n<p>My version is: %s</p>\n", hostname, version)
}

func main() {

	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	InfoLogger.Println("Starting up...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()

		if strings.HasSuffix(r.URL.Path, "favicon.ico") {
			ErrorLogger.Println("Favicon requested: stop asking for my favicon!")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		InfoLogger.Println("Favicon not requested: Genetrating response.")
		responseText := Sysinfo(hostname, version)
		fmt.Fprint(w, responseText)
		InfoLogger.Println("Rendered page. finishing.")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
