package endpoint

type IEndpoint[T any] interface {
    Init()
    Path() string
    Get() []T 
    Post(item T) T
    Put(item T) T
    Delete(item T) bool
}
