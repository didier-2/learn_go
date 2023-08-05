package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"go.learn/pkg/apis"
	"google.golang.org/protobuf/proto"

	"log"
	"os"
)

func NewRecord(filepath string) *record {
	return &record{
		filePath:         filepath,
		yamlFilePath:     filepath + ".yaml",
		protobufFilePath: filepath + ".proto.base64",
	}
}

type record struct {
	filePath         string
	yamlFilePath     string
	protobufFilePath string
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
	{
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
	}
	{
		data, err := proto.Marshal(pi)
		if err != nil {
			fmt.Println("marshal出错：", err)
			return err
		}
		{
			if err := r.writeFileWithAppendProtobuf(data); err != nil {
				log.Println("写入PROTOBUF时出错", err)
				return err
			}
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

func (r *record) writeFileWithAppendProtobuf(data []byte) error {
	file, err := os.OpenFile(r.protobufFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) //linux 权限
	if err != nil {
		fmt.Println("无法打开文件", r.filePath, "错误信息是", err)
		os.Exit(1)
	}
	defer file.Close()

	//newData := append([]byte("---\n"), data...)
	_, err = file.Write([]byte(base64.StdEncoding.EncodeToString(data)))
	if err != nil {
		return err
	}
	return nil
}