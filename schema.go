package main

import (
	"fmt"
)

// GetSchemaStatements provides, in order, the slice of strings that can be used to create the gs2 shcemea
// the argument 'schema' sets the correct schmea name in the statements
func GetSchemaStatements(schema string) []string {
	statements := []string{}
	statements = append(statements, fmt.Sprintf("CREATE SCHEMA \"%s\"", schema))
	statements = append(statements, fmt.Sprintf("CREATE TABLE \"%s\".\"PRODUCT_TYPE\" (ID INTEGER PRIMARY KEY NOT NULL GENERATED BY DEFAULT AS IDENTITY, TYPE_NAME NVARCHAR(8) UNIQUE, DESCRIPTION NVARCHAR(1024))", schema))
	statements = append(statements, fmt.Sprintf("CREATE TABLE \"%s\".\"PRODUCT_COLOUR\" (ID INTEGER PRIMARY KEY NOT NULL GENERATED BY DEFAULT AS IDENTITY,COLOUR NVARCHAR(64) UNIQUE)", schema))
	statements = append(statements, fmt.Sprintf("CREATE TABLE \"%s\".\"MANUF\" (ID INTEGER PRIMARY KEY NOT NULL GENERATED BY DEFAULT AS IDENTITY,MANUF NVARCHAR(64) UNIQUE)", schema))
	statements = append(statements, fmt.Sprintf("CREATE TABLE \"%s\".\"PRODUCT\" (ID INTEGER PRIMARY KEY NOT NULL GENERATED BY DEFAULT AS IDENTITY,NAME NVARCHAR(64) NOT NULL,TYPE INTEGER NOT NULL,MANUF INTEGER NOT NULL,COLOUR INTEGER,BUY_PRICE DECIMAL,SELL_PRICE DECIMAL, RAND_WEIGHT INTEGER)", schema))
	statements = append(statements, fmt.Sprintf("ALTER TABLE \"%s\".\"PRODUCT\" ADD CONSTRAINT FK_PRD_TYPE FOREIGN KEY (TYPE) REFERENCES \"%s\".\"PRODUCT_TYPE\" (ID) ON DELETE CASCADE", schema, schema))
	statements = append(statements, fmt.Sprintf("ALTER TABLE \"%s\".\"PRODUCT\" ADD CONSTRAINT FK_PRD_COLOUR FOREIGN KEY (COLOUR) REFERENCES \"%s\".\"PRODUCT_COLOUR\" (ID) ON DELETE CASCADE", schema, schema))
	statements = append(statements, fmt.Sprintf("ALTER TABLE \"%s\".\"PRODUCT\" ADD CONSTRAINT FK_PRD_MANUF FOREIGN KEY (MANUF) REFERENCES \"%s\".\"MANUF\" (ID) ON DELETE CASCADE", schema, schema))
	statements = append(statements, fmt.Sprintf("CREATE TABLE \"%s\".\"STOCK\" (ID INTEGER PRIMARY KEY NOT NULL GENERATED BY DEFAULT AS IDENTITY,PRODUCT INTEGER NOT NULL UNIQUE,QTY INTEGER NOT NULL)", schema))
	statements = append(statements, fmt.Sprintf("ALTER TABLE \"%s\".\"STOCK\" ADD CONSTRAINT FK_STOCK_ID FOREIGN KEY (PRODUCT) REFERENCES \"%s\".\"PRODUCT\" (ID) ON DELETE CASCADE", schema, schema))
	statements = append(statements, fmt.Sprintf("CREATE TABLE \"%s\".\"CUSTOMERS\" (ID INTEGER PRIMARY KEY NOT NULL GENERATED BY DEFAULT AS IDENTITY,SAL NVARCHAR(16),FNAME NVARCHAR(64) NOT NULL,LNAME NVARCHAR(64) NOT NULL,ADDR1 NVARCHAR(128) NOT NULL,CITY NVARCHAR(64) NOT NULL,DISCOUNT_PCT DECIMAL)", schema))
	statements = append(statements, fmt.Sprintf("CREATE TABLE \"%s\".\"SALES\" (ID INTEGER PRIMARY KEY NOT NULL GENERATED BY DEFAULT AS IDENTITY,CUSTOMER INTEGER NOT NULL,DATE DATE)", schema))
	statements = append(statements, fmt.Sprintf("ALTER TABLE \"%s\".\"SALES\" ADD CONSTRAINT FK_CUSTOMER_ID FOREIGN KEY (CUSTOMER) REFERENCES \"%s\".\"CUSTOMERS\" (ID) ON DELETE CASCADE", schema, schema))
	statements = append(statements, fmt.Sprintf("CREATE TABLE \"%s\".\"SALES_PRODUCTS\" (ID INTEGER PRIMARY KEY NOT NULL GENERATED BY DEFAULT AS IDENTITY,SALE_ID INTEGER NOT NULL,PRODUCT INTEGER NOT NULL,QTY INTEGER NOT NULL)", schema))
	statements = append(statements, fmt.Sprintf("ALTER TABLE \"%s\".\"SALES_PRODUCTS\" ADD CONSTRAINT FK_PRODUCT_ID FOREIGN KEY (PRODUCT) REFERENCES \"%s\".\"PRODUCT\" (ID) ON DELETE CASCADE;", schema, schema))
	statements = append(statements, fmt.Sprintf("ALTER TABLE \"%s\".\"SALES_PRODUCTS\" ADD CONSTRAINT FK_SALE_ID FOREIGN KEY (SALE_ID) REFERENCES \"%s\".\"SALES\" (ID) ON DELETE CASCADE", schema, schema))
	statements = append(statements, fmt.Sprintf("CREATE VIEW \"%s\".\"V_PRODUCT\" AS SELECT PD.ID AS ID,PD.NAME AS NAME,TP.TYPE_NAME AS TYPE,PC.COLOUR AS COLOUR,MN.MANUF AS MANUF, PD.BUY_PRICE AS BUY_PRICE, PD.SELL_PRICE AS SELL_PRICE FROM \"%s\".\"PRODUCT\" AS PD LEFT JOIN \"%s\".\"PRODUCT_TYPE\" AS TP ON PD.TYPE = TP.ID LEFT JOIN \"%s\".\"PRODUCT_COLOUR\" AS PC ON PD.COLOUR = PC.ID LEFT JOIN \"%s\".MANUF AS MN ON PD.MANUF = MN.ID", schema, schema, schema, schema, schema))
	statements = append(statements, fmt.Sprintf("CREATE PROCEDURE \"%s\".\"PLACE_ORDER\" (IN cust_id INT, IN prod_id INT, IN ord_date DATE, IN qty INT) AS BEGIN INSERT INTO \"%s\".\"SALES\" (CUSTOMER, DATE) VALUES (cust_id, ord_date); INSERT INTO \"%s\".\"SALES_PRODUCTS\" (SALE_ID, PRODUCT, QTY) VALUES ((SELECT current_identity_value() FROM DUMMY), prod_id, qty); END;", schema, schema, schema))
	statements = append(statements, fmt.Sprintf("CREATE VIEW \"%s\".\"V_ORDERS\" AS SELECT S.ID AS SID, S.DATE AS DATE, VP.NAME AS PRODUCT_NAME, VP.TYPE AS PRODUCT_TYPE, VP.COLOUR AS COLOUR, VP.MANUF AS MANUFACTURER, VP.BUY_PRICE AS BUY_PRICE, VP.SELL_PRICE AS SELL_PRICE, (VP.SELL_PRICE-(VP.SELL_PRICE*(C.DISCOUNT_PCT/100)))-VP.BUY_PRICE AS PROFIT, C.SAL AS SAL, C.FNAME AS FNAME, C.LNAME AS LNAME, C.CITY AS CITY FROM \"%s\".\"SALES\" AS S LEFT JOIN \"%s\".\"SALES_PRODUCTS\" AS SP ON S.ID = SP.SALE_ID LEFT JOIN \"%s\".\"V_PRODUCT\" AS VP ON SP.PRODUCT = VP.ID LEFT JOIN \"%s\".\"CUSTOMERS\" AS C ON S.CUSTOMER = C.ID", schema, schema, schema, schema, schema))

	return statements
}

