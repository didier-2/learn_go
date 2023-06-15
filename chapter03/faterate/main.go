package main

import (
	"fmt"
	"go.learn/chapter03/faterate/calc"
)

func main() {
	for {
		mainFatRateBody()
		if cont := whetherContinue(); !cont {
			break
		}
	}
}

//func recoverMainBody() {
//	if re := recover(); re != nil {
//		fmt.Printf("warning:catch critical error: %v\n", re)
//		debug.PrintStack()
//	}
//}

func mainFatRateBody() {
	//defer recoverMainBody()
	weight, tall, age, _, sex := getMaterialsFromInput()

	//fatRate:=calc.FatRate()
	fatRate, err := calcFatRate(weight, tall, age, sex)
	if err != nil {
		fmt.Println("warning:计算体脂率出错，不能再继续")
		return
	}

	//getHealthinessSuggestions(sex, age, fatRate)
	getHealthinessSuggestions(age, fatRate, getHealthinessSuggestionFromMale)
	getHealthinessSuggestions(age, fatRate, getHealthinessSuggestionFromFemale)

}

func getHealthinessSuggestions(age int, fatRate float64, getSuggestion func(age int, fatRate float64)) {
	//if sex == "男" {
	//	getSuggestion()
	//	getHealthinessSuggestionFromMale(age, fatRate)
	//} else {
	//	getHealthinessSuggestionFromFemale(age, fatRate)
	//}

	getSuggestion(age, fatRate)
}

func getHealthinessSuggestionFromFemale(age int, fatRate float64) {
	// 女性体脂表
	if age >= 18 && age <= 39 {
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
	} else if age >= 40 && age <= 59 {
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
	} else if age >= 60 {
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

func getHealthinessSuggestionFromMale(age int, fatRate float64) {
	// 男性体脂表
	if age >= 18 && age <= 39 {
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
	} else if age >= 40 && age <= 59 {
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
	} else if age >= 60 {
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
}

func calcFatRate(weight float64, tall float64, age int, sex string) (fatRate float64, err error) {
	//计算体脂
	bmi, err := calc.Bmi(tall, weight)
	if err != nil {
		return 0, err
	}
	//bmi := weight / (tall * tall)
	fatRate = calc.FatRate(bmi, age, sex) //(1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	fmt.Println("体脂率是：", fatRate)
	return
}

func getMaterialsFromInput() (float64, float64, int, int, string) {
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

	var sexWeight int
	sex := "男"
	fmt.Print("性别:")
	fmt.Scanln(&sex)

	if sex == "男" {
		sexWeight = 1
	} else if sex == "女" {
		sexWeight = 0
	}
	return weight, tall, age, sexWeight, sex
}

func whetherContinue() bool {
	var whetherContinue string
	fmt.Println("是否继续录入y/n")
	fmt.Scanln(&whetherContinue)
	if whetherContinue != "y" {
		return false
	}
	return true
}
