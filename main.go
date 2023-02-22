package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const format = "02.01.2006"

func convertDate(format, date string) (time.Time, bool) {
	t, err := time.Parse(format, date)
	if err != nil {
		return time.Time{}, false
	}
	return t, t.Format(format) == date
}

func calculate(date time.Time) int64 {
	now := time.Now()
	diff := now.Sub(date)
	return int64(diff.Hours() / 24)
}

func main() {
	var (
		dateStr   = flag.String("date", "", "date for calculation, format: dd.mm.yyyy")
		countDays int64
	)
	flag.Parse()

	if len(*dateStr) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	/* check the format of the input date and convert it to time.Time */
	dateTime, checkFormat := convertDate(format, *dateStr)
	if !checkFormat {
		fmt.Printf("Your date %q does not match the format 'dd.mm.yyyy'\n", *dateStr)
		os.Exit(1)
	}

	countDays = calculate(dateTime)

	if countDays < 0 {
		countDays *= -1
		fmt.Printf("%d day(s) left until %s\n", countDays, *dateStr)
	} else {
		fmt.Printf("%d day(s) have passed since %s\n", countDays, *dateStr)
	}
}
