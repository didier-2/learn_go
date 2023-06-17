package main

var defaultCalaculator = Calculator{}

type Calculator struct {
	left, right int
	op          string
	result      int
}
