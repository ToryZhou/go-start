package main

import "fmt"

/* 声明全局变量 */
var a_scope int = 20;

func main() {
	/* main 函数中声明局部变量 */
	var a_scope int = 10
	var b int = 20
	var c int = 0

	fmt.Printf("main()函数中 a = %d\n",  a_scope);
	c = sum( a_scope, b);
	fmt.Printf("main()函数中 c = %d\n",  c);
}

/* 函数定义-两数相加 */
func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n",  a);
	fmt.Printf("sum() 函数中 b = %d\n",  b);

	return a + b;
}