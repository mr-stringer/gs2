package main

import (
	"log"

	"github.com/jmcvetta/randutil"
)

//RandomProductID supplies weighted random prodcuts IDs when requested
func RandomProductID(products []randutil.Choice, ID chan<- int, quit <-chan bool) {
	for {
		/*Just stick a value in the pipe and block if it is full*/
		IDInterface, err := randutil.WeightedChoice(products)
		if err != nil {
			log.Fatalf("RandomProductID: fails to draw random product ID this will cause gs2 to exit")
		}
		IDVal, ok := IDInterface.Item.(int)
		if !ok {
			log.Fatalf("RandomProductID: fails to draw cast product ID to int, this will cause gs2 to exit")
		}
		/*stick it down the pipe.*/
		ID <- IDVal
	}
}

//RandomCustomerID supplies weighted random prodcuts IDs when requested
func RandomCustomerID(customers []int, ID chan<- int, quit <-chan bool) {
	for {
		/*Just stick a value in the pipe and block if it is full*/
		rndel := RandInt(0, len(customers)-1)
		ID <- customers[rndel]
	}
}
