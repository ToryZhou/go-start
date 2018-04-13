package main

import "fmt"

func main() {
	test7_1()
	test7_2()
}

func test7_1() {
	var C, c int //声明变量
	C = 1        /*这里不写入FOR循环是因为For语句执行之初会将C的值变为1，当我们goto A时for语句会重新执行（不是重新一轮循环）*/
A:
	for C < 100 {
		C++ //C=1不能写入for这里就不能写入
		for c = 2; c < C; c++ {
			if C%c == 0 {
				goto A //若发现因子则不是素数
			}
		}
		fmt.Println(C, "是素数")
	}
}

func test7_2() {
	for i := 1; i <= 9; i++ { // i 控制行，以及计算的最大值
		for j := 1; j <= i; j++ { // j 控制每行的计算个数
			fmt.Printf("%d*%d=%d ", j, i, j*i)
		}
		fmt.Println("")
	}
}
