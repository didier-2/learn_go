package main

import "fmt"

func (c *Calculator) Add() int { //注意要用指针

	tempResult := c.left + c.right
	c.result = tempResult
	fmt.Println("调用add函数，c的result=", c.result)
	return tempResult
}
func (c Calculator) Sub() int {
	return c.left - c.right
}
func (c Calculator) Multiple() int {
	return c.left * c.right
}
func (c Calculator) Divide() int {
	return c.left / c.right
}
func (c Calculator) Reminder() int {
	return c.left % c.right
}
