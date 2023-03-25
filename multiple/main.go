package main

import "fmt"

func main() {
	//bmis := []float64{1.1, 2.222, 3.2323123}
	avgBmi := calculateAvg(1, 12, 2323, 4, 2, 4, 42, 4)

	fmt.Println(avgBmi)
	avgBmi = calculateAvgSlice([]float64{})
	fmt.Println(avgBmi)
	getScoresOfStudent("Tom")

}
func calculateAvg(bmis ...float64) float64 {
	var total float64 = 0
	for _, bmi := range bmis {
		total = total + bmi
	}
	return total / float64(len(bmis))
}
func calculateAvgSlice(bmis []float64) float64 {
	var total float64 = 0
	for _, bmi := range bmis {
		total = total + bmi
	}
	return total / float64(len(bmis))
}

func getScoresOfStudent(stdName string) (chinese int, match int, english int, physics int, nature int) {
	return 0, 0, 0, 0, 0
}
