package main

import (
	"fmt"
	"main/api"
    "main/persistence"
    "net/http"
)


func main() {
    fmt.Println("Starting...")

    db := persistence.Connect()
    defer db.Close()

    api.RegisterEndpoint("/db", func(w http.ResponseWriter, req *http.Request) {
        rows, err := db.Query("select * from account;")
        if err != nil {
            api.WriteToResponse(w, err.Error())
            return
        }
        api.WriteToResponse(w, rows)
    })
    api.Start()
}
