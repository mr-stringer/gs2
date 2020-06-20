package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jmcvetta/randutil"
)

//GetYears returns a slice of randutil.Choice which contains years and their weighting
func GetYears(StartYear, EndYear int) ([]randutil.Choice, error) {
	if StartYear > EndYear {
		return nil, errors.New("getYears: Start year is greater than end year")
	}

	/*Assign weighting to years*/
	years := make([]randutil.Choice, (EndYear-StartYear)+1)
	for i := 0; i < len(years); i++ {
		years[i].Item = StartYear
		years[i].Weight = RandInt(10, 20)
		StartYear++
	}
	return years, nil
}

//GetMonths returns a slice of randutil.Choice which contains months and their weighting
func GetMonths() []randutil.Choice {
	months := []randutil.Choice{
		{Weight: 10, Item: "01"},
		{Weight: 5, Item: "02"},
		{Weight: 3, Item: "03"},
		{Weight: 2, Item: "04"},
		{Weight: 3, Item: "05"},
		{Weight: 3, Item: "06"},
		{Weight: 3, Item: "07"},
		{Weight: 4, Item: "08"},
		{Weight: 5, Item: "09"},
		{Weight: 10, Item: "10"},
		{Weight: 12, Item: "11"},
		{Weight: 15, Item: "12"},
	}
	return months
}

//chanReturn is a simple struct that is popped down the channel when a goroutine quits.
type chanReturn struct {
	ok      bool
	message string
}

//RandomDate produces weighted random dates and puts passes it into a channel
func RandomDate(c configuration, date chan<- time.Time, ret chan<- chanReturn) {
	const shortForm = "2006-01-02"

	years, err := GetYears(c.StartYear, c.EndYear)
	if err != nil {
		log.Printf("RandomDate: Failed to get years\n")
		ret <- chanReturn{ok: false, message: err.Error()}
	}

	months := GetMonths()

	/*loop until told otherwise*/
	for {
		if len(date) < cap(date) {
			yearChoice, err := randutil.WeightedChoice(years)
			if err != nil {
				log.Printf("RandomDate: Failed drawing WeightedChoice for year\n")
				log.Print(err)
				ret <- chanReturn{ok: false, message: err.Error()}
			}
			year, ok := yearChoice.Item.(int)
			if !ok {
				log.Printf("RandomDate: Failed converting year to integer\n")
				ret <- chanReturn{ok: false, message: err.Error()}
			}

			monthChoice, err := randutil.WeightedChoice(months)
			if err != nil {
				log.Printf("RandomDate: Failed drawing WeightedChoice for month\n")
				log.Print(err)
				ret <- chanReturn{ok: false, message: err.Error()}
			}

			month, ok := monthChoice.Item.(string)
			if !ok {
				log.Printf("RandomDate: Failed converting month to integer\n")
				ret <- chanReturn{ok: false, message: err.Error()}
			}

			t, err := time.Parse(shortForm, fmt.Sprintf("%d-%s-01", year, month))
			if err != nil {
				log.Printf("RandomDate: failed to parse date\n")
				ret <- chanReturn{ok: false, message: err.Error()}
			}
			date <- t

		}
	}
}
