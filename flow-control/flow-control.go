package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	// sum := 0; 
	// for i:=0; i < 10 ; i++ {
	// 	sum += i 
	// }
	// defer fmt.Printf("Hàm return thì chạy defer")
	// fmt.Println(sum)
	// fmt.Println(sqrt(2), sqrt(-4))

	// fmt.Println(
	// 	pow(3, 2, 10),
	// 	pow(3, 3, 20),
	// )
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Printf("Hàm return thì chạy defer")
	fmt.Println("counting")

	

	fmt.Println("done")

}