package main

import "fmt"

func main() {
	m1 := map[string]float64{"xz": 90, "xc": 85, "xv": 65, "xb": 49, "xn": 85, "xm": 65, "xa": 95, "xs": 95, "xd": 95, "xf": 95, "xg": 95, "xh": 95, "xj": 95, "xk": 95, "xl": 95, "xq": 95, "xw": 95, "xe": 95, "xr": 95, "xt": 95}
	var s float64
	var pingjun float64
	for _, scroe := range m1 {
		s = s + scroe
		fmt.Println(s)
	}
	pingjun = s / float64(len(m1))
	fmt.Println(pingjun)
}
