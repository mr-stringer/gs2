package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmcvetta/randutil"
)

//RandomSal accepts a bool on the request channel and sends data on the sal channel.
//Sending 'true' on the request channel will trigger a random salutation to be send on the sal channel.
//Sending 'false' on the request channel will close the sal channel and end the function.
//If the process of randomising a sal fails, the function will shutdown the application.  This is not a nice
//way of handling the error and may be improved in a later version
func RandomSal(request <-chan bool, sal chan<- string) {

	sals := GetWeightedSals()

	/*forever loop*/
	for {
		r := <-request
		if r == true {
			res, err := randutil.WeightedChoice(sals)
			if err != nil {
				log.Printf("Error drawing WeightedChoice")
				log.Print(err)
				// unclean exit
				os.Exit(-1)
			}
			sal <- fmt.Sprintf("%v", res.Item)
		} else {
			/*request to close*/
			return
		}
	}
}

//RandomFirstIntial accepts a bool on the request channel and sends data on the fint channel.
//Sending 'true' on the request channel will trigger a random uppercase character fint channel.
//Sending 'false' on the request channel will end the function.
//If the process of randomising a fint fails, the function will shutdown the application.  This is not a nice
//way of handling the error and may be improved in a later version
func RandomFirstIntial(request <-chan bool, fint chan<- string) {

	fints := GetWeightedFirstIntial()

	/*forever loop*/
	for {
		r := <-request
		if r == true {
			res, err := randutil.WeightedChoice(fints)
			if err != nil {
				log.Printf("Error drawing WeightedChoice")
				log.Print(err)
				// unclean exit
				os.Exit(-1)
			}
			fint <- fmt.Sprintf("%v", res.Item)
		} else {
			/*request to close*/

			return
		}
	}
}

//RandomSurname accepts a bool on the request channel and sends data on the surname channel.
//Sending 'true' on the request channel will trigger a random surname to be sent on the surname channel.
//Sending 'false' on the request channel will end the function.
//If the process of randomising a surname fails, the function will shutdown the application.  This is not a nice
//way of handling the error and may be improved in a later version
func RandomSurname(request <-chan bool, surname chan<- string) {

	surnames := GetWeightedSurnames()

	/*forever loop*/
	for {
		r := <-request
		if r == true {
			res, err := randutil.WeightedChoice(surnames)
			if err != nil {
				log.Printf("Error drawing WeightedChoice")
				log.Print(err)
				// unclean exit
				os.Exit(-1)
			}
			surname <- fmt.Sprintf("%v", res.Item)
		} else {
			/*request to close*/
			return
		}
	}
}

//RandomStreetName accepts a bool on the request channel and sends data on the streetName channel.
//Sending 'true' on the request channel will trigger a random street name to be sent on the streetName channel.
//Sending 'false' on the request channel will end the function.
//If the process of randomising a street name fails, the function will shutdown the application.  This is not a nice
//way of handling the error and may be improved in a later version
func RandomStreetName(request <-chan bool, streetName chan<- string) {

	streetNames := GetWeightedStreetNames()

	/*forever loop*/
	for {
		r := <-request
		if r == true {
			res, err := randutil.WeightedChoice(streetNames)
			if err != nil {
				log.Printf("Error drawing WeightedChoice")
				log.Print(err)
				// unclean exit
				os.Exit(-1)
			}
			streetName <- fmt.Sprintf("%v", res.Item)
		} else {
			/*request to close*/
			return
		}
	}
}

//RandomTownName accepts a bool on the request channel and sends data on the town channel.
//Sending 'true' on the request channel will trigger a random town name to be sent on the townName channel.
//Sending 'false' on the request channel will end the function.
//If the process of randomising a town name fails, the function will shutdown the application.  This is not a nice
//way of handling the error and may be improved in a later version
func RandomTownName(request <-chan bool, townName chan<- string) {

	townNames := GetWeightedTownNames()

	/*forever loop*/
	for {
		r := <-request
		if r == true {
			res, err := randutil.WeightedChoice(townNames)
			if err != nil {
				log.Printf("Error drawing WeightedChoice")
				log.Print(err)
				// unclean exit
				os.Exit(-1)
			}
			townName <- fmt.Sprintf("%v", res.Item)
		} else {
			/*request to close*/
			return
		}
	}
}

//RandomDiscount accepts a bool on the request channel and sends data on the discount channel.
//Sending 'true' on the request channel will trigger a random discount % to be sent on the discount channel.
//Sending 'false' on the request channel will end the function.
//If the process of randomising a discount fails, the function will shutdown the application.  This is not a nice
//way of handling the error and may be improved in a later version
func RandomDiscount(request <-chan bool, discount chan<- int) {

	discounts := GetWeightedDiscount()

	/*forever loop*/
	for {
		r := <-request
		if r {
			res, err := randutil.WeightedChoice(discounts)
			if err != nil {
				log.Printf("Error drawing WeightedChoice")
				log.Print(err)
				// unclean exit
				os.Exit(-1)
			}
			discount <- res.Item.(int)
		} else {
			/*request to close*/
			return
		}
	}
}

//RandomStreetNumber accepts a bool on the request channel and sends data on the streetNum channel.
//Sending 'true' on the request channel will trigger a random street number to be sent on the streetNum channel.
//Sending 'false' on the request channel will end the function.
//If the process of randomising a street number fails, the function will shutdown the application.  This is not a nice
//way of handling the error and may be improved in a later version
func RandomStreetNumber(request <-chan bool, streetNum chan<- int) {
	Nums := []randutil.Choice{
		{Weight: 10, Item: 0}, //1 to 50
		{Weight: 7, Item: 1},  //51 to 100
		{Weight: 4, Item: 2},  //101 to 200
		{Weight: 2, Item: 3},  // 201 to 500
		{Weight: 1, Item: 4},  // 501 to 700
	}
	for { //forever
		r := <-request
		if r == true {
			res, err := randutil.WeightedChoice(Nums)
			if err != nil {
				log.Printf("Error drawing WeightedChoice")
				log.Print(err)
				//exit from here
				os.Exit(-1)
			}
			sn := res.Item.(int)
			switch sn {
			case 0:
				streetNum <- RandInt(1, 50)
			case 1:
				streetNum <- RandInt(51, 100)
			case 2:
				streetNum <- RandInt(101, 200)
			case 3:
				streetNum <- RandInt(201, 500)
			case 4:
				streetNum <- RandInt(501, 700)
			default:
				log.Printf("Not sure what to do")
				os.Exit(-1)
			}
		} else {
			//log.Printf("I've been asked to quit\n")
			return

		}
	}
}

//RecordsForWorkers is used to calculate how many records each worker should get
//In most instances, the number of records will not be able to be evenly split across workers.
//RecordsForWorkers take two integer arguments. totalRecords is the total number of records to insert
//workers is the number of workers.
//If the all workers cannot recieve an equal number of records, the first worker will be given the difference
//The returned value 'first' will be the value to be given to the first worker, the value 'rest' is to be given to all other workers.
func RecordsForWorkers(totalRecords, workers int) (first, rest int) {
	//see if the records can be equally split
	mod := totalRecords % workers
	if (mod) == 0 {
		return (totalRecords / workers), (totalRecords / workers)
	}

	return mod + (totalRecords / workers), (totalRecords / workers)
}
