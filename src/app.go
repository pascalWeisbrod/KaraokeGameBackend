package main

import (
	"fmt"
    "main/persistence"
)


func main() {
    fmt.Println("Starting...")
    db := persistence.Connect()
    db.Query("show tables;")
    defer db.Close()
}
