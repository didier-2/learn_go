package main

import (
	"fmt"
)

func main() {
	//person := getFakePersonInfo() //获取个人信息
	//c := Calc{}
	//c.BMI(person) //计算
	//c.FatRate(person)
	//fmt.Println(person)
	//sugg := fatRateSuggestion{}
	//suggestion := sugg.GetSuggestion(person) //拿建议
	//fmt.Println(suggestion)
	frSvc := &fatRateService{s: GetFaRateSuggestion()}
	fakePerson := getFakePersonInfo()
	fmt.Println(frSvc.GiveSuggestionToperson(fakePerson))
	for {
		p := getPersonInfoFromInput()
		fmt.Println(frSvc.GiveSuggestionToperson(p))
	}
}

func getPersonInfoFromInput() *Person {
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

	return &Person{
		name:   name,
		sex:    sex,
		tall:   tall,
		weight: weight,
		age:    age,
	}

}

func getFakePersonInfo() *Person {
	return &Person{
		name:   "xd",
		sex:    "男",
		tall:   176,
		weight: 76,
		age:    23,
	}
}
