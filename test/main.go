package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	// 函数式编程,直接传递一个匿名函数
	ReceiveM(func(r int) int {
		return r * r
	}, 3)
	SumRange(1, 2, 3)
}

// ReceiveM 函数式编程,接受一个函数类型执行这个函数
func ReceiveM(m func(r int) int, r int) {
	fmt.Printf("ReceiveM exec the func name is %s.\n", runtime.FuncForPC(reflect.ValueOf(m).Pointer()).Name())
	fmt.Printf("ReceiveM exec the func had a result return %v\n", m(r))
}

// SumRange 循环加运算可变参数
func SumRange(nums ...int) int {
	tmpAdd := 0
	for i := 0; i < len(nums); i++ {
		// also below
	}
	for i, v := range nums {
		fmt.Printf("SumRange current index %v\n", i)
		//tmpAdd+=nums[i]
		tmpAdd += v
	}
	fmt.Printf("SumRange result %v\n", tmpAdd)
	return tmpAdd
}

// 不建议使用的函数写法,仅适用于短小函数，无论如何皆不建议编写此种指定返回值变量的函数
func div(a, b int) (rc, rd int) {
	rc, rd = rc*2, rd*3
	// 赋值给返回值变量,直接返回返回值
	return
}
