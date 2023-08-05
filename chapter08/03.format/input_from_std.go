package main

import (
	"fmt"
	"go.learn/pkg/apis"
)

type InputFromStd struct {
}

func (InputFromStd) GetInput() *apis.PersonalInformation {
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

	return &apis.PersonalInformation{
		Name:   name,
		Sex:    sex,
		Tall:   float32(tall),
		Weight: float32(weight),
		Age:    float32(age),
	}
}
