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
    Id          string      `json:"id"`
    Title       string      `json:"title"`
    Name        string      `json:"name"`
    Position    []string    `json:"position"`
}

var employees []Employee
//type employees []Employee -> error, type employees is not an expression golang

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome in home..")
    fmt.Println("Endpoint hit: homePage")
}

func returnAllEmployees(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint hit: returnAllEmployees")
    json.NewEncoder(w).Encode(employees)
}  

func returnSingleEmployee(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint hit: returnSingleEmployee")
    params := mux.Vars(r)
    for _, item := range employees{
        if item.Id == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Employee{})
    return
}

//mux
//https://github.com/gorilla/mux
func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)

    employees = append(employees, Employee{Id:"1", Title:"spec1", Name:"test1", Position: []string{"admin1", "dev1"}})
    employees = append(employees, Employee{Id:"2", Title:"spec1", Name:"test2", Position: []string{"admin2", "ops1"}})

    // valid for: type employees []Employee
    // employees := employees{
    //     Employee{Id:"1", Title:"spec1", Name:"test1", Position: []string{"admin1", "dev1"}},
    //     Employee{Id:"2", Title:"spec1", Name:"test2", Position: []string{"admin2", "ops1"}},
    // }

    myRouter.HandleFunc("/", homePage).Methods("GET")
    myRouter.HandleFunc("/emp", returnAllEmployees).Methods("GET")
    myRouter.HandleFunc("/emp/{id}", returnSingleEmployee).Methods("GET")
    log.Fatal(http.ListenAndServe(":9999", myRouter))
}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    handleRequests()
}
