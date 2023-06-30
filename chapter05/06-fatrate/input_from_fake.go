package main

type FakeInput struct {
}

func (*FakeInput) GetInput() Person {
	return Person{
		name:   "xd",
		sex:    "男",
		tall:   176,
		weight: 76,
		age:    23,
	}
}
