package gojicorsmiddleware_test

import (
	"fmt"
	"github.com/rpagliuca/gojicorsmiddleware"
	"goji.io"
	"goji.io/pat"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/my-resource"), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "success")
	})
	mux.Use(gojicorsmiddleware.CorsMiddlewareHandler)

	server := httptest.NewServer(mux)
	client := server.Client()

	req, _ := http.NewRequest("GET", server.URL+"/my-resource", nil)
	resp, _ := client.Do(req)

	if resp.Header.Get("Access-Control-Allow-Headers") != "*" {
		t.Error("Access-Control-Allow-Headers not defined correctly")
	}

	if resp.Header.Get("Access-Control-Allow-Origin") != "*" {
		t.Error("Access-Control-Allow-Origin not defined correctly")
	}

	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Error("Error reading response body")
	}

	if string(contents) != "success" {
		t.Error("Inner handler returned unexpected response")
	}
}
