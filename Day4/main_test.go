
package main

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"freshers_bootcamp/Day4/Config"
	"freshers_bootcamp/Day4/Routes"
	"strings"
	"testing"
)

var (
	err       error
	APIRouter *gin.Engine
)

func TestDatabaseConnection(t *testing.T) {
	APIRouter = Router.SetupRouter()
	Config.DB, err = Config.ConnectDB()
	if err != nil {
		panic(err)
	}
}
func TestCreateUser(t *testing.T) {
	jsonParam := `{"name":"test User"}`

	req, err := http.NewRequest("POST", "/user", strings.NewReader(string(jsonParam)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// Router := Router.SetupRouter()
	handler := http.Handler(APIRouter)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestCreateProduct(t *testing.T) {
	jsonParam := `{"name":" product","quantity":100, "price" : 500}`

	req, err := http.NewRequest("POST", "/product", strings.NewReader(string(jsonParam)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.Handler(APIRouter)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestCreateOrder(t *testing.T) {

	jsonParam := `{"ProductID":1, "UserID":1}`

	req, err := http.NewRequest("POST", "/order", strings.NewReader(string(jsonParam)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.Handler(APIRouter)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestGetProduct(t *testing.T) {
	req, err := http.NewRequest("GET", "/product/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.Handler(APIRouter)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}