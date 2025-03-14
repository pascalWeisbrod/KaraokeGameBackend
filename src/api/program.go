package api

import (
    "fmt"
    "net/http"
    "main/model"
    "encoding/json"
)

var users = model.NewUserList()

func Start() {
    http.HandleFunc("/hello", helloworld) 
    http.HandleFunc("/reset", registerUser)
    http.HandleFunc("/join", registerUser)
    
    http.ListenAndServe(":443", nil)
}

func registerUser(w http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    var user model.User
    err := decoder.Decode(&user)
    if err != nil {
        fmt.Fprintf(w, err.Error())
        return
    }

    users.Add(user)

    encoder := json.NewEncoder(w)
    encoder.Encode(users)
} 

func resetUserList(w http.ResponseWriter, req *http.Request) {
    users = model.NewUserList()
}

func helloworld(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Hello World")
}
