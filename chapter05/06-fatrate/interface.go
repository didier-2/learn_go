package main

type InoutService interface {
	GetInput() Person
}
type OutputService interface {
	Output(Person, string)
}
