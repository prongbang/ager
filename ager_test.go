package ager_test

import (
	"github.com/prongbang/ager"
	"testing"
	"time"
)

func TestAge(t *testing.T) {
	// Given
	format := "2/1/2006"
	todayDate, _ := time.Parse(format, "9/2/2022")
	birthDate, _ := time.Parse(format, "8/1/2022")

	// When
	age := ager.Age(birthDate, todayDate)
	aging := ager.Format(age.Year, age.Month, age.Day)

	// Then
	if aging != "1D-1M" {
		t.Error("Error", aging)
	}
}

func TestFormat(t *testing.T) {
	age := ager.Duration{
		Day:   1,
		Month: 2,
		Year:  2023,
	}

	// When
	aging := ager.Format(age.Year, age.Month, age.Day)

	// Then
	if aging != "1D-2M-2023Y" {
		t.Error("Error", aging)
	}
}

func TestAdd(t *testing.T) {
	// Given
	format := "2/1/2006"
	todayDate, _ := time.Parse(format, "9/2/2022")
	duration := ager.Duration{
		Day:   1,
		Month: 1,
		Year:  2,
	}

	// When
	actual := ager.Add(todayDate, duration)

	// Then
	if actual.Day() != 10 && actual.Month() != 3 && actual.Year() != 2024 {
		t.Error("Error", actual)
	}
}

func TestDifference(t *testing.T) {
	// Given
	format := "2/1/2006"
	fromDate, _ := time.Parse(format, "9/1/2022")
	toDate, _ := time.Parse(format, "8/3/2023")

	// When
	actual := ager.Difference(fromDate, toDate)

	// Then
	if actual.Day != 27 && actual.Month != 1 && actual.Year != 1 {
		t.Error("Error", actual)
	}
}
