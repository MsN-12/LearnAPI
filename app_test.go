package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	err := a.Initialise(DBUser, DBPassword, "test")
	if err != nil {
		log.Fatalln("initialise failed")
	}
	createTable()
	m.Run()
}
func TestGetProduct(t *testing.T) {
	clearTable()
	addProduct("keyboard", 110, 500)
	request, _ := http.NewRequest("GET", "/product/1", nil)
	rsp := sendRequest(request)
	checkStatusCode(t, http.StatusOK, rsp.Code)

}
func TestCreateProduct(t *testing.T) {
	clearTable()
	var product = []byte(`{"name": "chair", "quantity":1, "price":100}`)
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(product))
	req.Header.Set("Content-Type", "application/json")
	rsp := sendRequest(req)
	checkStatusCode(t, http.StatusCreated, rsp.Code)

	var m map[string]interface{}
	json.Unmarshal(rsp.Body.Bytes(), &m)
	if m["name"] != "chair" {
		t.Errorf("Expected name : %v, Got: %v", "chair", m["name"])
	}
	if m["quantity"] != 1.0 {
		t.Errorf("Expected quantity : %v, Got: %v", 1, m["quantity"])
	}
}

func createTable() {
	createTableQuery := `CREATE TABLE IF NOT EXISTS products(
    id int not null AUTO_INCREMENT,
    name varchar(255) not null,
    quantity int,
    price float(10,7),
    PRIMARY KEY (id)
);`
	_, err := a.DB.Exec(createTableQuery)
	if err != nil {
		log.Fatalln(err)
	}
}
func clearTable() {
	a.DB.Exec("DELETE from products")
	a.DB.Exec("alter table products AUTO_INCREMENT=1")
}
func addProduct(name string, quantity int, price float64) {
	query := fmt.Sprintf("INSERT into products(name, quantity, price) VALUES('%v',%v,%v)", name, quantity, price)
	a.DB.Exec(query)
}
func sendRequest(request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, request)
	return recorder
}
func checkStatusCode(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Expected status: %v, Recived: %v", expected, actual)
	}
}
