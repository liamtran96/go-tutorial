# Flow control take note

### for loop
- for: k có dấu ngoặc tròn
- golang k có while: sử dụng for như while

```go
  package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

```

### if statement

- k có dấu ngoặc tròn như for
- short hand của if: vừa khai báo vừa khởi tạo 

```go
  package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

```

# Key word "defer" trong golang

- Định nghĩa: Các câu lệnh phía sau từ khoá defer sẽ đc thực thi ngay sau khi hàm kết thúc(return)

```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
// hello 
// world
```

## stacking defer

- Cơ chế stacking defer là khi gọi nhiêu defer cùng lúc thì thằng nào vào trc ra sau (firt in last out) 

```
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

//done 
// counting
// 9 => 0
```
