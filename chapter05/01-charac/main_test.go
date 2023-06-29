package main

import (
	"fmt"
	"testing"
)

func TestAssertion(t *testing.T) {
	r := TestBox{}
	//var b Box = r
	var c Close = r

	switch cDetail := c.(type) {
	case Refrigerator:
		fmt.Println("是预期的对象")
		fmt.Println(cDetail.Size)
	case TestBox:
		fmt.Println("这是一个Box不能当冰箱用")

	}
	{
		refrigerator, ok := checkIsRefrigerator(c)
		if ok {
			fmt.Println("是个冰箱，开门装大象", refrigerator)
		} else {
			fmt.Println("这不是1个冰箱")
		}
	}
	//r2 := c.(Refrigerator)
	//fmt.Println(r2.Size)
}

func checkIsRefrigerator(c Close) (Refrigerator, bool) {
	r, ok := c.(Refrigerator)
	return r, ok
}

type TestBox struct {
}

func (tb TestBox) Close() error {
	return nil
}
