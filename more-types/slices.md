#Slices 
-  Là 1 mảng đc cắt ra từ 1 mảng khác đc gọi là slices. Và slices là 1 cấu trúc dữ liệu

```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

//[3 5 7]
```
- [bỏ đi bao nhiêu phần tử : lấy bao nhiêu phần tử]
- [:] => để trống thì mặc định là lấy hết

```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[:]
	fmt.Println(s)
}

//[2 3 5 7 11 13]

```

### Bản chất của slice là 1 con trỏ đến mảng mà nó cắt

```go 
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
  s[0] = 200
	fmt.Println(s)
	fmt.Println(primes)
}

//[200 5 7]
//[2 200 5 7 11 13]

```

#Làm sao để copy 1 slice

- Cắt ra 1 slices mới với lệnh make() sau đó for qua nó để gắn vào