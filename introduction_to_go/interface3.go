package main

import (
	"fmt"
	"io"
	"net/http"
)

// START OMIT
func handle(w io.Writer, msg string) {
	fmt.Fprintln(w, msg)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handle(w, r.URL.Path[1:]) // HL
	})
  
	fmt.Println("start listening on port 4000")
	http.ListenAndServe(":4000", nil)
}

// END OMIT
