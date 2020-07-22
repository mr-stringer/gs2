package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

type configuration struct {
	Hostname    string
	Port        string
	Username    string
	Password    string
	Schema      string
	DropSchema  bool
	Workers     int
	Customers   int
	Orders      int
	TrnxRecords int
	//InsertMode string Insert mode is no longer required.  Will always run using transactions
	StartYear int
	EndYear   int
	Verbose   bool
}

// getConfig the specified configuration file and attempts to marshall it's content into the
// configuration struct it is called from.
// getConfig will error if, the config file cannot be found, the json is malformed or if the marshalled configuration fails verification.
func (c *configuration) getConfig(cfile string) error {
	/*try abd open the config file*/
	file, err := os.Open(cfile)
	if err != nil {
		/*let the caller deal with the error*/
		return err
	}
	defer file.Close()
	log.Print("Config file opened\n")

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		return err
	}
	log.Printf("JSON Formatting is OK\n")
	/*verify the config*/
	err = c.verifyConfig()
	if err != nil {
		return err
	}
	log.Printf("Config file parsed OK\n")

	return nil
}

// verifyConfig checks that all required fields are correctly populated
// The only exception is DropSchema, if this is not set it defaults to `false`, which is the safe option
func (c *configuration) verifyConfig() error {
	/*Use reflection to get the values of the fields to ensure that each field is populated with something.*/
	v := reflect.ValueOf(*c)
	typeOfS := v.Type()

	var flagError = false

	for i := 0; i < v.NumField(); i++ {
		/*cast everthing to a string to test values*/
		s1 := fmt.Sprintf("%v", v.Field(i).Interface())
		if s1 == "" || s1 == "0" {
			/*Exclude TrnxRecords*/
			if typeOfS.Field(i).Name == "TrnxRecords" || typeOfS.Field(i).Name == "Verbose" {
				continue
			}
			log.Printf("Field: %s has no value set\n", typeOfS.Field(i).Name)
			flagError = true
		}
	}

	if flagError {
		log.Printf("One or more fields in the configuration was not set.  Ensure that the configuration file has all of the following parameters set:\n\tConfigName\n\tHostname\n\tPort\n\tUsername\n\tPassword\n\tSchema\n\tWorkers\n\tCustomers\n\tOrders\n\tInsertMode\n\tStartYear\n\tEndYear\n\nIf DropScema is omitted, it is set to false.")
		return errors.New("Incomplete configuration file")
	}

	/* Insert mode no longer support, large inserts will always been conducted as transactions.
	if c.InsertMode != "Batch" && c.InsertMode != "Single" {
		log.Printf("verifyConfig: Batchmode must be equal be set to Batch or Single, %s not supported", c.InsertMode)
		return errors.New("Batch mode incorrectly set")
	}
	*/

	/*check that start year is before or same as end year*/
	if c.StartYear > c.EndYear {
		log.Printf("verifyConfig: StartYear=%d, EndYear=%d! - StartYear must be equal or less than EndYear", c.StartYear, c.EndYear)
		return errors.New("StartYear and EndYear are not consistent")
	}

	/*Check username - SYSTEM not allowed*/
	if strings.ToUpper(c.Username) == "SYSTEM" {
		log.Printf("verifyConfig: The username SYSTEM is not allowed\n")
		return errors.New("Username SYSTEM not allowed")
	}

	/*Only allow years from 1001 to 3000*/
	if c.StartYear < 1001 || c.EndYear > 2999 {
		log.Printf("verifyConfig: Start Year must be 1001 or greater and End Year can be no greater than 2999")
		return errors.New("Date range not accepted")
	}

	/*If TrnxRecords is not set, set it to 10,000*/
	if c.TrnxRecords == 0 {
		c.TrnxRecords = 10000
	}

	/*Trnx must be between 100 and 10,000,000 check that it is*/
	if c.TrnxRecords > 10000000 || c.TrnxRecords < 100 {
		log.Printf("verifyConfig: TrnxRecords must cannot be lower than 100 or higher than 10,000,000")
		return errors.New("TrnxRecords not accepted")
	}

	return nil

}
