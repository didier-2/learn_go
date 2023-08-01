package main

import (
	"encoding/json"
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
	"go.learn/pkg/apis"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	input := InputFromStd{}
	calc := &Calc{}
	rk := &FatRateRank{}
	record := NewRecord("D:/Desktop/record.txt")
	for {
		pi := input.GetInput()
		if err := record.savePersonalInformation(pi); err != nil {
			log.Fatal("保存失败：", err)
		}
		fr, err := calc.FatRate(pi)
		if err != nil {
			log.Fatal("计算体脂率失败", err)

		}
		rk.inputRecord(pi.Name, fr)
		rankResult, _ := rk.getRand(pi.Name)
		fmt.Println("排名结果：", rankResult)
		//caseStudy()
	}

}

func caseStudy() {
	filepath := "D:/Desktop/小丁2.json"
	personalInformation := apis.PersonalInformation{
		Name:   "小强",
		Sex:    "男",
		Tall:   1.70,
		Weight: 71,
		Age:    35,
	}

	fmt.Printf("%+v\n\n", personalInformation)
	data, err := json.Marshal(personalInformation)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Marshal的结果是(原生)：", data)
	fmt.Println("Marshal的结果是：", string(data))
	writeFile(filepath, data)
	readFile(filepath)
}

func readFile(filepath string) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	//infos := strings.Split(string(data), ",")
	fmt.Println("读取出来的内容是：", string(data))

	personalInformation := apis.PersonalInformation{}
	json.Unmarshal(data, &personalInformation) //todo handle error
	fmt.Println("开始计算体脂信息", personalInformation)
	bmi, err := gobmi.BMI(personalInformation.Weight, personalInformation.Tall) //todo handle error
	fmt.Printf("%s的bmi是：%v", personalInformation.Name, bmi)
	fatRate := gobmi.CalcFatRate(bmi, personalInformation.Age, personalInformation.Sex)
	fmt.Printf("%s的体脂率是：%v", personalInformation.Name, fatRate)

}

func writeFile(filepath string, data []byte) {
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("无法打开文件", filepath, "错误信息是", err)
		os.Exit(1)
	}
	defer file.Close()

	//b := make([]byte, 10)

	//var n int
	_, err = file.Write([]byte(data))
	fmt.Println(err)
	//_, err = file.Write([]byte("this is first line\n"))
	//fmt.Println(err)
	//_, err = file.WriteString("第二行内容\n")
	//fmt.Println(err)
	//_, err = file.WriteAt([]byte("CHANGED"), 0)
	//fmt.Println(err)
}
