package main

import "fmt"

type InputFromStd struct {
}

func (InputFromStd) GetInput() Person {
	//录入各项
	var name string
	fmt.Print("姓名:")
	fmt.Scanln(&name)

	var weight float64
	fmt.Print("体重:")
	fmt.Scanln(&weight)

	var tall float64
	fmt.Print("身高:")
	fmt.Scanln(&tall)

	var age int
	fmt.Print("年龄:")
	fmt.Scanln(&age)

	sex := "男"
	fmt.Print("性别:")
	fmt.Scanln(&sex)

	return Person{
		name:   name,
		sex:    sex,
		tall:   tall,
		weight: weight,
		age:    age,
	}
}
