package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func connectDb() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:123456@tcp(localhost)/testdb"))
	if err != nil {
		log.Fatal("数据库链接失败", err)
	}
	fmt.Println("链接数据库成功")
	return conn
}

func creatNewPerson(conn *gorm.DB) error {
	resp := conn.Create(&PersonalInformation{
		Name:   "小傲傲",
		Sex:    "男",
		Tall:   1.70,
		Weight: 71,
		Age:    35,
	})

	if err := resp.Error; err != nil {
		fmt.Println("创建人某某时失败", err)
		return err
	}
	fmt.Println("创建人某某时成功")
	return nil
}

func ormSelect(conn *gorm.DB) {

	output := make([]*PersonalInformation, 0)
	resp := conn.Where(&PersonalInformation{Name: "小强"}).Find(&output)
	if resp.Error != nil {
		fmt.Println("查找失败", resp.Error)
		return
	}
	data, _ := json.Marshal(output)
	fmt.Printf("查询结果：%v\n", string(data))
}
func ormSelectWithInaccurateQuery(conn *gorm.DB) {
	output := make([]*PersonalInformation, 0)
	resp := conn.Where("tall>?", 1.7).Find(&output)
	if resp.Error != nil {
		fmt.Println("模糊查找失败", resp.Error)
		return
	}
	data, _ := json.Marshal(output)
	fmt.Printf("模糊查询结果：%v\n", string(data))
}

func ormSelectWithInaccurateQueryHack(conn *gorm.DB) {
	output := make([]*PersonalInformation, 0)
	resp := conn.Where("name=? and sex=?", "'小强' -- ", "男").Find(&output)
	if resp.Error != nil {
		fmt.Println("模糊查找失败", resp.Error)
		return
	}
	data, _ := json.Marshal(output)
	fmt.Printf("模糊查询结果：%v\n", string(data))
}

func updateExistingPerson(conn *gorm.DB) error {
	resp := conn.Updates(&PersonalInformation{
		Id:     5,
		Name:   "小傲傲",
		Sex:    "男",
		Tall:   1.70,
		Weight: 71,
		Age:    35,
	})

	if err := resp.Error; err != nil {
		fmt.Println("更新人某某时失败", err)
		return err
	}
	fmt.Println("更新人某某时成功")
	return nil

}
func updateExistingPersonSelectedFiles(conn *gorm.DB) error {
	p := &PersonalInformation{
		Id:     5,
		Name:   "小傲傲",
		Sex:    "男",
		Tall:   1.80,
		Weight: 91,
		Age:    55,
	}
	resp := conn.Model(p).Select("name", "tall").Updates(p)

	if err := resp.Error; err != nil {
		fmt.Println("更新人某某时失败", err)
		return err
	}
	fmt.Println("更新人某某时成功")
	return nil
}

func deletePersons(conn *gorm.DB) {
	p := &PersonalInformation{
		Id: 5,
	}
	resp := conn.Delete(p)

	if err := resp.Error; err != nil {
		fmt.Println("删除人某某时失败", err)
		return
	}
	fmt.Println("删除人某某时成功")
	return
}

func main() {
	conn := connectDb()
	fmt.Println(conn)
	//if err := creatNewPerson(conn); err != nil {
	//	//todo
	//}
	//ormSelect(conn)
	//ormSelectWithInaccurateQuery(conn)

	//ormSelectWithInaccurateQueryHack(conn)
	//updateExistingPerson(conn)
	//updateExistingPersonSelectedFiles(conn)
	deletePersons(conn)
}
