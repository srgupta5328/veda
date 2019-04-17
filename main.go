package main

import (
	"log"
	"net/http"

	"github.com/srgupta5328/veda/pkg/app"
)

func main() {
	a := app.App{}
	a.Router = app.NewRouter()
	app.Logger.Info("Spinning Up Applications on K8s")
	log.Fatal(http.ListenAndServe(app.Port, a.Router))
}
