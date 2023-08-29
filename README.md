# golang
Homework Task 
Необходимо написать отдельный пакет, в котором будет реализован in-memory cache со следующими методами:

- `Set(key string, value interface{})` - запись значения `value` в кеш по ключу `key`
- `Get(key string)`
- `Delete(key)`

## Example

```go
func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId := cache.Get("userId")

	fmt.Println(userId)
}
```

## Solution
```go
package cashe

type Cashe struct {
	cashMap map[string]interface{}
}

func New() *Cashe {
	return &Cashe{
		cashMap: make(map[string]interface{}),
	}
}

func (c *Cashe) Set(key string, value interface{}) {
	c.cashMap[key] = value
}

func (c *Cashe) Get(key string) interface{} {
	return c.cashMap[key]
}

func (c *Cashe) Delete(key string) {
	delete(c.cashMap, key)
}
```
