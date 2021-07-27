package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"freshers_bootcamp/Day3/Config"
	"freshers_bootcamp/Day3/Controllers"
	"freshers_bootcamp/Day3/Models"
	"freshers_bootcamp/Day3/Routes"

	"github.com/jinzhu/gorm"
)


func TestGet(t *testing.T) {
	//SQL Connection using GORM
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()

	//Setting the router
	router := Routes.SetupRouter()
	router.GET("/user-api/user/", Controllers.GetUsers)

	//Get request
	request, _ := http.NewRequest("GET", "/user-api/user/", nil)

	//Recording the response
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedOutput := `[{"id":1,"name":"Guguloth","LastName":"Tarun","DOB":"03-11-1999","address":"Hyderabad","subject":"Mechanical","marks":78},{"id":2,"name":"Guguloth","LastName":"Varun","DOB":"01-01-2001","address":"Hyderabad","subject":"computers","marks":75}]`
	if response.Body.String() != expectedOutput {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expectedOutput)
	}

}

func TestPost(t *testing.T) {
	//SQL Connection using GORM
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	Config.DB.AutoMigrate(&Models.User{})


	//Setting the router
	router := Routes.SetupRouter()
	router.POST("/user-api/user/",Controllers.CreateUser)

	//send request
	newStud := Models.User{
		Name: "Test",
		LastName: "Check",
		DOB: "xyz",
		Address: "location",
		Subject: "Maths",
		Marks: 85,
	}

	responseBody,_ := json.Marshal(newStud)
	req, _ := http.NewRequest("POST", "/user-api/user/", bytes.NewBuffer([]byte(responseBody)))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedOutput := `{"id":144,"name":"Test","LastName":"Check","DOB":"xyz","address":"location","subject":"Maths","marks":85}`
	if response.Body.String() != expectedOutput {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expectedOutput)
	}

}

func TestPut(t *testing.T) {
	//SQL database using GORM
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	Config.DB.AutoMigrate(&Models.User{})

	//setting up router
	router := Routes.SetupRouter()
	router.PUT("/user-api/user/1",Controllers.UpdateUser)

	//send request
	newStudent := Models.User{
		LastName: "P",
		Marks:     85,
	}

	responseBody,_ := json.Marshal(newStudent)
	req, _ := http.NewRequest("PUT", "/user-api/user/1/", bytes.NewBuffer([]byte(responseBody)))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	expectedOutput := `{"id":1,"name":"Raajitha","LastName":"Potala","DOB":"22-03-2000","address":"Visakhapatnam","subject":"Communication","marks":85}`
	if response.Body.String() != expectedOutput {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expectedOutput)
	}
}

func TestDelete(t *testing.T) {
	//SQL Connection using GORM
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()

	//setup router
	router := Routes.SetupRouter()
	router.DELETE("/user-api/user/2/",Controllers.DeleteUser )

	//Get request
	req, _ := http.NewRequest("DELETE", "/user-api/user/21/", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	//checking test case
	if response.Code != 307 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Code, 200)
	}
}