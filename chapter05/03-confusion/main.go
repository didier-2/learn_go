package main

func main() {
	dd := downloadFromNetDisk{
		secret:   &mobileTokenDynamic{mobileNumber: "1232155465"},
		filepath: "接口编程.ppt",
	}
	dd.Download()
}

type DynamicSecret interface {
	GetSecret() string
}

func dynamicSecret() string {
	return ""
}

type mobileTokenDynamic struct {
	mobileNumber string
}

func (d *mobileTokenDynamic) GetSecret() string {
	return "something"
}

//通常开发时侯，第一个版本叫做：happy path
//剩下的是痛：他无法应对变更。简单的变更会带来更痛苦的维护
