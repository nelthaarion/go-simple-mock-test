package controllers

import (
	"errors"
	"io/ioutil"
	"net/http/httptest"
	"testProject/services"
	"testing"
)

type RouteServiceMock struct {
	ExpectedReturn func() ([]byte, error)
}

func (rsm RouteServiceMock) HandlePing() ([]byte, error) {
	return rsm.ExpectedReturn()
}

func TestPingWithoutError(t *testing.T) {
	rsm := RouteServiceMock{}
	rsm.ExpectedReturn = func() ([]byte, error) {
		return []byte("PONG"), nil
	}
	services.RouteServiceInstance = rsm
	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	PingHandler(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == 500 {
		t.Error(err.Error())
	}
	if string(data) != "PONG" {
		t.Error(err.Error())
	}
}

func TestPingWithError(t *testing.T) {
	rsm := RouteServiceMock{}
	rsm.ExpectedReturn = func() ([]byte, error) {
		return []byte("PONG"), errors.New("JOBS DONE")
	}
	services.RouteServiceInstance = rsm
	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	PingHandler(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	if resp.StatusCode != 500 {
		t.Error("error is null")
	}
}
