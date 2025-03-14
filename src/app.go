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
        rows, err := db.Query("select name from song;")
        defer rows.Close()

        if err != nil {
            api.WriteToResponse(w, err.Error())
            return
        }

        type Named struct { Name string }
        var data []Named

        for rows.Next() {

            var named Named
            if err := rows.Scan(&named.Name); err != nil { 
                api.WriteToResponse(w, err.Error())
                return 
            }
            data = append(data, named)
        }

        api.WriteToResponse(w, data)
    })
    api.Start()
}
