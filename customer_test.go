package main

import (
	"log"
	"testing"
)

func TestRandomSal(t *testing.T) {
	type args struct {
		request chan bool
		sal     chan string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Good 01", args{make(chan bool), make(chan string)}},
	}

	/*Get the expected outcomes*/
	sals := GetWeightedSals()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go RandomSal(tt.args.request, tt.args.sal)
			for i := 0; i < 100; i++ {
				tt.args.request <- true
				res := <-tt.args.sal

				found := bool(false)

				for _, v := range sals {
					if v.Item == res {
						log.Printf("Match found - %s:%v", res, v.Item)
						found = true
						break
					}
				}
				if !found {
					t.Errorf("RandomSal() unexpected output")
				}
			}
		})
		close(tt.args.request)
		close(tt.args.sal)
	}
}

func TestRandomFirstIntial(t *testing.T) {
	type args struct {
		request chan bool
		fint    chan string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Good 01", args{make(chan bool), make(chan string)}},
	}

	/*Get the expected outcomes*/
	fints := GetWeightedFirstIntial()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go RandomFirstIntial(tt.args.request, tt.args.fint)
			for i := 0; i < 100; i++ {
				tt.args.request <- true
				res := <-tt.args.fint

				found := bool(false)

				for _, v := range fints {
					if v.Item == res {
						log.Printf("Match found - %s:%v", res, v.Item)
						found = true
						break
					}
				}
				if !found {
					t.Errorf("TestRandomFirstIntial() unexpected output")
				}
			}
		})
		close(tt.args.request)
		close(tt.args.fint)
	}
}

func TestRandomSurname(t *testing.T) {
	type args struct {
		request chan bool
		surname chan string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Good 01", args{make(chan bool), make(chan string)}},
	}

	/*Get the expected outcomes*/
	surname := GetWeightedSurnames()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go RandomSurname(tt.args.request, tt.args.surname)
			for i := 0; i < 100; i++ {
				tt.args.request <- true
				res := <-tt.args.surname

				found := bool(false)

				for _, v := range surname {
					if v.Item == res {
						log.Printf("Match found - %s:%v", res, v.Item)
						found = true
						break
					}
				}
				if !found {
					t.Errorf("RandomSurname() unexpected output")
				}
			}
		})
		close(tt.args.request)
		close(tt.args.surname)
	}
}

func TestRandomStreetName(t *testing.T) {
	type args struct {
		request    chan bool
		streetName chan string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Good 01", args{make(chan bool), make(chan string)}},
	}

	/*Get the expected outcomes*/
	streets := GetWeightedStreetNames()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go RandomStreetName(tt.args.request, tt.args.streetName)
			for i := 0; i < 100; i++ {
				tt.args.request <- true
				res := <-tt.args.streetName

				found := bool(false)

				for _, v := range streets {
					if v.Item == res {
						log.Printf("Match found - %s:%v", res, v.Item)
						found = true
						break
					}
				}
				if !found {
					t.Errorf("RandomStreetName() unexpected output")
				}
			}
		})
		close(tt.args.request)
		close(tt.args.streetName)
	}
}

func TestRandomTowntName(t *testing.T) {
	type args struct {
		request  chan bool
		townName chan string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Good 01", args{make(chan bool), make(chan string)}},
	}

	/*Get the expected outcomes*/
	townNames := GetWeightedTownNames()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go RandomTownName(tt.args.request, tt.args.townName)
			for i := 0; i < 100; i++ {
				tt.args.request <- true
				res := <-tt.args.townName

				found := bool(false)

				for _, v := range townNames {
					if v.Item == res {
						log.Printf("Match found - %s:%v", res, v.Item)
						found = true
						break
					}
				}
				if !found {
					t.Errorf("RandomTownName() unexpected output")
				}
			}
		})
		close(tt.args.request)
		close(tt.args.townName)
	}
}

func TestRandomDiscount(t *testing.T) {
	type args struct {
		request  chan bool
		discount chan int
	}
	tests := []struct {
		name string
		args args
	}{
		{"Good 01", args{make(chan bool), make(chan int)}},
	}

	/*Get the expected outcomes*/
	discounts := GetWeightedDiscount()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go RandomDiscount(tt.args.request, tt.args.discount)
			for i := 0; i < 100; i++ {
				tt.args.request <- true
				res := <-tt.args.discount

				found := bool(false)

				for _, v := range discounts {
					if v.Item == res {
						log.Printf("Match found - %d:%v", res, v.Item.(int))
						found = true
						break
					}
				}
				if !found {
					t.Errorf("RandomDiscount() unexpected output")
				}
			}
		})
		close(tt.args.request)
		close(tt.args.discount)
	}
}

func TestRandomStreetNumber(t *testing.T) {
	type args struct {
		request   chan bool
		streetNum chan int
	}
	tests := []struct {
		name string
		args args
	}{
		{"Good", args{make(chan bool), make(chan int)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go RandomStreetNumber(tt.args.request, tt.args.streetNum)

			for i := 0; i < 100; i++ {
				tt.args.request <- true
				res := <-tt.args.streetNum
				if res <= 0 || res >= 701 {
					t.Errorf("TestRandomStreetNumber expect output to be between 1 and 700 but got %d\n", res)
				}
			}
		})
	}
}

func TestRecordsForWorkers(t *testing.T) {
	type args struct {
		totalRecords int
		workers      int
	}
	tests := []struct {
		name      string
		args      args
		wantFirst int
		wantRest  int
	}{
		{"11 records, 10 workers", args{11, 10}, 2, 1},
		{"100 records, 20 workers", args{100, 20}, 5, 5},
		{"1,000,000 records, 64 workers", args{1000000, 64}, 15625, 15625},
		{"1,000,000 records, 70 workers", args{1000000, 70}, 14335, 14285},
		{"77,777,777 records, 100 workers", args{77777777, 100}, 777854, 777777},
		{"500,000,000 records, 37 workers", args{500000000, 37}, 13513532, 13513513},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirst, gotRest := RecordsForWorkers(tt.args.totalRecords, tt.args.workers)
			if gotFirst != tt.wantFirst {
				t.Errorf("RecordsForWorkers() gotFirst = %v, want %v", gotFirst, tt.wantFirst)
			}
			if gotRest != tt.wantRest {
				t.Errorf("RecordsForWorkers() gotRest = %v, want %v", gotRest, tt.wantRest)
			}
		})
	}
}
