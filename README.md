Installing GO
---
1. brew install go
2. ```export PATH=${HOME}/go/bin:$PATH``` add to .zshrc
3. ```source ~/.zshrc```


- VSCode for IDE
- [Standard GO Libraries](https://golang.org/pkg/)
- Installing other packages:
- Learning Go: [A tour of GO](https://tour.golang.org)

1. main package is executable and need main func
2. variable declaration, type at the last: var a int, var b int and  var x, y int
3. function return  func add(x,y int) int {}
3. func multiple return: function split(a int) (x, y) int {} also call naked func. returns x and y
4. variable with intializer: var i, j int = 1, 2
5. Short var declaration l:=40 , available within the func
6. variable declarations may be "factored" into blocks, as with import statements:
```go
    import (
	"fmt"
	"math/cmplx"
    )

    var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
    )
```
7.  default values for variables: 0 for numeric type, false for bool and "" for string.
8.




