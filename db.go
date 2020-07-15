package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	_ "github.com/SAP/go-hdb/driver" /*register driver*/
	"github.com/jmcvetta/randutil"
)

type gsConn struct {
	Conn      *sql.DB
	Connected bool /*defaults to false which is what I want :) */
}

//InsertPayload provides the data required to insert an order
type InsertPayload struct {
	CustomerID int
	ProductID  int
	Date       time.Time
}

// Creates a connection to the database thus populating the Conn struct.
func (g *gsConn) Init(c configuration) error {
	log.Print("Initalising connection")
	err := g.Connect(g.PrintDsn(c.Username, c.Password, c.Hostname, c.Port))
	if err != nil {
		log.Panicf("Could not initialise database connection.  Probably wrong details or DB is not up\n")
		return err
	}
	return nil
}

// printDsn is used via the Connect function, it's simply prints the formatted DSN string
func (g *gsConn) PrintDsn(username, password, hostname, port string) string {
	return fmt.Sprintf("hdb://%s:%s@%s:%s", username, password, hostname, port)
}

// Connect should be called only from Init and is used to populate the Conn variable
func (g *gsConn) Connect(Dsn string) error {
	var err error
	g.Conn, err = sql.Open("hdb", Dsn)
	if err != nil {
		return err
	}
	g.Connected = true
	return nil

}

// GetHanaVersion returns a string containing the version of HANA
func (g gsConn) GetHanaVersion() (string, error) {
	var version string
	var q1 string
	q1 = "SELECT VERSION FROM \"SYS\".\"M_DATABASE\""
	r1 := g.Conn.QueryRow(q1)
	err := r1.Scan(&version)
	return version, err
}

// Checks that the user givene in the argument 'user' has been granted the 'MONITORING' role.  Returns true if the role is granted.UserHasMonRole
// Returns an error if an error occurs.
func (g gsConn) UserHasMonRole(user string) (bool, error) {
	var count int = 0
	//count = 0
	q1 := fmt.Sprintf("SELECT COUNT(GRANTEE) FROM \"SYS\".\"GRANTED_ROLES\" WHERE GRANTEE = '%s' AND ROLE_NAME = 'MONITORING'", strings.ToUpper(user))
	r1 := g.Conn.QueryRow(q1)
	err := r1.Scan(&count)
	if err != nil {
		log.Printf("Scan Failed\n")
		log.Printf("%s\n", err)
		return false, err
	}
	if count > 0 {
		log.Printf("MONITORING role not gtanted to %s\n", user)
		return true, nil
	}
	log.Printf("MONITORING role is gtanted to %s\n", user)
	return false, nil
}

// Checks if the schema in the argument 'schema' is present.  Returns true if the schema is present, otherwises it returns false.CheckSchema
// Returns an error if an error occurs
func (g gsConn) CheckSchema(schema string) (bool, error) {
	q1 := fmt.Sprintf("SELECT COUNT(SCHEMA_NAME) FROM \"PUBLIC\".\"SCHEMAS\" WHERE SCHEMA_NAME = '%s'", schema)
	count := 0
	r1 := g.Conn.QueryRow(q1)
	err := r1.Scan(&count)
	if err != nil {
		log.Printf("Scan failed\n")
		return false, err
	}
	if count > 0 {
		log.Printf("\"%s\" schema is present\n", schema)
		return true, nil
	}
	log.Printf("\"%s\" schema is not present\n", schema)
	return false, nil
}

// Drops the schema giving in the argument 'schema'.  Returns an error if the DROP fails.
func (g *gsConn) DropSchema(schema string) error {
	log.Printf("Dropping schema \"%s\"\n", schema)
	q1 := fmt.Sprintf("DROP SCHEMA \"%s\" CASCADE", schema)
	_, err := g.Conn.Exec(q1)
	if err != nil {
		log.Printf("Schema drop failed\n")
		return err
	}
	log.Printf("Schema \"%s\" dropped\n", schema)
	return nil
}

