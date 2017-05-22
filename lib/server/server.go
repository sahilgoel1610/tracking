package server

import "net/http"
import "fmt"

func Start() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/letsee", letseeHandler)

	http.ListenAndServe(":8082", nil)
}


func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home path")
}

func letseeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "LetSee")
}