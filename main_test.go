package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.init()
	code := m.Run()
	os.Exit(code)

}

func TestIndex(t *testing.T) {

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	a.router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Bad HTTP return code : %d", rr.Code)
	}
	var m map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)
	if m["message"] != "Hello from go api server" {
		t.Errorf("Bad response message: %v", m["message"])
	}
}

func TestHealth(t *testing.T) {

	req, _ := http.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()
	a.router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Bad HTTP return code : %d", rr.Code)
	}
	var m map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)
	if m["health"] != "OK" {
		t.Errorf("Bad response message: %v", m["message"])
	}
}

func TestLiveness(t *testing.T) {

	req, _ := http.NewRequest("GET", "/liveness", nil)
	rr := httptest.NewRecorder()
	a.router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Bad HTTP return code : %d", rr.Code)
	}
	var m map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)
	if m["status"] != "UP" {
		t.Errorf("Bad response message: %v", m["message"])
	}
}
