package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	//"encoding/json"
    //models "api-rest-go/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
    repository "api-rest-go/repository"
 )

 var db *gorm.DB

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", GetHome).Methods("GET")
	router.HandleFunc("/people", GetPeople).Methods("GET")
    router.HandleFunc("/person/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/person/{id}", CreatePerson).Methods("POST")
    router.HandleFunc("/person/{id}", DeletePerson).Methods("DELETE")

    db = repository.OpenConnection()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"Está é a home. Leia o README para saber quais são as requisições possíveis.")
}

func GetPeople(w http.ResponseWriter, r *http.Request) {}

func GetPerson(w http.ResponseWriter, r *http.Request) {}

func CreatePerson(w http.ResponseWriter, r *http.Request) {}

func DeletePerson(w http.ResponseWriter, r *http.Request) {}