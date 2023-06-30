package main

import "log"

type fatRateService struct {
	input  InoutService
	s      *fatRateSuggestion
	output OutputService
}

func (frsvc *fatRateService) GiveSuggestionToperson(person *Person) {
	if err := person.calcBmi(); err != nil {
		log.Printf("无法给%s计算BMI:%v", person.name, err)
		return
	}
	person.calcFatRate()
	frsvc.output.Output(*person, frsvc.s.GetSuggestion(person))
}

//func (frsvc *fatRateService) GiveSuggestionTopersons(persons ...*Person) map[*Person]string {
//	out := map[*Person]string{}
//	for _, item := range persons {
//		out[item] = frsvc.GiveSuggestionToperson(item)
//	}
//	return out
//
//}
