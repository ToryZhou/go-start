package main

var isActive bool                   // 全局变量声明
var enabled, disabled = true, false // 忽略类型的声明
func test2_1() {
	var available bool // 一般声明
	valid := false     // 简短声明
	available = true   // 赋值操作
	println(isActive, enabled, disabled, available, valid)
}

func main()  {
	test2_1()
}