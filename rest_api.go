package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

//str.115
type Employee struct {
    Id          int         `json:"id"`
    Title       string      `json:"title"`
    Name        string      `json:"name"`
    Position    []string    `json:"position"`
}
type Employees []Employee

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome in home..")
    fmt.Println("Endpoint hit: homePage")
}

//simple requests - not mux
// func handleRequests() {
//     http.HandleFunc("/", homePage)
//     http.HandleFunc("/emp", returnAllEmployees)
//     log.Fatal(http.ListenAndServe(":9999", nil))
// }

//mux
//https://github.com/gorilla/mux
func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/emp", returnAllEmployees)
    myRouter.HandleFunc("/employee/{key}", returnSingleEmployee)
    log.Fatal(http.ListenAndServe(":1234", myRouter))
}

func returnSingleEmployee(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]
    fmt.Println("Key: " + key)
    fmt.Println(vars)
}

func returnAllEmployees(w http.ResponseWriter, r *http.Request){
    employees := Employees{
        Employee{Title:"spec1", Name:"test1", Position: []string{"admin1", "dev1"}},
        Employee{Title:"spec1", Name:"test2", Position: []string{"admin2", "ops1"}},
    }
    fmt.Println("Endpoint hit: returnAllEmployees")
    json.NewEncoder(w).Encode(employees)
}    

func main() {
    // mux
    fmt.Println("Rest API v2.0 - Mux Routers")

    handleRequests()
}