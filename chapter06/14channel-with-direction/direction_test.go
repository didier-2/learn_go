package main

import "testing"

func TestDirection(t *testing.T) {
	c := make(chan int, 100)
	inOnly(c)
	outOnly(c)
}

func inOnly(c chan<- int) {
	c <- 1
	//<-c//当c是单向（入）channel时不能再取数。编译错误
}
func outOnly(c <-chan int) {
	//c<-1//当c是单向（出）channel时不能再存入数。编译错误
	<-c
}
