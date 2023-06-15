package calc

import "github.com/armstrongli/go-bmi"

func FatRate(bmi float64, age int, sex string) (fatRate float64) {
	gobmi.CalcFatRate(bmi, age, sex)
	return
}