// Create the schema, tables and views
func (g *gsConn) CreateSchema(schema string, drop bool) error {
	log.Printf("Creating Schema \"%s\"\n", schema)

	exists, err := g.CheckSchema(schema)
	if err != nil {
		return err
	}

	if exists && drop {
		err = g.DropSchema(schema)
		if err != nil {
			return err
		}
	} else if exists && !drop {
		log.Printf("Schema \"%s\" already exists, remove it manually or set DropSchema to true in the configuration file\n", schema)
		return fmt.Errorf("error - cannot remove schema")
	}

	err = g.TransactExecRows(GetSchemaStatements(schema))
	if err != nil {
		return err
	}
	log.Printf("Schema \"%s\" created\n", schema)
	return nil
}

// Populates the schema with masterdata
func (g gsConn) CreateMasterData(schema string) error {

	log.Printf("Inserting master data")
	err := g.TransactExecRows(GetMasterDataStatements(schema))
	if err != nil {
		return err
	}
	log.Printf("Master data inserted")
	return nil
}

// Inserts products into the schema
func (g gsConn) InsertProducts(schema string) error {
	log.Printf("Inserting product data")
	err := g.TransactExecRows(GetProductStatements(schema))
	if err != nil {
		return err
	}
	log.Printf("Product data inserted")
	return nil
}

//TransactRows takes a single argument, a slice of strings.  The will all be passed to the database uning the Exec function in a single transaction
// if this fails for any reason, the function will return an error.  The log will contain information about the error.  If an error occurs it
// is unlikely that the program should continue.
func (g gsConn) TransactExecRows(statements []string) error {
	var trnxError error

	trnx, err := g.Conn.Begin()
	if err != nil {
		log.Printf("Failed to start transaction\n")
		return err
	}

	for _, i := range statements {
		/*log.Printf("Inserting: %s\n", i)*/
		_, trnxError = trnx.Exec(i)
		if trnxError != nil {
			log.Printf("Error executing: %s\n", i)
			break
		}
	}

	if trnxError != nil {
		err = trnx.Rollback()
		if err != nil {
			log.Printf("Failed to rollback transaction - database could be inconsistent\n")
			return err
		}
		fmt.Printf("Sucessfully rolledback the failed transaction\n")
		return trnxError
	}

	err = trnx.Commit()
	if err != nil {
		log.Printf("Failed to commit transaction - database could be inconsistent\n")
		return err
	}

	return nil
}

//WorkerInsertCustomers is design to work as a go rountine as part of a weight group
// wid should be a unique int representing the Worker ID
// count is the number of customers to insert
// schema is the schema in which to inser the customers
// wg is a pointer to the waitgroup to report to
func (g gsConn) WorkerInsertCustomers(wid, count int, schema string, wg *sync.WaitGroup) {
	var execError error
	//var rollbackErr error

	//log.Printf("WORKER-%d: Inserting %d random customers\n", wid, count)
	trnx, err := g.Conn.Begin()
	if err != nil {
		log.Printf("WORKER-%d: Failed to start transaction", wid)
		/*I'll write a better way of handling errors at some point*/
		log.Printf("WORKER-%d: Forcing application to quit, this won't be pretty", wid)
		os.Exit(-1)
	}

	/*make the channels for the random functions*/
	sal := make(chan string)
	fint := make(chan string)
	surname := make(chan string)
	streetnum := make(chan int)
	streetname := make(chan string)
	townname := make(chan string)
	disco := make(chan int)

	/*fire off the go routines*/
	go RandomSal(count, sal)
	go RandomFirstIntial(count, fint)
	go RandomSurname(count, surname)
	go RandomStreetNumber(count, streetnum)
	go RandomStreetName(count, streetname)
	go RandomTownName(count, townname)
	go RandomDiscount(count, disco)

	/*loops for the value giving in 'count'.  It starts by requesting random data and then inserts it.*/
	for i := 0; i < count; i++ {
		stmt1 := fmt.Sprintf("INSERT INTO \"%s\".\"CUSTOMERS\" (SAL, FNAME, LNAME, ADDR1, CITY, DISCOUNT_PCT) VALUES ('%s', '%s', '%s', '%d %s', '%s', '%d');", schema, <-sal, <-fint, <-surname, <-streetnum, <-streetname, <-townname, <-disco)

		_, execError = trnx.Exec(stmt1)
		if execError != nil {
			log.Printf("WORKER-%d: Failed to execute \"%s\"\n", wid, stmt1)
			break
		}
	}

	/*Deal with the DB state*/
	if execError != nil {
		log.Printf("WORKER-%d: An insert failed, attempting rollback", wid)
		log.Print(execError)
		err = trnx.Rollback()
		if err != nil {
			log.Printf("WORKER-%d: Rollback failed, database could be inconsistent", wid)
			log.Printf("WORKER-%d: Forcing program-wide exit", wid)
			os.Exit(-1)
		}
	}
	err = trnx.Commit()
	if err != nil {
		log.Printf("WORKER-%d: Commit failed, database could be inconsistent", wid)
		log.Printf("WORKER-%d: Forcing program-wide exit", wid)
		os.Exit(-1)

	}
	//log.Printf("WORKER-%d: Commit sucessful\n", wid)
	wg.Done()
	return
}

