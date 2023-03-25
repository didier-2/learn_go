package main

import (
	"fmt"
	"time"
)

var tall float64
var weight float64

func main() {

	panicAndRecover()
	fmt.Println("done panicAndRecover,sleep somewhile")
	time.Sleep(10 * time.Second)

	closureMain()
	fmt.Println("done closureMain,sleep somewhile")
	time.Sleep(10 * time.Second)

	//close call
	fmt.Println(calcSum(21, 12, 12, 123, 23, 123, 123))
	fmt.Println(calcSum(34, 3, 4, 343, 434, 353, 53, 53))
	fmt.Println(calcSum(34, 3, 4, 343, 434, 353, 53, 53))
	fmt.Println(calcSum(34, 3, 4, 343, 43, 353, 53, 53))
	fmt.Println(calcSum(34, 3, 4, 343, 434, 353, 53, 53))
	fmt.Println(calcSum(34, 3, 4, 343, 434, 353, 53, 53))
	showUsedTime()
	fmt.Println(fib(1))
	fmt.Println("done calc,sleep somewhile")
	time.Sleep(10 * time.Second)

	//sampleSubdomain()
	fmt.Println("全局变量赋值前")
	caleBMI()
	tall, weight = 1.70, 70.0
	fmt.Println("全局变量赋值后")
	caleBMI()

	//重新定义重名的局部变量
	tall, weight := 100.70, 700.0
	caleBMI()
	tall, weight = 100.70, 700.0
	caleBMI()

	calculatorAdd := func(a, b float64) float64 {
		return a + b
	}
	result := calculatorAdd(1, 2)
	fmt.Println(result)
	{
		//fmt.Scanln....
		personTall := 1.11
		personWeight := 23.0
		calculatorAdd(personTall, personWeight)
	}
	{
		//fmt.Scanln....
		personTall := 1.11
		personWeight := 23.0
		calculatorAdd(personTall, personWeight)
	}
	fmt.Println(tall, weight)
}
func caleBMI() float64 {
	//tall,weight:="1.70",23
	fmt.Println(tall, weight)
	return 0
}
func sampleSubdomain() {
	name := "小丁"
	fmt.Println("名字", name)
	{
		fmt.Println("名字", name)
		name = "k"
		fmt.Println("名字", name)
	}
	fmt.Println("名字", name)

	if name == "小丁" {
		a := 3
		fmt.Println(a)
	} else {
		a := 4
		fmt.Println(a)
	}
}

func sampleSubdomainIf() {
	if v := caleBMI(); v == 0 {
		fmt.Println(v)
	} else {
		fmt.Println(v)
	}
	//fmt.Println(v) //无效，v的有效范围if block
}
func sampleSubdomainFor() {
	for i := 0; i < 10; i++ {
		fmt.Println("hallo golang,", i)
	}
	//fmt.Println(i) //i的有效范围for block
}
func fib(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
