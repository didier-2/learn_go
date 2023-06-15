package main

import (
	"fmt"
	"os"
	"time"
)

func deferGuess() {
	startTime := time.Now()
	defer fmt.Println("duration:", time.Now().Sub(startTime))
	defer func() {
		fmt.Println("duration upgraded:", time.Now().Sub(startTime))
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("start time: ", startTime)
}

func openFile() {
	fileName := "/test.txt"
	fileObj, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	defer fileObj.Close()

}
