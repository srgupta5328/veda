package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/srgupta5328/veda/pkg/app"
)

func TestHealth(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:9000/", nil)
	res := httptest.NewRecorder()

	app.Health(res, req)

	got := res.Body.String()
	want := "Welcome to the Coin Market Application in Go"

	if got != want {
		t.Errorf("Got: '%s', Want: '%s'", got, want)
	}
}
