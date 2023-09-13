package apis

//type PersonalInformation struct {
//	Name   string  `json:"name"` //注意私有成员变量在序列化和反序列化时会被忽略
//	Sex    string  `json:"sex"`
//	Tall   float64 `json:"tall"`
//	Weight float64 `json:"weight"`
//	Age    int     `json:"age"`
//}

//type PersonalInformationFatRate struct {
//	Name    string
//	FatRate float64
//}

//type PersonalRank struct {
//	Name       string
//	Sex        string
//	RankNumber int
//	FatRate    float64
//}

func (*PersonalInformation) TableName() string {
	return "personal_information"
}
