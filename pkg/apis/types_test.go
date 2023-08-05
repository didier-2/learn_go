package apis

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml" //这是第三方开源的yaml工具，可以解决yaml格式化时引用protobuf生成的特殊字段
	"google.golang.org/protobuf/proto"
	originalyaml "gopkg.in/yaml.v3"
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
	data, err := originalyaml.Marshal(personalInformation)
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
	originalyaml.Unmarshal([]byte(data), &personalInformation)
	fmt.Println(personalInformation)
}
func TestMarshalYamlSpecial(t *testing.T) {
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

func TestMarshalProtobuf(t *testing.T) {
	personalInformation := &PersonalInformation{
		Name:   "小强",
		Sex:    "男",
		Tall:   1.70,
		Weight: 71,
		Age:    35,
	}
	data, err := proto.Marshal(personalInformation)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(data)
	fmt.Println(string(data))
	//通常在非程序交互过程中，要保留原生protobuf，可以直接写入文件。如果想要单行宝春，需要转码
	//选择的通用转码是：base64
	output64Data := base64.StdEncoding.EncodeToString(data)
	fmt.Println(output64Data)
}
