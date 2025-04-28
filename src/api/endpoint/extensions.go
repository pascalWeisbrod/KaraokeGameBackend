package endpoint

import (
    "net/http"
    "encoding/json"
)

type error struct {
    Message string
}

func WriteToResponse[T any](w http.ResponseWriter, data any) {
    encoder := json.NewEncoder(w)
    encoder.Encode(data)
}

func TryReadFromRequest[T any](req *http.Request) (bool, T) {
    decoder := json.NewDecoder(req.Body)
    var item T
    err := decoder.Decode(&item) 
    if err != nil {
        var empty T
        return false, empty
    }
    return true, item
}

func RouteRequest[T any](ep IEndpoint[T]) func(w http.ResponseWriter, req *http.Request) {
    return func(w http.ResponseWriter, req *http.Request) {
        success, item := TryReadFromRequest[T](req)
        if req.Method != http.MethodGet && !success {
            WriteToResponse[error](w, error{Message: "Something went wrong!"})
            return 
        }

        switch req.Method {
        case http.MethodGet:
            WriteToResponse[T](w, ep.Get())
        case http.MethodPost:
            WriteToResponse[T](w, ep.Post(item))
        case http.MethodPut:
            WriteToResponse[T](w, ep.Put(item))
        case http.MethodDelete:
            WriteToResponse[T](w, ep.Delete(item))
        }
    }
}
