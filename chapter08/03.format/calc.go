package main

import (
	gobmi "github.com/armstrongli/go-bmi"
	"go.learn/pkg/apis"
	"log"
)

type Calc struct {
}

func (Calc) BMI(person *apis.PersonalInformation) (float64, error) {
	bmi, err := gobmi.BMI(person.Weight, person.Tall)
	if err != nil {
		log.Println("error when calculating bmi :", err)
		return -1, err
	}

	return bmi, nil
}
func (c *Calc) FatRate(person *apis.PersonalInformation) (float64, error) {
	bmi, err := c.BMI(person)
	if err != nil {
		return -1, err
	}
	return gobmi.CalcFatRate(bmi, person.Age, person.Sex), nil
}
