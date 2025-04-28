package endpoints

import (
	"main/model"
	"main/persistence/queries"
)

type SongEndpoint struct {
    Songs queries.SongQueries
}

func (e SongEndpoint) Init() {
}

func (e SongEndpoint) Path() string {
    return "/songs"
}

func (e SongEndpoint) Get() []model.Song {
    query := e.Songs.GetSongs()
    if !query.Success {
        var empty []model.Song
        return empty
    }
    return query.Data
}

func (e SongEndpoint) Post(item model.Song) model.Song {
    query := e.Songs.PostSong(item)
    if !query.Success {
        var empty model.Song
        return empty
    }
    return query.Data[0]
}

func (e SongEndpoint) Put(item model.Song) model.Song {
    panic("")
}

func (e SongEndpoint) Delete(item model.Song) bool {
    return false
}
