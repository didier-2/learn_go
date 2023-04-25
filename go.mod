module go.learn

go 1.20

//replace go.learn.tools => ../go.learn.tools

require (
	github.com/armstrongli/go-bmi v0.0.1
	github.com/spf13/cobra v1.7.0
//github.com/armstrongli/go-bmi v0.0.1
//go.learn.tools v0.0.0-00010101000000-000000000000
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace github.com/armstrongli/go-bmi v0.0.1 => .\staging\src\github.com\armstrongli\go-bmi

replace github.com/spf13/cobra v1.7.0 => github.com/spf13/cobra v1.2.0
