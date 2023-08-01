package apis

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"testing"
)

func TestMarshalJson(t *testing.T) {
	personalInformation := PersonalInformation{
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
}
func TestUnmarshalJson(t *testing.T) {
	data := `{"name":"小强","sex":"男","tall":1.7,"weight":71,"age":35}`
	personalInformation := PersonalInformation{}
	json.Unmarshal([]byte(data), &personalInformation)
	fmt.Println(personalInformation)
}

func TestMarshalYaml(t *testing.T) {
	personalInformation := PersonalInformation{
		Name:   "小强",
		Sex:    "男",
		Tall:   1.70,
		Weight: 71,
		Age:    35,
	}

	fmt.Printf("%+v\n\n", personalInformation)
	data, err := yaml.Marshal(personalInformation)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Marshal的结果是(原生)：", data)
	fmt.Println("Marshal的结果是(string)：", string(data))
}
func TestUnmarshalYaml(t *testing.T) {
	data := `
name: 小强
sex: 男
tall: 1.7
weight: 71
age: 35
`
	personalInformation := PersonalInformation{}
	yaml.Unmarshal([]byte(data), &personalInformation)
	fmt.Println(personalInformation)
}
