package main

import "testing"

func Test_fatRateSuggestion_GetSuggestion(t *testing.T) {
	sugg := GetFaRateSuggestion()
	tests := []Person{
		{
			sex:     "男",
			age:     35,
			fatRate: 0.24,
		},
	}
	if got := sugg.GetSuggestion(&tests[0]); got != "肥胖" {
		t.Fail()
	}

}

func Test_fatRateSuggestion_GetSuggestion1(t *testing.T) {
	sugg := GetFaRateSuggestion()
	type args struct {
		person *Person
	}
	tests := []struct {
		name string //测试的名字
		args args   //字段
		want string //期望
	}{
		{
			name: "0.24-35",
			args: args{
				person: &Person{
					sex:     "男",
					age:     35,
					fatRate: 0.24,
				},
			},
			want: "肥胖",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sugg
			if got := s.GetSuggestion(tt.args.person); got != tt.want {
				t.Errorf("GetSuggestion() = %v, want %v", got, tt.want)
			}
		})
	}
}
