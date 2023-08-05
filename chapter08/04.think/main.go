package main

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"go.learn/pkg/apis"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	//5万条记录
	counter := 50000
	persons := make([]*apis.PersonalInformation, 0, counter)

	start := time.Now()
	data, _ := ioutil.ReadFile("D:\\Desktop\\data.json")
	json.Unmarshal(data, &persons)
	finish := time.Now()
	fmt.Println("json unmarshal", finish.Sub(start))
	//for i := 0; i < counter; i++ {
	//	persons = append(persons, &apis.PersonalInformation{
	//		Name:   "xiao",
	//		Sex:    "男",
	//		Tall:   1.70,
	//		Weight: 71,
	//		Age:    35,
	//	})
	//}

	//json yaml protobuf分别序列化，记录序列化耗时
	//保存文件，记录耗时
	{
		fmt.Println("序列化Json")
		starTime := time.Now()
		data, err := json.Marshal(persons)
		if err != nil {
			log.Fatal(err)
		}
		finishMarshalTime := time.Now()
		ioutil.WriteFile("D:\\Desktop\\data.json", data, 0777)
		finishUnmarshalTime := time.Now()
		fmt.Println("序列化耗时", finishMarshalTime.Sub(starTime))
		fmt.Println("写文件耗时", finishUnmarshalTime.Sub(finishMarshalTime))

	}
	{
		fmt.Println("序列化Yaml")
		starTime := time.Now()
		data, err := yaml.Marshal(persons)
		if err != nil {
			log.Fatal(err)
		}
		finishMarshalTime := time.Now()
		ioutil.WriteFile("D:\\Desktop\\data.yaml", data, 0777)
		finishUnmarshalTime := time.Now()
		fmt.Println("序列化耗时", finishMarshalTime.Sub(starTime))
		fmt.Println("写文件耗时", finishUnmarshalTime.Sub(finishMarshalTime))

	}
	{
		fmt.Println("序列化PROTOBUF")
		pLister := &apis.PersonalInformationList{Items: persons}
		starTime := time.Now()
		data, err := proto.Marshal(pLister)
		if err != nil {
			log.Fatal(err)
		}
		finishMarshalTime := time.Now()
		ioutil.WriteFile("D:\\Desktop\\data.PROTOBUF", data, 0777)
		finishUnmarshalTime := time.Now()
		fmt.Println("序列化耗时", finishMarshalTime.Sub(starTime))
		fmt.Println("写文件耗时", finishUnmarshalTime.Sub(finishMarshalTime))

	}

}
