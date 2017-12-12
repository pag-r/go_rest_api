package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

//str.115
type Employee struct {
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
func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/emp", returnAllEmployees)
    log.Fatal(http.ListenAndServe(":9999", nil))
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
    handleRequests()
}