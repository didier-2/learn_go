package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.learn/pkg/apis"
	"log"
)

func main() {

	learnDB, err := getDBConnection()
	defer learnDB.Close()
	err = testDBConnection(err, learnDB)

	queryAllWithHack(err, learnDB)
	//queryAllData(err, learnDB)
	//insertData(learnDB) //todo
	//queryAllData(err, learnDB)
}

func insertData(learnDB *sql.DB) error {
	_, err := learnDB.Exec(fmt.Sprintf("insert into personal_information(name,sex,tall,weight,age) values ('%s','%s','%f','%f','%d')",
		"xiaob",
		"女",
		1.6,
		45.0,
		23,
	))
	if err != nil {
		fmt.Printf("新增数据失败:%v\n", err)
		return err
	}
	return nil
}
func queryAllWithHack(err error, learnDB *sql.DB) {
	_sql := fmt.Sprintf(`select  name,age from personal_information  where name ='%s'and sex='%s'`, "小强", "男")

	raws, err := learnDB.Query(_sql) //更换列的顺序，因为数据类型不匹配导致失败
	//警告：如果不巧，数据类型兼容，会引发更大的灾难
	if err != nil {
		log.Fatal(err)
	}

	list := &apis.PersonalInformationList{}
	for raws.Next() {
		var name string
		var age int
		if err := raws.Scan(&name, &age); err != nil {
			log.Printf("扫描数据失败，跳过该行")
		}
		fmt.Printf("name:%s\t,age:%d\n", name, age)
		list.Items = append(list.Items, &apis.PersonalInformation{
			Name: name,
			Age:  int64(age),
		})
	}
	data, _ := json.Marshal(list)
	fmt.Println(string(data))
}
func queryAllData(err error, learnDB *sql.DB) {
	raws, err := learnDB.Query("select  name,age from personal_information ") //更换列的顺序，因为数据类型不匹配导致失败
	//警告：如果不巧，数据类型兼容，会引发更大的灾难
	if err != nil {
		log.Fatal(err)
	}

	list := &apis.PersonalInformationList{}
	for raws.Next() {
		var name string
		var age int
		if err := raws.Scan(&name, &age); err != nil {
			log.Printf("扫描数据失败，跳过该行")
		}
		fmt.Printf("name:%s\t,age:%d\n", name, age)
		list.Items = append(list.Items, &apis.PersonalInformation{
			Name: name,
			Age:  int64(age),
		})
	}
	data, _ := json.Marshal(list)
	fmt.Println(string(data))
}

func testDBConnection(err error, learnDB *sql.DB) error {
	if err = learnDB.Ping(); nil != err {
		log.Fatalf("DB测试失败:%v", err)
	}
	fmt.Println("数据库链接成功")
	return err
}

func getDBConnection() (*sql.DB, error) {
	learnDB, err := sql.Open("mysql", "root:123456@tcp(localhost)/testdb")
	if err != nil {
		log.Fatalf("链接数据库失败:%v", err)
	}
	return learnDB, err
}
