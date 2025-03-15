package persistence

type Response[T any] struct {
    Success bool
    Data []T
    ErrorMessage string
}
