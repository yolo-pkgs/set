# set

Because I can, that's why. And no, it's not thread-safe

```go
package main

import (
	"fmt"

	"github.com/yolo-pkgs/set"
)

func main() {
	s := set.New[int]()
	s.Add(0, 0, 1)
	s.Add(2)
	fmt.Println(s)
}
```
