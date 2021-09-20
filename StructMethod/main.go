package main

import "fmt"

type StringExs string
type Point struct {
	x float64
	y float64
}

func (v StringExs) Jo() string {
	return string(v + " Expand")
}

func (v Point) PointAdd() Point {
	return Point{x: v.x + 1, y: v.y + 1}
}
func (v *Point) PointPtrPower(iv *Point) *Point {
	*v = Point{x: v.x * iv.x, y: v.y * iv.y}
	return v
}

func main() {
	// StringExs是个字符串类型,定义一个变量e,但它其实更像是抽象的字符串值类型
	var e StringExs
	e = "xiao"
	// 调用string方法，相当于java的类方法，但作者觉得更像c#的扩展方法
	println(e.Jo())

	// Point是结构体类型，更像java面向对象的类型
	var pt = Point{x: 1, y: 2}
	// 调用结构体变量的方法,输出返回的值。但由于struct仍是值类型,方法返回之后传递的值是最初值,除非指定接收返回的值
	fmt.Printf("%v", pt.PointAdd())
	// 可指定接收返回的值:
	// pt = pt.PointAdd()
	// 否则输出的仍是最初值:
	fmt.Printf("%v", pt)

	// Point是结构体类型，但定义成指针类型，因此几乎等同与java的类型引用
	var p *Point
	// 取赋值的地址
	p = &Point{x: 2, y: 3}
	// 调用指针结构体p的方法相乘，因为是指针，所以不必指定接收返回的值，因为接收的是地址，会根本上改变传递的值
	fmt.Printf("%v", *(p.PointPtrPower(p)))
	// 输出指针p的值
	fmt.Printf("%v", *p)
	return
}
