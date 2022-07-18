# Cache
___

In-memory cache implementation, Go package
```go
type Cache interface {
	Set(key string, value any) error
	Get(key string) (any, error)
	Delete(key string) error
}
```
___
# example 1
```go
c := cache.NewCache()
```
___
# example 2
```go
c := cache.NewCache()
```
___