package main

import (
	"fmt"
	"main/api"
	"main/model"
	"main/persistence"
	"main/persistence/queries"
	"net/http"
)


func main() {
    fmt.Println("Starting...")

    db := persistence.Connect()
    defer db.Close()

    api.RegisterEndpoint("/songs", func(w http.ResponseWriter, req *http.Request) {
        data := queries.GetSongs(db)
        api.WriteToResponse[persistence.Response[model.Song]](w, data)
    })
    api.Start()
}
