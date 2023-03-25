package main

import (
	"fmt"
)

func main() {
	var totalFatRate float64
	names := [3]string{}
	weights := [3]float64{}
	talls := [3]float64{}
	ages := [3]int{}
	bmis := [3]float64{}
	fatRates := [3]float64{}
	var fatRate float64

	for i := 0; i < 3; i++ {
		name, weight, tall, age, sex := getInfoFromInput()

		bmi := caleBMI(tall, weight)
		bmis[i] = bmi

		if sex == "男" {
			sexWeight = 1
		} else if sex == "女" {
			sexWeight = 0
		}

		var femalefatRate float64 = (1.2*bmis[i] + 0.23*float64(ages[i]) - 5.4 - 10.8*float64(0)) / 100
		var malefatRate float64 = (1.2*bmis[i] + 0.23*float64(ages[i]) - 5.4 - 10.8*float64(1)) / 100

		fatRates[i] = fatRate
		fmt.Println("体脂率是：", fatRate)

		if sex == "男" {
			// 男性体脂表
			if ages[i] >= 18 && ages[i] <= 39 {
				if fatRate <= 0.1 {
					fmt.Println("偏瘦")
				} else if fatRate > 0.1 && fatRate <= 0.16 {
					fmt.Println("标准")
				} else if fatRate > 0.16 && fatRate <= 0.21 {
					fmt.Println("偏胖")
				} else if fatRate > 0.21 && fatRate <= 0.26 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("过胖")
				}
			} else if ages[i] >= 40 && ages[i] <= 59 {
				if fatRate <= 0.11 {
					fmt.Println("偏瘦")
				} else if fatRate > 0.11 && fatRate <= 0.17 {
					fmt.Println("标准")
				} else if fatRate > 0.17 && fatRate <= 0.22 {
					fmt.Println("偏胖")
				} else if fatRate > 0.22 && fatRate <= 0.27 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("过胖")
				}
			} else if ages[i] >= 60 {
				if fatRate <= 0.13 {
					fmt.Println("偏瘦")
				} else if fatRate > 0.13 && fatRate <= 0.19 {
					fmt.Println("标准")
				} else if fatRate > 0.19 && fatRate <= 0.24 {
					fmt.Println("偏胖")
				} else if fatRate > 0.24 && fatRate <= 0.29 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("过胖")
				}
			} else {
				fmt.Println("我们不看未成年人")
			}
		} else {
			// 女性体脂表
			if ages[i] >= 18 && ages[i] <= 39 {
				if fatRate <= 0.2 {
					fmt.Println("偏瘦")
				} else if fatRate > 0.2 && fatRate <= 0.27 {
					fmt.Println("标准")
				} else if fatRate > 0.27 && fatRate <= 0.34 {
					fmt.Println("偏胖")
				} else if fatRate > 0.34 && fatRate <= 0.39 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("过胖")
				}
			} else if ages[i] >= 40 && ages[i] <= 59 {
				if fatRate <= 0.21 {
					fmt.Println("偏瘦")
				} else if fatRate > 0.21 && fatRate <= 0.28 {
					fmt.Println("标准")
				} else if fatRate > 0.28 && fatRate <= 0.35 {
					fmt.Println("偏胖")
				} else if fatRate > 0.35 && fatRate <= 0.40 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("过胖")
				}
			} else if ages[i] >= 60 {
				if fatRate <= 0.22 {
					fmt.Println("偏瘦")
				} else if fatRate > 0.22 && fatRate <= 0.29 {
					fmt.Println("标准")
				} else if fatRate > 0.29 && fatRate <= 0.36 {
					fmt.Println("偏胖")
				} else if fatRate > 0.36 && fatRate <= 0.41 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("过胖")
				}
			} else {
				fmt.Println("我们不看未成年人")
			}
		}
		//var whetherContinue string
		//fmt.Println("是否继续录入y/n")
		//fmt.Scanln(&whetherContinue)
		//if whetherContinue != "y" {
		//	break
		//}
	}

	for i := 0; i < 3; i++ {
		fmt.Println(names[i], weights[i], talls[i], ages[i], bmis[i], fatRates[i])
	}
	fmt.Println(totalFatRate / 3)
}

func caleBMI(tall float64, weight float64) float64 {
	return tall / (weight * weight)
}

func getInfoFromInput() (name string, weight float64, tall float64, age int, sex string) {
	fmt.Print("name")
	fmt.Scanln(&name)

	//var weight float64
	fmt.Println("输入体重/kg")
	fmt.Scanln(&weight)

	//var tall float64
	fmt.Println("输入身高/m")
	fmt.Scanln(&tall)
	//bmis[i] = weights / (talls * talls)
	//var age int
	fmt.Println("输入年龄/岁")
	fmt.Scanln(&age)

	//var sexWeight int
	//var sex string
	fmt.Println("输入性别，男/女")
	fmt.Scanln(&sex)
	return
}