// GetMasterDataStatements returns, in order, a slice of strings that can be used to populate the schema with the master data
// the argument 'schema' sets the correct schmea name in the statements
func GetMasterDataStatements(schema string) []string {
	statements := []string{}
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT_TYPE\" (TYPE_NAME, DESCRIPTION) VALUES ('EG', 'Electric Guitar')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT_TYPE\" (TYPE_NAME, DESCRIPTION) VALUES ('AG', 'Acoustic Guitar')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT_TYPE\" (TYPE_NAME, DESCRIPTION) VALUES ('EAG', 'Electro-Acoustic Guitar')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT_TYPE\" (TYPE_NAME, DESCRIPTION) VALUES ('EB', 'Electric Bass Guitar')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT_TYPE\" (TYPE_NAME, DESCRIPTION) VALUES ('AB', 'Acoustic Bass Guitar')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT_TYPE\" (TYPE_NAME, DESCRIPTION) VALUES ('EAB', 'Electro-Acoustic Bass Guitar')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT_TYPE\" (TYPE_NAME, DESCRIPTION) VALUES ('FX', 'Effects Pedal')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT_TYPE\" (TYPE_NAME, DESCRIPTION) VALUES ('AC', 'Accessories (strings, plectrums, etc')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT_COLOUR\" (COLOUR) VALUES ('Sunburst')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT_COLOUR\" (COLOUR) VALUES ('White')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT_COLOUR\" (COLOUR) VALUES ('Black')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT_COLOUR\" (COLOUR) VALUES ('Cream')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT_COLOUR\" (COLOUR) VALUES ('Red')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT_COLOUR\" (COLOUR) VALUES ('Mint')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT_COLOUR\" (COLOUR) VALUES ('Yellow')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT_COLOUR\" (COLOUR) VALUES ('Azure')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT_COLOUR\" (COLOUR) VALUES ('Aqua')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT_COLOUR\" (COLOUR) VALUES ('Natural')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT_COLOUR\" (COLOUR) VALUES ('Amber')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"MANUF\" (MANUF) VALUES('Fenber')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"MANUF\" (MANUF) VALUES('Ibanjez')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"MANUF\" (MANUF) VALUES('Gibbon')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"MANUF\" (MANUF) VALUES('Jacksun')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"MANUF\" (MANUF) VALUES('Jamaha')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"MANUF\" (MANUF) VALUES('Buss')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"MANUF\" (MANUF) VALUES('ECX')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"MANUF\" (MANUF) VALUES('BlackCar')", schema))
	statements = append(statements, fmt.Sprintf("INSERT INTO \"%s\".\"MANUF\" (MANUF) VALUES('BernieBall')", schema))

	return statements
}
