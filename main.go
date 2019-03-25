package main

import (
	"fmt"
	"strconv"
	"strings"

	//"strings"
	"time"

	"github.com/fatih/color"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

func isLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}

	return false
}

func getDaysOfYear() int {
	days := 0

	year := time.Now().Year()

	for month := 1; month <= 12; month++ {
		if month != 2 {
			if month == 4 || month == 6 || month == 9 || month == 11 {
				days += 30

			} else {
				days += 31
			}
		} else {
			if isLeapYear(year) {
				days += 29
			} else {
				days += 28
			}
		}
	}
	return days
}

func main() {
	strUp := `
 __               __                   
|  \  _      _   |__)  _  _   _  .  _  
|__/ (_| \/ _)   | \  (- ||| (_| | | ) 
         /`
	now := time.Now()
	monthInt := int(now.Month())
	monthStr := strconv.Itoa(monthInt)
	if monthInt < 10 {
		monthStr = "0" + strconv.Itoa(monthInt)
	}
	today, err := time.Parse(TIME_LAYOUT, strconv.Itoa(now.Year())+
		"-"+monthStr+
		"-"+strconv.Itoa(now.Day())+
		" 00:00:00")
	if err != nil {
		panic(err)
	}
	startOfYear, err := time.Parse(TIME_LAYOUT, strconv.Itoa(now.Year())+"-01-01 00:00:00")
	if err != nil {
		panic(err)
	}
	passDays := today.Sub(startOfYear)
	passDaysInt := int(passDays.Hours() / 24)

	color.Set(color.FgGreen, color.Bold)
	defer color.Unset()
	fmt.Println(strUp)
	daysOfYear := getDaysOfYear()
	width := 40
	currentSaucerSize := int(float64(passDaysInt) / float64(daysOfYear) * float64(width))
	strDown := "[" +
		strings.Repeat("🔴️", currentSaucerSize) +
		strings.Repeat("🔵", width-currentSaucerSize) +
		"] " +
		strconv.Itoa(passDaysInt) + "/" + strconv.Itoa(daysOfYear)
	fmt.Println(strDown)
}
