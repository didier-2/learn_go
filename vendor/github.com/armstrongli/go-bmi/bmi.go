package gobmi

import "fmt"

func BMI(weightKG, heightM float64) (bmi float64, err error) {
	if weightKG < 0 {
		err = fmt.Errorf("weight cannoot be negative")
		return
	}
	if heightM < 0 {
		err = fmt.Errorf("heightM cannoot be negative")
		return
	}
	if heightM == 0 {
		err = fmt.Errorf("heightM cannoot be 0")
		return
	}
	bmi = weightKG / (heightM * heightM)
	return
}
