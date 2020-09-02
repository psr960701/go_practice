package main

import (
	"fmt"
	"math"
)

func main() {
	var age int // 变量声明 初始值为0
	fmt.Println("my age is", age)
	age = 29 // 赋值
	fmt.Println("my age is", age)
	age = 54 // 赋值
	fmt.Println("my new age is", age)
	var number = 233 //声明变量并初始化
	fmt.Println("intresting number is", number)
	var name = "three" //自动推测类型为字符串
	var nowage = 24
	fmt.Println(name, "is", nowage)
	var height, weight = 180, 160 //如果要初始化就得两个一起初始化
	fmt.Println("height is", height, "weight is", weight)
	var (
		newName  = "peng three"
		heartAge = 18
		exWeight = 150
	)
	fmt.Println(newName, heartAge, exWeight)
	//简短声明
	shortA, shortB, shortName := 10, 20, "tPeng" //必须都拥有初始值
	fmt.Println(shortA, shortB, shortName)       //声明后必须使用 否则报错

	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("minimum value is ", c)
	strAge := 29
	// strAge="test" //不允许这么做 cannot use "test" (type untyped string) as type int in assignment
	fmt.Println("strToAge", strAge)
}
