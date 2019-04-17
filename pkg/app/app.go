package app

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	Port = ":9000"
)

var Logger, _ = zap.NewDevelopment()

type App struct {
	Router *mux.Router
}
