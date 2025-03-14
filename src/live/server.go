package live

import (
    "net"
)

func Connect() {
    ln, err := net.Listen("udp", ":8080") 
    if err != nil {
        panic(err)
    }
    defer ln.Close()
    
    for {
        conn, err := ln.Accept()
        if err != nil {
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(connection net.Conn) {
}
