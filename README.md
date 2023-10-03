# SimpleCache
Simple map-based in-memory cache with Golang (Zhashkevych's course)

## Example:

```
package main

import (
	"fmt"

	"github.com/AngelicaNice/SimpleCache/cache"
)

func main() {
	cache := cache.New()
	cache.Set("key_int", 4)
	cache.Set("key_string", "some_string")
	fmt.Println(cache)
	cache.Delete("key_int")
	fmt.Println(cache)
	fmt.Println(cache.Get("some_string"))
	fmt.Println(cache.Get("key_string"))
}
```
