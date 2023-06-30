package main

func main() {
	//person := getFakePersonInfo() //获取个人信息
	//c := Calc{}
	//c.BMI(person) //计算
	//c.FatRate(person)
	//fmt.Println(person)
	//sugg := fatRateSuggestion{}
	//suggestion := sugg.GetSuggestion(person) //拿建议
	//fmt.Println(suggestion)
	frSvc := &fatRateService{
		s:      GetFaRateSuggestion(),
		input:  &FakeInput{},
		output: &StdOut{},
	}
	//fakePerson := getFakePersonInfo()
	//fmt.Println(frSvc.GiveSuggestionToperson(fakePerson))
	for {
		p := frSvc.input.GetInput()
		frSvc.GiveSuggestionToperson(&p)
	}
}
