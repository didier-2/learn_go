package calc

import "fmt"

func CalcBMI(tall float64, weight float64) (bmi float64, err error) {
	if tall <= 0 {
		return 0, fmt.Errorf("身高不能为0或者负数")
	}
	//todo
	return weight / (tall * tall), nil
}
