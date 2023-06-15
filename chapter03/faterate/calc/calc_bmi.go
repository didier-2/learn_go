package calc

import (
	"github.com/armstrongli/go-bmi"
	//_ "go.learn/staging/src/github.com/armstrongli/go-bmi"
	//gobmi "go.learn/staging/src/github.com/armstrongli/go-bmi"
)

func Bmi(tall float64, weight float64) (bmi float64, err error) {
	bmi, _ = gobmi.BMI(weight, tall)
	return
	//if tall <= 0 {
	//	panic("身高不能为0或者负数")
	//}
	////todo
	//return weight / (tall * tall)

}
