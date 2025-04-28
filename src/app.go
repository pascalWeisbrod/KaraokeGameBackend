package main

import (
	"fmt"
	"main/api"
	"main/api/endpoint/endpoints"
	"main/persistence"
	"main/persistence/queries"
	"os"
)

func main() {
	fmt.Println("Starting...")

	db_user := os.Getenv("KARAOKE_DB_USER")
	db_pass := os.Getenv("KARAOKE_DB_PW")
	db_addr := os.Getenv("KARAOKE_DB_ADDR")
	db_name := os.Getenv("KARAOKE_DB_NAME")
	db := persistence.Connect(db_user, db_pass, db_addr, db_name)
	defer db.Close()

    song_service := queries.InitiateSongService(db)

	api_port := os.Getenv("KARAOKE_API_PORT")
	server := api.RequestServer(api_port)
    api.RegisterIEndpoint(&server, endpoints.SongEndpoint{Songs: song_service})
    api.RegisterEndpoint(&server, "/get-karaoke-file", endpoints.GetFile)
    api.RegisterEndpoint(&server, "/post-karaoke-file", endpoints.PostFile)
	server.Start()
}
