package main

import "fmt"

func main() {

	test_base()
	test_use()
	test_null()
}

func test_base() {
	var a int = 10

	fmt.Printf("变量的地址: %x\n", &a)
}
func test_use() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}

func test_null() {
	var ptr *int

	fmt.Printf("ptr 的值为 : %x\n", ptr)
	/* ptr 不是空指针 */
	if (ptr != nil) {
		fmt.Println("ptr 不是空指针 ")
	}
	/* ptr 是空指针 */
	if (ptr == nil) {
		fmt.Println("ptr 是空指针 ")
	}
}
