package queries

import (
	"fmt"
	"database/sql"
	"main/model"
	"main/persistence"
)

type SongQueries struct {
    db *sql.DB
}

func InitiateSongService (dbConnection *sql.DB) SongQueries {
    return SongQueries{db: dbConnection}
}

func (s SongQueries) GetSongs() persistence.Response[model.Song] {
    rows, err := s.db.Query("select * from song;")
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

func (s SongQueries) PostSong(song model.Song) persistence.Response[model.Song] {
    result, err := s.db.Exec("insert into song (name, album, text) values (?, ?, ?)", song.Name, song.Album, song.Text)

    if err != nil {
        return persistence.Response[model.Song]{ Success: false, ErrorMessage: "Something went wrong trying to access the database." }
    }

    resultID, err := result.LastInsertId()
    song.ID = int(resultID)

    if err != nil {
        return persistence.Response[model.Song]{ Success: false, ErrorMessage: "Something went wrong trying to access the database." }
    }

    return persistence.Response[model.Song]{ Success: true, Data: []model.Song{song} }
}
