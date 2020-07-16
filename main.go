package main

import (
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	/*let us begin*/
	start := time.Now()
	log.Printf("gs2 is initialising")

	/*Read in the config*/
	c := configuration{}
	err := c.getConfig("gs2.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	/*Create the DB struct and initilise it*/
	db := gsConn{}
	err = db.Init(c)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Conn.Close()

	err = db.Conn.Ping()
	if err != nil {
		log.Fatalf("Can't ping DB, can't continue\n")
	}

	/*Create the schema*/
	err = db.CreateSchema(c.Schema, c.DropSchema)
	if err != nil {
		log.Printf("Failed to create the schema\n")
		log.Print(err)
		log.Printf("gs2 failed to complete the all tasks\n")
		os.Exit(-1)
	}

	/*Create the master data*/
	err = db.CreateMasterData(c.Schema)
	if err != nil {
		log.Printf("Failed to populate masterdata\n")
		log.Print(err)
		log.Printf("gs2 failed to complete the all tasks\n")
		os.Exit(-1)
	}

	/*Insert products*/
	err = db.InsertProducts(c.Schema)
	if err != nil {
		log.Printf("Failed to populate masterdata\n")
		log.Print(err)
		log.Printf("gs2 failed to complete the all tasks\n")
		os.Exit(-1)
	}

	/*Insert customers*/
	log.Printf("Spawning %d workers to create %d customer\n", c.Workers, c.Customers)
	first, rest := RecordsForWorkers(c.Customers, c.Workers)
	var wg1 sync.WaitGroup
	for i := 0; i < c.Workers; i++ {
		wg1.Add(1)
		if i == 0 {
			go db.WorkerInsertCustomers(i, first, c.Schema, &wg1)
		} else {
			go db.WorkerInsertCustomers(i, rest, c.Schema, &wg1)
		}
	}
	wg1.Wait()
	log.Printf("All customers created\n")

	/*Create the return channel*/
	retchan := make(chan chanReturn)

	/*Create the payload chan and start the routine*/
	pl := make(chan InsertPayload, c.Workers)
	go db.CreatePayload(c, pl)

	InsStart := time.Now()
	/*Insert customers*/
	log.Printf("Spawning %d workers to create %d orders\n", c.Workers, c.Orders)
	first, rest = RecordsForWorkers(c.Orders, c.Workers)

	for i := 0; i < c.Workers; i++ {
		if i == 0 {
			go db.PlaceOrders(c, first, i, retchan, pl)
		} else {
			go db.PlaceOrders(c, rest, i, retchan, pl)
		}
	}

	var (
		ok     int
		failed int
	)

	for i := 0; i < c.Workers; i++ {
		ret := <-retchan
		if !ret.ok {
			failed++
			log.Printf("MAIN: Worker failed:%s\n", ret.message)
		} else {
			log.Printf("MAIN: Worker OK:%s\n", ret.message)
			ok++
		}
	}
	InsEnd := time.Now()
	End := time.Now()

	close(pl)

	log.Printf("========================\n")
	log.Printf("Workers completed: %d\n", ok)
	log.Printf("Workers failed:    %d\n", failed)
	log.Printf("========================\n")
	log.Printf("Orders took %.2f seconds\n", InsEnd.Sub(InsStart).Seconds())
	log.Printf("Total time %.2f seconds\n", End.Sub(start).Seconds())
	os.Exit(0)
}
