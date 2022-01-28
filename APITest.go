package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEntries(t *testing.T) {
	req, err := http.NewRequest("GET", "/Desafio/1.0.0/project/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SearchProjects)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	//expected := `[{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb@g.com","phone_number":"0987654321"},{"id":2,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"},{"id":6,"first_name":"FirstNameSample","last_name":"LastNameSample","email_address":"lr@gmail.com","phone_number":"1111111111"}]`
	if len(rr.Body.String()) == 0 {
		t.Errorf("handler returned unexpected body: got-> %v ",
			rr.Body.String())
	}
}

func TestGetProjectByName(t *testing.T) {

	req, err := http.NewRequest("GET", "/Desafio/1.0.0/project/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("name", "prj1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ProjectProjectnameGet)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"Name":        "prj1","DisplayName": "Project_1",
		Description: " Projeto da equipe 1",
		 "CreateTime":  time.Now(),
		"UpdateTime":  time.Now(),
		"Apis":        []ApiItem{},
	}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got-> %v ", rr.Body.String())
	}
}

func TestGetEntryByNameNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/Desafio/1.0.0/project/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("name", "prj23")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ProjectProjectnameGet)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestCreateEntry(t *testing.T) {

	var jsonStr = []byte(`{
		Name:        "prj30",
		DisplayName: "Project_3",
		Description: " Projeto da equipe 3"	}`)

	req, err := http.NewRequest("POST", "/entry", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddProject)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{
		Name:        "prj30",
		DisplayName: "Project_3",
		Description: " Projeto da equipe 3"	}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestDeleteProject(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/Desafio/1.0.0/project/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("name", "prj2")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ProjectProjectnameDelete)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":4,"first_name":"xyz change","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestEditProject(t *testing.T) {

	var jsonStr = []byte(`{"name": "prj3","DisplayName": "Project_3_T","Description": " Projeto da equipe 3_Teste",}`)

	req, err := http.NewRequest("PUT", "localhost:8080/Desafio/1.0.0/project/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ProjectProjectnamePut)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"name": "prj3","DisplayName": "Project_3_T","Description": " Projeto da equipe 3_Teste",}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
