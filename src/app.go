package main

import (
	"fmt"
	"main/api"
	"main/persistence"
	"main/persistence/queries"
	"net/http"
)


func main() {
    fmt.Println("Starting...")

    db := persistence.Connect()
    defer db.Close()

    api.RegisterEndpoint("/db", func(w http.ResponseWriter, req *http.Request) {
        data := queries.GetSongNames(db)
        api.WriteToResponse[persistence.Response[queries.Named]](w, data)
    })
    api.Start()
}
