package main

import "fmt"

func main() {
	a, b := 1, 2
	add(&a, &b)
	fmt.Println(a)
	c := &a //c的类型是*int,c指向a的盒子，*c就可以拿到a里面的东西，3
	d := &c //d的类型是**int，d指向c的盒子，d本身是指针，它存的东西，也是指针
	fmt.Println("d=", d, "*d=", *d, "**d=", **d)

	m := map[string]string{}
	mp1 := &m //mp1的类型是*map[string]string
	fmt.Println(mp1)
	put(m)
	fmt.Println("*mp1=", *mp1)

	f1 := add //f1=func（int,int）
	f1(&a, &b)
	fmt.Println("f1,add=", a)
	f1p1 := &f1 //f1p1=*func(int,int)
	(*f1p1)(&a, &b)
	fmt.Println("f1p1,add=", a)

	{
		var nothing *int
		//*nothing=3//注意这里是没有指向任何东西的int指针
		fmt.Println("nothing=", nothing)
	}
	{
		var nothingMap map[string]string
		//nothingMap["a"] = "BBB"
		fmt.Println("nothingMap=", nothingMap)
	}
	{
		var nothingSlice []int
		nothingSlice = append(nothingSlice, 100)
		fmt.Println("nothingSlice=", nothingSlice)
	}
}
func put(m map[string]string) {
	m["a"] = "AAA"
}

func add(a, b *int) {
	*a = *a + *b
}
