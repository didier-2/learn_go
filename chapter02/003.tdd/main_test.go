package _03_tdd

import "testing"

func TestCase1Part1(t *testing.T) {
	inputRecord("王强", 0.38) //录入
	inputRecord("王强", 0.32)
	{
		randOfWQ, fatRandOfWQ := getRand("王强")
		if randOfWQ != 1 {
			t.Fatalf("预期王强第一，但是得到的是:%d", randOfWQ)
		}
		if fatRandOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂是0.32，但得到的是:%f", fatRandOfWQ)
		}
	}
}

func TestCase1(t *testing.T) {
	inputRecord("王强", 0.38) //录入
	inputRecord("王强", 0.32)
	{
		randOfWQ, fatRandOfWQ := getRand("王强")
		if randOfWQ != 1 {
			t.Fatalf("预期王强第一，但是得到的是:%d", randOfWQ)
		}
		if fatRandOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂是0.32，但得到的是:%f", fatRandOfWQ)
		}
	}
	inputRecord("李静", 0.28)
	{
		randOfWQ, fatRandOfWQ := getRand("王强")
		if randOfWQ != 2 {
			t.Fatalf("预期王强第一，但是得到的是:%d", randOfWQ)
		}
		if fatRandOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂是0.32，但得到的是:%f", fatRandOfWQ)
		}
	}
	{
		randOfLJ, fatRandOfLJ := getRand("李静")
		if randOfLJ != 1 {
			t.Fatalf("预期李静第一，但是得到的是:%d", randOfLJ)
		}
		if fatRandOfLJ != 0.28 {
			t.Fatalf("预期李静的体脂是0.28，但得到的是:%f", fatRandOfLJ)
		}
	}

}
