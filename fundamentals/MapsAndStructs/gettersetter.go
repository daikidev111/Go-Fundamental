package stuff

import (
	"errors"
	"fmt"
	"strconv"
	"time"
) 

var pl = fmt.Println

var Name string = "David"

func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, i := range intArr {
		strArrr = append(strArr, strconv.Itoa(i))
	}
	return strArr
}


// struct
type Date struct {
	day int
	month int
	year int
}


// setter
func (d *Date) setDay(day int) error {
	if (day < 1 || day > 31) {
		return errors.New("incorrect day value¥")
	}
	d.day = day
	return nil
}

func (d *Date) setMonth(month int) error {
	if (month < 1 || month > 12) {
		return errors.New("incorrect month value")
	}
	d.month = month
	return nil
}

func (d *Date) setYear(year int) error {
	if (year < 1900 || year > time.Now().Year()) {
		return errors.New("incorrect year value")
	}
	d.year = year
	return nil
}

// getter
func (d *Date) Day() int {return d.day}
func (d *Date) Month() int {return d.month}
func (d *Date) Year() int {return d.year}