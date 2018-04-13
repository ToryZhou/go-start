package main

import "fmt"

import "unsafe"
const (
	a4 = "abc"
	b4 = len(a4)
	c4 = unsafe.Sizeof(a4)
)

func main() {
	test4_1()
	test4_2()
	test4_3()
}

func test4_1(){
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Println("面积为 : %d", area)
	println()
	println(a, b, c)
}

func test4_2()  {
	println(a4, b4, c4)
}

func test4_3()  {
	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)
}

