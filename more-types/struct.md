# Struct
- Struct là 1 cấu trúc dữ liệu, struct khác object
```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
// Vertex là 1 kiểu dữ liệu tronng go
```
- Nhìn vào ví dụ bên dưới để thấy struct mang đầy đủ các thuộc tính của type 

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	a:= Vertex{1, 2}
	b := a // b COPY VALUE FORM a
  // nếu Vertex là 1 object sthì b sẽ trỏ đến vùng nhớ của a, nên struct không giống object
	b.X = 10
	fmt.Println(a.X,b.Y)
  
}
// 1 2
```

- Lưu ý: khi sử dụng struct thì tạo thành 1 cấu trúc dữ liệu có size băng tổng size của mỗi field

```go
type Employee struct {
  ID int  
  Name string
  Age int16
  Gender string
  Active bool
}
```

```go
package main

import (
	"fmt"
	"unsafe"
)

type Employee struct {
  ID int
  Name string
  Age int16
  Gender string
  Active bool
}

func main() {
  var e Employee
  fmt.Printf("Size of %T struct: %d bytes", e, unsafe.Sizeof(e))
}
// 56 bytes
```

[Link reference](https://perennialsky.medium.com/memory-allocation-for-a-struct-in-golang-b5057b8ccf23)
