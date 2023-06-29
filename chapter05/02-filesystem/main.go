package main

import "fmt"

func main() {
	{
		var equipment IOInterface = &Soft{}
		var data string
		//if equipment == "纸带" {
		//	data = readFromPaper()
		//} else if equipment == "磁带" {
		//	data = readFromMag()
		//} else if equipment == "1.4软盘" {
		data = equipment.Read()
		//}
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Mag{}
		var data string
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Paper{}
		var data string
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Sata{}
		var data string
		data = equipment.Read()
		fmt.Println(data)
	}
}

type IOInterface interface {
	Read() (data string)
}
type Soft struct {
}

func (Soft) Read() string {
	return "啦啦啦啦软盘"
}

type Sata struct {
}

func (Sata) Read() string {
	return "安安静静的Sata"
}

//	func readFrom14Sort() string {
//		return "啦啦啦啦软盘"
//	}
type Mag struct {
}

func (Mag) Read() string {
	return "滋滋滋滋滋滋磁带"
}

type Paper struct {
}

func (Paper) Read() string {
	return "从纸带读***********00000000000"
}
