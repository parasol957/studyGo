package main

import "fmt" //fmt 提供格式化,输出,输入的函数

func main() {
	fmt.Println("你好\t世界") //一个制表位，实现对齐的功能
	fmt.Println("你好\n世界") //实现换行的功能
	fmt.Println("你好\\世界") //一个\
	fmt.Println("你好\"世界") //一个"
	fmt.Println("你好\r世界") //表示把'世界'顶到最前面去，覆盖掉以前的内容

	//姓名 张三
	//性别 男
	fmt.Println("姓名\t张三\n性别\t男\n年龄\t16")

}
