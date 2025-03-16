package queries

import (
	"database/sql"
	"fmt"
	"main/model"
	"main/persistence"
)

type Named struct { Name string }

func GetSongs(db *sql.DB) persistence.Response[model.Song] {
    rows, err := db.Query("select * from song;")
    defer rows.Close()

    if err != nil {
        return persistence.Response[model.Song]{ Success: false, ErrorMessage: "Something went wrong trying to access the database." }
    }
    
    var songs []model.Song
    for rows.Next() {
        var song model.Song
        if err := rows.Scan(&song.ID, &song.Name, &song.Album, &song.Text); err != nil {
            fmt.Println(err.Error())
            return persistence.Response[model.Song]{ Success: false, ErrorMessage: "Something went wrong jsonifying" }
        }
        songs = append(songs, song)
    }
    return persistence.Response[model.Song]{ Success: true, Data: songs }
}
func GetSongNames(db *sql.DB) persistence.Response[Named] {
    rows, err := db.Query("select name from song;")
    defer rows.Close()

    if err != nil {
        return persistence.Response[Named]{ Success: false, ErrorMessage: err.Error() }
    }
    
    var data []Named
    for rows.Next() {
        var named Named
        if err := rows.Scan(&named.Name); err != nil {
            return persistence.Response[Named]{ Success: false, ErrorMessage: err.Error() }
        }
        data = append(data, named)
    }
    return persistence.Response[Named]{ Success: true, Data: data }
}
