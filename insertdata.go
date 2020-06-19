package main

import (
	"log"

	"github.com/jmcvetta/randutil"
)

//RandomProductID supplies weighted random prodcuts IDs when requested
func RandomProductID(products []randutil.Choice, ID chan<- int, quit <-chan bool, wid int) {
	for {
		select {
		/*if quit is sent!*/
		case <-quit:
			/*clean up*/
			return
		default:
			/*if no quit is sent check if there is room in the channel for another ID*/
			if len(ID) < cap(ID) {
				IDInterface, err := randutil.WeightedChoice(products)
				if err != nil {
					log.Fatalf("WORKER-%d: fails to draw random product ID this will cause gs2 to exit", wid)
				}
				IDVal, ok := IDInterface.Item.(int)
				if !ok {
					log.Fatalf("WORKER-%d: fails to draw cast product ID to int, this will cause gs2 to exit", wid)
				}
				/*stick it down the pipe.*/
				ID <- IDVal
			}
		}
	}
}

//RandomCustomerID supplies weighted random prodcuts IDs when requested
func RandomCustomerID(customers []int, ID chan<- int, quit <-chan bool, wid int) {
	for {
		select {
		/*if quit is sent!*/
		case <-quit:
			/*clean up*/
			return
		default:
			rndel := RandInt(0, len(customers)-1)
			ID <- customers[rndel]
		}
	}
}
