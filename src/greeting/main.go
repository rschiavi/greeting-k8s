package main

import (
	"fmt"
	"log"
	"net/http"
)

func greeting(texto string) string {
	return fmt.Sprintf("<b>%s</b>", texto)
}

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, _ *http.Request) {
		res.Write([]byte(greeting("Code.education Rocks!")))
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
