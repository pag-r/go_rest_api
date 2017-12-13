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

var Employees []Employee
//type Employees []Employee -> error, type employees is not an expression golang

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome in home..")
    fmt.Println("Endpoint hit: homePage")
}

func returnAllEmployees(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint hit: returnAllEmployees")
    json.NewEncoder(w).Encode(Employees)
}  

func returnSingleEmployee(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint hit: returnSingleEmployee")
    params := mux.Vars(r)
    for _, item := range Employees{
        if item.Id == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Employee{})
    return
}

// TO DO
// missing 'position' in curl


func createSingleEmployee(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint hit: createSingleEmployee")

    // curl -i -X POST -d '{"id":"4","title":"test","name":"test"}' localhost:9999/emp
    r.ParseForm()
    fmt.Println("form: ", r.Form)
    //map[{"id":"4","title":"test","name":"test","position":null}:[]]
    var employee Employee
    for key, _ := range r.Form {
         fmt.Println("key:",key)
        //key: {"id":"4","title":"test","name":"test","position":null}
        err := json.Unmarshal([]byte(key), &employee)
        if err != nil {
            log.Println(err.Error())
        }
    }
    Employees = append(Employees, Employee{Id: employee.Id, Title: employee.Title, Name: employee.Name})

    // params := mux.Vars(r)
    // var employee Employee
    // _ = json.NewDecoder(r.Body).Decode(&Employees)
    
    // employee.Id = params["id"]
    // employee.Title = params["title"]

    // Employees = append(Employees, Employee{})
    // json.NewEncoder(w).Encode(Employees)
}

//mux
//https://github.com/gorilla/mux
func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)

    Employees = append(Employees, Employee{Id:"1", Title:"spec1", Name:"test1", Position: []string{"admin1", "dev1"}})
    Employees = append(Employees, Employee{Id:"2", Title:"spec1", Name:"test2", Position: []string{"admin2", "ops1"}})

    // valid for: type Employees []Employee
    // employees := Employees{
    //     Employee{Id:"1", Title:"spec1", Name:"test1", Position: []string{"admin1", "dev1"}},
    //     Employee{Id:"2", Title:"spec1", Name:"test2", Position: []string{"admin2", "ops1"}},
    // }

    myRouter.HandleFunc("/", homePage).Methods("GET")
    myRouter.HandleFunc("/emp", returnAllEmployees).Methods("GET")
    myRouter.HandleFunc("/emp/{id}", returnSingleEmployee).Methods("GET")
    //myRouter.HandleFunc("/emp/{id}", createSingleEmployee).Methods("POST")
    myRouter.HandleFunc("/emp", createSingleEmployee).Methods("POST")
    log.Fatal(http.ListenAndServe(":9999", myRouter))
}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    handleRequests()
}
