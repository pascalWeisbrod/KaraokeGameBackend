package queries

import (
    "database/sql"
    "main/persistence"
)

type Named struct { Name string }

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