//GetProductsIDs returns a slice of Choice which is used to draw weighted random product IDs
func (g gsConn) GetProductIDs(schema string) ([]randutil.Choice, error) {
	q1 := fmt.Sprintf("SELECT COUNT(ID) FROM \"%s\".\"PRODUCT\"", schema)
	r := g.Conn.QueryRow(q1)
	var count int
	err := r.Scan(&count)
	if err != nil {
		log.Printf("GetProductIDs:Error scanning result: %s\n", q1)
		return []randutil.Choice{}, err
	}

	productIDs := make([]randutil.Choice, count)

	q2 := fmt.Sprintf("SELECT ID, RAND_WEIGHT FROM \"%s\".\"PRODUCT\"", schema)
	rows, err := g.Conn.Query(q2)
	if err != nil {
		log.Printf("GetProductIDs:Error running query: %s\n", q1)
		return productIDs, err
	}
	defer rows.Close()

	c1 := 0
	var id int
	var weight int
	for rows.Next() {
		err = rows.Scan(&id, &weight)
		if err != nil {
			log.Printf("GetProductIDs:Error scanning result: %s\n", q2)
			return productIDs, err
		}
		productIDs[c1] = randutil.Choice{Weight: weight, Item: id}
		c1++
	}
	return productIDs, nil

}

//GetCustomerIDs returns a slice of customer IDs which is used to randomise customers
func (g gsConn) GetCustomerIDs(schema string) ([]int, error) {
	log.Printf("GetCustomerIDs: Getting customers IDs from the database")
	q1 := fmt.Sprintf("SELECT COUNT(ID) FROM \"%s\".\"CUSTOMERS\"", schema)
	r := g.Conn.QueryRow(q1)
	var count int
	err := r.Scan(&count)
	if err != nil {
		log.Printf("GetCustomerIDs:Error scanning result: %s\n", q1)
		return []int{}, err
	}

	CustomerIDs := make([]int, count)
	q2 := fmt.Sprintf("SELECT ID FROM \"%s\".\"CUSTOMERS\"", schema)
	rows, err := g.Conn.Query(q2)
	if err != nil {
		log.Printf("GetCustomerIDs:Error running query: %s\n", q2)
		return []int{}, errors.New("Found an error")
	}
	defer rows.Close()

	c1 := 0
	var id int
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Printf("GetProductIDs:Error scanning result: %s\n", q2)
			return CustomerIDs, err
		}
		CustomerIDs[c1] = id
		c1++
	}
	return CustomerIDs, nil

}

//CreatePayload is a function that runs as a goroutines and creates the random data for the customer inserts
func (g gsConn) CreatePayload(c configuration, plChan chan<- InsertPayload) {
	log.Printf("CreatePayload: Getting Product IDs\n")
	prodIDs, err := g.GetProductIDs(c.Schema)
	if err != nil {
		log.Printf("CreatePayload: failed to get product IDs\n")
		log.Printf("CreatePayload: will now quit and not clean up\n")
		/*ugly quit*/
		os.Exit(-1)
	}

	log.Printf("CreatePayload: Customer Product IDs\n")
	CustIDs, err := g.GetCustomerIDs(c.Schema)
	if err != nil {
		log.Printf("CreatePayload: failed to get customer IDs\n")
		log.Printf("CreatePayload: will now quit and not clean up\n")
		/*ugly quit*/
		os.Exit(-1)
	}

	log.Printf("CreatePayload: Making Channels")
	/*Make channels for random products*/
	rndProd := make(chan int, c.Workers*2) /*bueffered channel length = workers*2*/

	/*Make channels for random customers*/
	rndCust := make(chan int, c.Workers*2) /*bueffered channel length = workers*2*/

	/*Make channels for random dates*/
	rndDate := make(chan time.Time, c.Workers*2) /*bueffered channel length = workers*2*/

	log.Printf("CreatePayload: Stating goroutines")
	go RandomProductID(c.Orders, prodIDs, rndProd)
	go RandomCustomerID(c.Orders, CustIDs, rndCust)
	go RandomDate(c, rndDate)

	for { /*forever*/
		if len(plChan) < cap(plChan) {
			plChan <- InsertPayload{CustomerID: <-rndCust, ProductID: <-rndProd, Date: <-rndDate}
		}
	}
}

