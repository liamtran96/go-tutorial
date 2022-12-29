# Flow control take note

### for loop
- for: k có dấu ngoặc tròn
- golang k có while: sử dụng for như while

```
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

```
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