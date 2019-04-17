package app_test

import (
	"testing"

	"github.com/srgupta5328/veda/pkg/app"
)

func TestNewRouter(t *testing.T) {
	t.Run("Initializing the router", func(t *testing.T) {
		got := app.NewRouter()
		if got == nil {
			t.Log("No Error Starting the router")
		}
	})

}
