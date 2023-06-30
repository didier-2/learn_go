package main

import "testing"

func TestFrServiceSuggestion(t *testing.T) {
	realOutput := &fakeOutput{}
	frSvc := &fatRateService{
		s:      GetFaRateSuggestion(),
		input:  &FakeInput{},
		output: realOutput,
	}
	p := frSvc.input.GetInput()
	expOutput := &fakeOutput{p: p,
		s: "偏重",
	}
	//expectedSugg := "偏重"
	frSvc.GiveSuggestionToperson(&p)
	//if expectedSugg != sugg {
	//	t.Fatalf("预期:%s,实际的输出：%s", expectedSugg, sugg)
	//}

	if realOutput.s != expOutput.s {
		t.Fatalf("预期:%s,实际的输出：%s", expOutput.s, realOutput.s)
	}
}
