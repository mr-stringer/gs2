package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmcvetta/randutil"
)

//RandomSal accepts count (int) and sends data on the sal channel.
//The function will generate random weighted salutations and send them down the sal channel until it reaches the count threshold.
//When the count threshold is reached, the function will close the channel and return.
//If the process of randomising a sal fails, the function will shutdown the application.  This is not a nice
//way of handling the error and may be improved in a later version
func RandomSal(count int, sal chan<- string) {

	sals := GetWeightedSals()

	/*forever loop*/
	for i := 0; i < count; i++ {
		res, err := randutil.WeightedChoice(sals)
		if err != nil {
			log.Printf("Error drawing WeightedChoice")
			log.Print(err)
			// unclean exit
			os.Exit(-1)
		}
		sal <- fmt.Sprintf("%v", res.Item)
	}
	close(sal)
}

//RandomFirstIntial accepts count (int) and sends data on the fint channel.
//The function will generate random weighted initials and send them down the fint channel until it reaches the count threshold.
//When the count threshold is reached, the function will close the channel and return.
//If the process of randomising a fint fails, the function will shutdown the application.  This is not a nice
//way of handling the error and may be improved in a later version
func RandomFirstIntial(count int, fint chan<- string) {

	fints := GetWeightedFirstIntial()

	/*forever loop*/
	for i := 0; i < count; i++ {
		res, err := randutil.WeightedChoice(fints)
		if err != nil {
			log.Printf("Error drawing WeightedChoice")
			log.Print(err)
			// unclean exit
			os.Exit(-1)
		}
		fint <- fmt.Sprintf("%v", res.Item)

	}
	close(fint)
}

//RandomSurname accepts count (int) and sends data on the surname channel.
//The function will generate random weighted surnames and send them down the surname channel until it reaches the count threshold.
//When the count threshold is reached, the function will close the channel and return.
//If the process of randomising a surname fails, the function will shutdown the application.  This is not a nice
//way of handling the error and may be improved in a later version
func RandomSurname(count int, surname chan<- string) {

	surnames := GetWeightedSurnames()

	/*forever loop*/
	for i := 0; i < count; i++ {
		res, err := randutil.WeightedChoice(surnames)
		if err != nil {
			log.Printf("Error drawing WeightedChoice")
			log.Print(err)
			// unclean exit
			os.Exit(-1)
		}
		surname <- fmt.Sprintf("%v", res.Item)
	}
	close(surname)
}

//RandomStreetName accepts count (int) and sends data on the streetName channel.
//The function will generate random weighted street names and send them down the streetName channel until it reaches the count threshold.
//When the count threshold is reached, the function will close the channel and return.
//If the process of randomising a street name fails, the function will shutdown the application.  This is not a nice
//way of handling the error and may be improved in a later version
func RandomStreetName(count int, streetName chan<- string) {

	streetNames := GetWeightedStreetNames()

	/*forever loop*/
	for i := 0; i < count; i++ {
		res, err := randutil.WeightedChoice(streetNames)
		if err != nil {
			log.Printf("Error drawing WeightedChoice")
			log.Print(err)
			// unclean exit
			os.Exit(-1)
		}
		streetName <- fmt.Sprintf("%v", res.Item)
	}
	close(streetName)
}

//RandomTownName accepts count (int) and sends data on the townName channel.
//The function will generate random weighted town names and send them down the townName channel until it reaches the count threshold.
//When the count threshold is reached, the function will close the channel and return.
//If the process of randomising a town name fails, the function will shutdown the application.  This is not a nice
//way of handling the error and may be improved in a later version
func RandomTownName(count int, townName chan<- string) {

	townNames := GetWeightedTownNames()

	/*forever loop*/
	for i := 0; i < count; i++ {
		res, err := randutil.WeightedChoice(townNames)
		if err != nil {
			log.Printf("Error drawing WeightedChoice")
			log.Print(err)
			// unclean exit
			os.Exit(-1)
		}
		townName <- fmt.Sprintf("%v", res.Item)

	}
	close(townName)
}

//RandomDiscount accepts count (int) and sends data on the discount channel.
//The function will generate random weighted discount and send them down the discount channel until it reaches the count threshold.
//When the count threshold is reached, the function will close the channel and return.
//If the process of randomising a discount fails, the function will shutdown the application.  This is not a nice
//way of handling the error and may be improved in a later version
func RandomDiscount(count int, discount chan<- int) {

	discounts := GetWeightedDiscount()

	/*forever loop*/
	for i := 0; i < count; i++ {
		res, err := randutil.WeightedChoice(discounts)
		if err != nil {
			log.Printf("Error drawing WeightedChoice")
			log.Print(err)
			// unclean exit
			os.Exit(-1)
		}
		discount <- res.Item.(int)
	}
	close(discount)
}

//RandomStreetNumber accepts count (int) and sends data on the streetNum channel.
//The function will generate random weighted street number and send them down the streetNum channel until it reaches the count threshold.
//When the count threshold is reached, the function will close the channel and return.
//If the process of randomising a street number fails, the function will shutdown the application.  This is not a nice
//way of handling the error and may be improved in a later version
func RandomStreetNumber(count int, streetNum chan<- int) {
	Nums := []randutil.Choice{
		{Weight: 10, Item: 0}, //1 to 50
		{Weight: 7, Item: 1},  //51 to 100
		{Weight: 4, Item: 2},  //101 to 200
		{Weight: 2, Item: 3},  // 201 to 500
		{Weight: 1, Item: 4},  // 501 to 700
	}
	for i := 0; i < count; i++ { //forever
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
	}
	close(streetNum)
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
