package main

import (
	"fmt"
	"math/rand"
)

type Math = int
type English = int
type Chinese = int

func main() {
	var mathScore int = 100
	var mathScore2 Math = 100
	mathScore2 = mathScore
	fmt.Println(mathScore2)
	getScoresOfstudent("dd")

	vg := &voteGate{[]*student{
		&student{name: fmt.Sprintf("%d", 1)},
		&student{name: fmt.Sprintf("%d", 2)},
		&student{name: fmt.Sprintf("%d", 3)},
		&student{name: fmt.Sprintf("%d", 4)},
		&student{name: fmt.Sprintf("%d", 5)},
	}}
	leader := vg.goRun()
	fmt.Println(leader)

}
func getScoresOfstudent(name string) (Math, Chinese, English) {
	//TODO
	return 0, 0, 0
}

type voteGate struct {
	students []*student
}

func (g *voteGate) goRun() *Leader {
	//TODO
	for _, item := range g.students {
		randInt := rand.Int()
		if randInt%2 == 0 {
			item.voteA(g.students[randInt%len(g.students)])
		} else {
			item.voteB(g.students[randInt%len(g.students)])

		}
	}

	maxSorce := -1
	maxSorceIndex := -1
	for i, item := range g.students {
		if maxSorce < item.agree {
			maxSorce = item.agree
			maxSorceIndex = i
		}
	}
	if maxSorceIndex >= 0 { //判断是否>0,因为如果没有学生，那么index的默认值是-1
		return (*Leader)(g.students[maxSorceIndex])
	}
	return nil
}

type Leader student

type student struct {
	name     string
	agree    int
	disagree int
}

func (std *student) voteA(target *student) {
	target.agree++
}
func (std *student) voteB(target *student) {
	target.disagree++
}
