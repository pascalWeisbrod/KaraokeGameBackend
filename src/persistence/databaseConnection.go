package persistence

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect(user string, pw string, addr string, db_name string) *sql.DB {
    // Capture connection properties.
    cfg := mysql.Config{
        User:   user,
        Passwd: pw,
        Net:    "tcp",
        Addr:   addr,
        DBName: db_name,
        AllowNativePasswords: true,
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Database Connected!")
    return db
}
