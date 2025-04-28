package api

import (
    "net/http"
    "main/api/endpoint"
)

type Server struct {
    port string
    running bool
    paths []string
}

func RequestServer(port string) Server {
    s := Server{ port: port, running: false, paths: make([]string, 0)}
    return s
}

func (s Server) Start() {
    if s.running {
        panic("Server was already started. Cannot start.")
    }
    s.running = true

    http.ListenAndServe(s.port, nil)
}

func RegisterEndpoint(s *Server, path string, ep func(w http.ResponseWriter, req *http.Request)) {
    if s.running {
        panic("Server is already running. Cannot register Endpoints.")
    }
    http.HandleFunc(path, ep)
}

func RegisterIEndpoint[T any](s *Server, ep endpoint.IEndpoint[T]) {
    if s.running {
        panic("Server is already running. Cannot register Endpoints.")
    }
    for _, p := range s.paths {
        if p == ep.Path() {
            panic("Path already exists : " + p)
        }
    }

    http.HandleFunc(ep.Path(), endpoint.RouteRequest(ep))
}
