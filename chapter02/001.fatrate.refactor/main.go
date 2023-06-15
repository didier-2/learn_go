package _01_fatrate_refactor

import "fmt"

func CalcBMI(tall float64, weight float64) (bmi float64) {
	if tall <= 0 {
		fmt.Println("身高不能为零或者负数")
	}
	return weight / (tall * tall)
}
