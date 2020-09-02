// bool
// 数字类型
// int8, int16, int32, int64, int
// uint8, uint16, uint32, uint64, uint
// float32, float64
// complex64, complex128
// byte
// rune
// byte 是 uint8 的别名。
// rune 是 int32 的别名。
// string;
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))   // a 的类型和大小  注意 用Printf来输出格式
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) // b 的类型和大小
	c, d := 5.67, 8.97
	fmt.Printf("\ntype of c %T d %T\n", c, d)
	sum := c + d
	diff := c - d
	fmt.Println("\nsum", sum, "diff", diff)
	no1, no2 := 56, 89
	fmt.Println("\nsum", no1+no2, "diff", no1-no2)
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("complexsum:", cadd)
	cmul := c1 * c2
	fmt.Println("complexmulti:", cmul)
	first := "Peng"
	last := "Three"
	name := first + " " + last
	fmt.Println("My name is", name)
	i := 55                //int
	j := 67.8              //float64
	sums := float64(i) + j //不允许 int + float64 不强制转换则报错
	fmt.Println(sums)
}
