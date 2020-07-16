package main

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestGetYears(t *testing.T) {
	type args struct {
		StartYear int
		EndYear   int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Good 01", args{2010, 2020}, false},
		{"Good 02", args{1970, 1970}, false},
		{"Good 03", args{1992, 1999}, false},
		{"Good 04", args{2090, 2090}, false},
		{"Good 05", args{1066, 1800}, false},
		{"Bad 01", args{2020, 2010}, true},
		{"Bad 02", args{1970, 1969}, true},
		{"Bad 03", args{1992, 1900}, true},
		{"Bad 04", args{2190, 2090}, true},
		{"Bad 05", args{1066, 1000}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetYears(tt.args.StartYear, tt.args.EndYear)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetYears() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, choice := range got {
				if choice.Item.(int) < tt.args.StartYear || choice.Item.(int) > tt.args.EndYear {
					t.Errorf("GetYears() error year:%d is not in bounds", choice.Item.(int))
				}
			}

		})
	}
}

func TestGetMonths(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test 01"},
	}

	expected := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got := GetMonths()
			for _, el := range got {
				if !func(ex []string, g string) bool {
					for _, e := range expected {
						if e == g {
							return true
						}
					}
					return false
				}(expected, el.Item.(string)) {
					t.Errorf("GetMonth() returned unexpected element")
				}
			}

		})
	}
}

func TestRandomDate(t *testing.T) {
	type args struct {
		c    configuration
		date chan time.Time
	}
	tests := []struct {
		name string
		args args
	}{
		/*Only need to set the start and end years in the configuration*/
		/*Also, do not need to test badly set years as these are weeded out in validation of the config */
		{"Good 01", args{configuration{StartYear: 2010, EndYear: 2015, Orders: 100}, make(chan time.Time, 10)}},
		{"Good 02", args{configuration{StartYear: 2012, EndYear: 2012, Orders: 100}, make(chan time.Time, 10)}},
		{"Good 03", args{configuration{StartYear: 1990, EndYear: 1990, Orders: 100}, make(chan time.Time, 10)}},
		{"Good 04", args{configuration{StartYear: 1066, EndYear: 1100, Orders: 100}, make(chan time.Time, 10)}},
		{"Good 05", args{configuration{StartYear: 1001, EndYear: 2888, Orders: 100}, make(chan time.Time, 10)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			df := "2006-01-02"
			start, err := time.Parse(df, fmt.Sprintf("%d-01-01", tt.args.c.StartYear))
			if err != nil {
				t.Errorf("TestRandomDate() failed to parse time")
			}
			end, err := time.Parse(df, fmt.Sprintf("%d-12-31", tt.args.c.EndYear))

			/*Make adjustments*/
			start = start.Add(-time.Second)
			end = end.Add(time.Hour * 24)

			go RandomDate(tt.args.c, tt.args.date)
			for i := 0; i < tt.args.c.Orders; i++ {
				tm := <-tt.args.date
				if tm.Before(start) || tm.After(end) {
					t.Errorf("RandomDate() date given out-of-bounds.  %s should not be greater before %s or after %s", tm.String(), start.String(), end.String())
				} else {
					log.Printf("%s falls between %s & %s", tm.String(), start.String(), end.String())
				}

			}

		})
	}
}
