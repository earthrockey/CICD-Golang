package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/earthrockey/CICD-Golang/model"
)

func TestOpenServer(t *testing.T) {
	log.Print("The service is start on http://localhost" + getPort())
	go HandleRequest()
}

func TestGetAllBook(t *testing.T) {
	log.Println("Test Get All Book")
	r, err := http.Get("http://localhost:8888/api/get/allbook")
	if err != nil {
		log.Print(err)
		t.Errorf("Error http.Get(\"http://localhost:8888/api/get/allbook\"): %s", err)
	}
	defer r.Body.Close()
	var books []model.Book
	err = json.NewDecoder(r.Body).Decode(&books)
	if err != nil {
		log.Print(err)
		t.Errorf("Error json.NewDecoder(r.Body).Decode(&books): %s", err)
	}
	fmt.Println(books)
}
