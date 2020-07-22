# gs2

gs2 (guitar shop 2) is tool that creates a testing workload for the SAP HANA database.  gs2 simulates intensive writes to a very simple normalised OLTP schema for a fictional guitar shop.

gs2 can be used to help with:

* Benchmarking performance
* Stress testing
* Understanding and validating HANA configuration changes and tuning

gs2 is written in go and therefore is statically compiled and requires virtually no dependacies, making it highly portable and  perfect for quick testing.

**Warning - gs2 can delete data inside your HANA database!**

Before using gs2, make sure you understand what it will do.  A certain combination of configuration parameters could delete your data.  Carefully read the configuration section before running!

## What gs2 does

gs2 is a tool for creating a test load on a HANA database.  The sequence of events is as follows.

* Reads and validates configuration files
* Connects to the target database
* Creates a schema
* Populates master data
* Generates orders

Generating orders is the bulk of the gs2 test workload.  gs2 will spawn worker threads that work in parallel provide weighted-random data for the PLACE_ORDER procudure in the database.  This procdure then populates data in the SALES and SALE_PRODUCTS tables.  The schema contains the view V_ORDERS which can be used for analytic testing.  It's hoped that updates to gs2 will provide inbuilt analytic workload testing.

### Weighted-random vs normal random

One of the goals of gs2 was to generate data that looks like normal data.  Over a long enough period, random data tends to look flat and doesn't make it very interesting to look at.  Therefore, gs2 uses weightings that make some products more popular than others.  Customers also have names that are more common than others and the towns that they live in are also weighted.  All of this gives the data a more 'real' feel.  The random data generation is done in gs2 as not to add artificial load on the database.

## Testing and Compiling

To compile gs2, you'll need a working golang environemnt.  For details of setting this up see <https://golang.org/doc/install>

You can run the test-suite by issuing the command

```golang
go test
```

gs2 can be compiled with the command

```golang
go install
```

This will place the gs2 binary in your `$GOBIN` directory.

## Prerequistes

In order to run gs2, you must have:

* A HANA database
* A client system on which gs2 will run (optional, but preferred)
* A user configured with the right permissions
* The SQL port of the database

gs2 has only be tested against HANA 2.0 SPS4 and newer.  It is not known if it will run against lower version of HANA.

It is preferred to run gs2 on a different host to the database itself.  This is so the load created by gs2 does not effect the database.

gs2 will not allow you to use the SYSTEM user, as this would be terrible security practice.  A non-SYSTEM user must be specified.  The new user must have the MONITOR role and the CREATE SCHEMA privilege.  The following SQL shows how these can be applied to a user:

```SQL
GRANT MONITORING TO TEST_USER;
GRANT CREATE SCHEMA TO TEST_USER;
```

gs2 will need to know the SQL port to use.  The best way to find the port is to connect to the tenant which used with gs2 and run the following query:

```SQL
SELECT SQL_PORT FROM M_SERVICES WHERE SQL_PORT != 0
```

The result should look similar to this:

```SQL
SQL_PORT
30041
```

You should not run gs2 against the SYSTEMDB.  gs2 will not stop you doing this, but it is unwise and untested.

## Running gs2

gs2 accepts no command line arguments and instead requires a JSON formatted configuration file named gs2.json to be present in the current working directory.  The configuration parameters are as follows:

| Parameter   | Type    | Description  |
|-------------|---------|--------------|
| Hostname    | String  | The hostname or IP address of the target HANA system.|
| Port        | Integer | The SQL port of the target HANA DB.|
| Username   | String  | The HANA username to be used on the target HANA system.|
| Password    | String  | The password for the user.|
| Schema      | String  | The schema to be used.|
| DropSchema  | Boolean | If set to `true`, gs2 will drop the target schema if it exists.  If set to `false`, gs2 will quit if the schema exists.  Only use `true` with caution.  If omitted, DropSchema will default to false.|
| Workers     | Integer | The number of worker goroutines to spawn.|
| Customers   | Integer | The number of customers to create.|
| Orders      | Integer | The number of orders to create.|
| TrnxRecords | Integer | Optional.  The number of records to be inserted in each transation.  Default is 10,000.  The lowest acceptable value is 100.  The highest acceptable is 10,000,000. |
| StartYear   | Integer | The year that orders start, lowest accepatable value is 1001.|
| EndYear     | Integer | The year that orders end, the highest acceptable value is 2999|
| Verbose     | Boolean | When set to `true`, gs2 will log the comitting of every transaction.  Verbose is set to `false` by default.|

Below is an example configuration file which is correctly formatted

```json
{
    "Hostname":"test_host",
    "Port":"30015",
    "Username":"TEST_USER",
    "Password":"HardPassword",
    "Schema":"GtrShop",
    "Workers":10,
    "Customers":100,
    "Orders"   :10000,
    "StartYear":2010,
    "EndYear":2020
}
```

With a valid configuration file in place, gs2 can be launched from the command line.

## Output

gs2 is very chatty and will produce a lot of output to screen.  Any problems found should be printed to screen too in (hopefully) plain English.

## Hints and Tips

Tune the configuration file for best performance.  Start with a low number of workers, customers and orders.  A resonable starting point would be 8 workers, 1000 customers and 100,000 orders.

Performance may be constrianed by either gs2 (high CPU use on the gs2 host), the HANA database (high CPU or disk throuhput on the DB host) or even network saturation.  Use the system monitoring tools at your disposal to help figure out the bottleneck.
