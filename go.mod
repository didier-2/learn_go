module go.learn

go 1.20

//replace go.learn.tools => ../go.learn.tools

require (
	github.com/armstrongli/go-bmi v0.0.1
	github.com/ghodss/yaml v1.0.0
	github.com/go-sql-driver/mysql v1.7.1
	github.com/spf13/cobra v1.7.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	gorm.io/driver/mysql v1.5.1
	gorm.io/gorm v1.25.1
//github.com/armstrongli/go-bmi v0.0.1
//go.learn.tools v0.0.0-00010101000000-000000000000
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/armstrongli/go-bmi v0.0.1 => .\staging\src\github.com\armstrongli\go-bmi

replace github.com/spf13/cobra v1.7.0 => github.com/spf13/cobra v1.2.0
