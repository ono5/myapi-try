package main

import (
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Worldi!"))
}

type application struct{}

func main() {
	app := &application{}
	app.routes()
}
