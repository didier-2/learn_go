package main

import (
	"encoding/json"
	"fmt"
	"go.learn/pkg/apis"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func NewRecord(filepath string) *record {
	return &record{
		filePath:     filepath,
		yamlFilePath: filepath + ".yaml",
	}
}

type record struct {
	filePath     string
	yamlFilePath string
}

func (r *record) savePersonalInformation(pi *apis.PersonalInformation) error {
	{
		data, err := json.Marshal(pi)
		if err != nil {
			fmt.Println("marshal出错：", err)
			return err
		}
		if err := r.writeFileWithAppendJson(data); err != nil {
			log.Println("写入json时出错", err)
			return err
		}
	}
	data, err := yaml.Marshal(pi)
	if err != nil {
		fmt.Println("marshal出错：", err)
		return err
	}
	{
		if err := r.writeFileWithAppendYaml(data); err != nil {
			log.Println("写入yaml时出错", err)
			return err
		}
	}
	return nil
}

func (r *record) writeFileWithAppendJson(data []byte) error {
	file, err := os.OpenFile(r.filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) //linux 权限
	if err != nil {
		fmt.Println("无法打开文件", r.filePath, "错误信息是", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(append(data, '\n'))
	if err != nil {
		return err
	}
	return nil
}
func (r *record) writeFileWithAppendYaml(data []byte) error {
	file, err := os.OpenFile(r.yamlFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) //linux 权限
	if err != nil {
		fmt.Println("无法打开文件", r.filePath, "错误信息是", err)
		os.Exit(1)
	}
	defer file.Close()

	newData := append([]byte("---\n"), data...)
	_, err = file.Write(newData)
	if err != nil {
		return err
	}
	return nil
}
