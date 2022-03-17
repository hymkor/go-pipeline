Pipeline Supporter
==================

[example](./example/example.go) :

With Go 1.18 Generics, any type's channel can be used for pipeline.

```go
package main

import (
    "fmt"

    "github.com/hymkor/go-pipeline"
)

func main() {
    fmt.Println("Start")
    pipeline.Run(func(in, out chan int) {
        for i := 0; i < 10; i++ {
            out <- i
        }
    }, func(in, out chan int) {
        for value := range in {
            out <- value + 1
        }
    }, func(in, out chan int) {
        for value := range in {
            fmt.Printf("%d\n", value)
        }
    })
    fmt.Println("Done.")
}
```
