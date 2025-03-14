package main

import (
	"fmt"
	"main/api"
)


func main() {
    fmt.Println("Starting...")
    go api.Start()
}
