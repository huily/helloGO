//学习《Go编程时光》https://golang.iswbm.com/c01/c01_06.html
//字典
//字典由key和value组成，它们各自有各自的类型
// key 不能是切片，不能是字典，不能是函数。
//声明：map[KEY_TYPE]VALUE_TYPE

package lesson6

import "fmt"

//第一种方法初始化并直接赋值
var mydict map[string]int8 = map[string]int8{"a": 1, "b": 2}

func Lesson6() {
	fmt.Println(mydict)

	//第二种方法：  “ :=  ” 初始化并赋值
	mydict1 := map[int]string{1: "A", 2: "B"}
	fmt.Println(mydict1)

	//第三种方法：  “ make()  ” 初始化，然后下一步再赋值
	////  make 函数相当于对其初始化
	mydict2 := make(map[int]float32)
	mydict2[1] = 1.0
}
