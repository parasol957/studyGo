package main

import "fmt"

func main() {
	var i int      //1指定后不赋值 就使用默认值
	i = 10         //赋值
	var num = 10   //2自行判断字符串类型
	text := "德玛西亚" //第三种类型 可以省略关键字 声明并赋值加:
	fmt.Println("i=", i)
	fmt.Println("num=", num)
	fmt.Println("text=", text)

}
