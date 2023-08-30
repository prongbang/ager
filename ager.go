package ager

import (
	"fmt"
	"time"
)

type Duration struct {
	Day   int
	Month int
	Year  int
}

func (dd Duration) String() string {
	return fmt.Sprintf("Year: %d, Month: %d, Day: %d", dd.Year, dd.Month, dd.Day)
}

var daysInMonth = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func DaysInMonth(year, month int) int {
	if month == int(time.February) && IsLeapYear(year) {
		return 29
	}
	return daysInMonth[month-1]
}

func IsLeapYear(year int) bool {
	return (year%4 == 0) && ((year%100 != 0) || (year%400 == 0))
}

func Difference(fromDate, toDate time.Time) Duration {
	endDate := toDate
	years := endDate.Year() - fromDate.Year()
	months := 0
	days := 0
	if fromDate.Month() > endDate.Month() {
		years--
		months = 12 + int(endDate.Month()) - int(fromDate.Month())
		if fromDate.Day() > endDate.Day() {
			months--
			days = DaysInMonth(fromDate.Year()+years, int((fromDate.Month()+time.Month(months-1))%12+1)) + endDate.Day() - fromDate.Day()
		} else {
			days = endDate.Day() - fromDate.Day()
		}
	} else if endDate.Month() == fromDate.Month() {
		if fromDate.Day() > endDate.Day() {
			years--
			months = 11
			days = DaysInMonth(fromDate.Year()+years, int((fromDate.Month()+time.Month(months-1))%12+1)) + endDate.Day() - fromDate.Day()
		} else {
			days = endDate.Day() - fromDate.Day()
		}
	} else {
		months = int(endDate.Month()) - int(fromDate.Month())
		if fromDate.Day() > endDate.Day() {
			months--
			days = DaysInMonth(fromDate.Year()+years, int(fromDate.Month()+time.Month(months))) + endDate.Day() - fromDate.Day()
		} else {
			days = endDate.Day() - fromDate.Day()
		}
	}
	return Duration{Day: days, Month: months, Year: years}
}

func Add(date time.Time, duration Duration) time.Time {
	years := date.Year() + duration.Year
	years += int((date.Month() + time.Month(duration.Month)) / 12)
	months := int((date.Month() + time.Month(duration.Month)) % 12)
	days := date.Day() + duration.Day - 1
	return time.Date(years, time.Month(months), 1, 0, 0, 0, 0, date.Location()).AddDate(0, 0, days)
}

func Age(birthDate time.Time, today ...time.Time) Duration {
	now := time.Now()
	if len(today) > 0 {
		now = today[0]
	}
	return Difference(birthDate, now)
}

func Format(years, months, days int) string {
	format := ""
	if days > 0 {
		format += fmt.Sprintf("%dD-", days)
	}
	if months > 0 {
		format += fmt.Sprintf("%dM-", months)
	}
	if years > 0 {
		format += fmt.Sprintf("%dY", years)
	}
	// Re-format
	size := len(format)
	if size > 0 {
		if format[size-1:] == "-" {
			format = format[:size-1]
		}
	} else {
		format = "NA"
	}
	return format
}