//PlaceOrders is used to place orders in the database.  The workers element of the configuration struct decides how many instances of the function will run
//It takes a number of arguments.
//configuration is the gs2 configuration struct
//count is the total number of orders for this instance to insert
//wid is the worker id, this should be unique
//errchan is a channel where the main process can check for any errors being sent back.
//A transaction size of 10,000 records is currently in place, this will be made variable in the future
func (g gsConn) PlaceOrders(c configuration, count int, wid int, retchan chan<- chanReturn, payloadChan <-chan InsertPayload) {
	/*Loop until we run out of things to do*/
	var loopMax int = 10000
	var done int = 0
	log.Printf("WORKER-%d: Entering outer loop", wid)
	for remaining := count; remaining > 0; {
		log.Printf("WORKER-%d: %d remaining", wid, remaining)

		var loop int
		if remaining > loopMax {
			loop = loopMax
		} else {
			loop = remaining
		}
		log.Printf("WORKER-%d: Attempting to start transaction of %d orders", wid, loop)

		trnx, err := g.Conn.Begin()
		if err != nil {
			log.Printf("WORKER-%d: failed to start transaction", wid)
			log.Printf("WORKER-%d: Will notify main and quit", wid)
			retchan <- chanReturn{ok: false, message: err.Error()}
			return
		}

		st1 := fmt.Sprintf("CALL \"%s\".\"PLACE_ORDER\"(?,?,?,1)", c.Schema)
		stmt, err := trnx.Prepare(st1)
		if err != nil {
			log.Printf("WORKER-%d: failed to prepare statement", wid)
			log.Printf("WORKER-%d: Will notify main and quit", wid)
			retchan <- chanReturn{ok: false, message: err.Error()}
			return
		}

		for i := 0; i < loop; i++ {
			pl := <-payloadChan
			/*start transacting*/

			_, err = stmt.Exec(pl.CustomerID, pl.ProductID, pl.Date)
			if err != nil {
				log.Printf("WORKER-%d: failed to execute statement, will attempt rollback", wid)
				rberr := trnx.Rollback()
				if rberr != nil {
					log.Printf("WORKER-%d: rollback failed!!!", wid)
					log.Printf("WORKER-%d: Will notify main and quit", wid)
					retchan <- chanReturn{ok: false, message: rberr.Error()}
					return
				}
				log.Printf("WORKER-%d: rollback OK", wid)
				log.Printf("WORKER-%d: Will notify main and quit", wid)
				retchan <- chanReturn{ok: false, message: err.Error()}
				return
			}
		}
		/*Close the statement*/
		err = stmt.Close()
		if err != nil {
			log.Printf("WORKER-%d: failed to close statement", wid)
			log.Printf("WORKER-%d: Will notify main and quit", wid)
			retchan <- chanReturn{ok: false, message: err.Error()}
			return
		}

		/*Commit*/
		err = trnx.Commit()
		if err != nil {
			log.Printf("WORKER-%d: failed to commit transaction", wid)
			log.Printf("WORKER-%d: Will notify main and quit", wid)
			retchan <- chanReturn{ok: false, message: err.Error()}
			return
		}
		done += loop
		remaining = remaining - loop
		log.Printf("WORKER-%d: Committed %.2f%% of orders", wid, float32(done)/float32(count)*100)

	}
	retchan <- chanReturn{ok: true, message: fmt.Sprintf("WORKER-%d: Completed all orders", wid)}
}
