package variabledemo

import (
	"fmt"
	"strings"
)

func VarType() {
	/* 	 变量的命名规则
	   	1. var 变量名 类型名= 值
		2. 变量名 := 值

	*/
	/*
		golang的数据类型有
		bool
		number
		string
	*/
	flag := false
	i := 1
	f := 1.0
	b := 'a'
	bb := []byte{'a'}
	str := "str"
	arr := [1]int{1}
	slice := make([]int, 2)
	map1 := make(map[string]int)
	type Peron struct{}
	t := Peron{}
	t2 := &t
	ff := func() {}
	ch := make(chan int)

	fmt.Printf("%T\n", flag)
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", bb)
	fmt.Printf("%T\n", str)
	fmt.Printf("%T\n", arr)
	fmt.Printf("%T\n", slice)
	fmt.Printf("%T\n", map1)
	fmt.Printf("%T\n", t)
	fmt.Printf("%T\n", t2)
	fmt.Printf("%T\n", ff)
	fmt.Printf("%T\n", ch)
}

func Condition() {
	a, b := 1, 2
	if a > b {
		println("a larger than b")
	} else if a == b {
		println("a equals b ")
	} else {
		println("a smaller than b ")
	}
}

func TypeSwitchDemo() {
	var a interface{} = 1
	switch v := a.(type) {
	case string:
		println("string type:", v)
	case int:
		println("int type:", v)
	default:
		println("unknown type")
	}
}

func GotoDemo() {
	a := 1
	if a > 1 {
		goto A
	} else if a == 1 {
		goto B
	} else {
		println("smaller than 1 ")
	}
A:
	println("goto larger than 1")

B:
	println("goto equals 1")
}

func LoopDemo() {
	for i := 0; i < 10; i++ {
		println(i)
	}
	x := 1
	for {
		if x == 5 {
			break
		}
		x++
	}

	// for range
	str := "hello"
	for index, value := range str {
		println(index, value)
	}
}

func ArrDemo() {
	// 数组初始化时，必须是个常量表达式
	var arr [1]int = [1]int{1}
	arr2 := new([2]int)
	println(arr[0])
	println(len(arr2))
}

func SliceDemo() {
	slice := make([]int, 10)
	println(len(slice))
}

func StringDemo() {
	str := "hello"
	println(len(str))
	println(string(str[0:1]))

	// 字符串拼接
	builder := strings.Builder{}
	builder.WriteString("hello")
	builder.WriteString("world")
	println(builder.String())
}

func MapDemo() {
	mp := map[string]string{
		"k1": "v1",
		"k2": "v2",
	}
	println(len(mp))
	for k, v := range mp {
		println(k, v)
	}

	mp2 := make(map[string]int, 10)
	mp2["a1"] = 1
	mp2["b2"] = 2
	delete(mp2, "a1")

	// set
	set := make(map[string]struct{})
	set["s1"] = struct{}{}
	set["s2"] = struct{}{}
	println(len(set))

}
