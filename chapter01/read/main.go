package main

import (
	"fmt"
	"github.com/spf13/cobra"
	calc2 "go.learn/chapter01/faterate/calc"
)

func main() {
	//录入
	var (
		name   string
		sex    string
		tall   float64
		weight float64
		age    int
	)

	cmd := &cobra.Command{
		Use:   "healthcheck",
		Short: "体脂计算器，根据身高、体重、性别、年龄、计算体脂比，并给出健康建议",
		Long:  "该体脂计算器是基于BMI得体脂计算器",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("name  ", name)
			fmt.Println("sex   ", sex)
			fmt.Println("tall  ", tall)
			fmt.Println("weight", weight)
			fmt.Println("age   ", age)
			//计算
			bmi := calc2.CalcBMI(tall, weight)
			fatrate := calc2.CalcFatRate(bmi, age, sex)
			fmt.Println("fatRate", fatrate)

			//评估结果

		},
	}
	cmd.Flags().StringVar(&name, "name", "", "姓名")
	cmd.Flags().StringVar(&sex, "sex", "", "性别")
	cmd.Flags().Float64Var(&tall, "tall", 0, "身高")
	cmd.Flags().Float64Var(&weight, "weight", 0, "体重")
	cmd.Flags().IntVar(&age, "age", 0, "年龄")
	cmd.Execute()

}
