package main

import (
	"fmt"
)

type product struct {
	Manufacturer string
	ProductName  string
	ProductType  string
	Colour       string
	BuyPrice     float32
	SellPrice    float32
	Weight       int //used to decide how popular the product is
}

func getProducts() []product {
	products := []product{
		{"Fenber", "Matocaster", "EG", "White", 379.23, 520.00, 10},
		{"Fenber", "Matocaster", "EG", "Black", 379.23, 520.00, 10},
		{"Fenber", "Matocaster", "EG", "Cream", 379.23, 520.00, 10},
		{"Fenber", "Matocaster", "EG", "Red", 379.23, 520.00, 10},
		{"Fenber", "Matocaster", "EG", "Yellow", 379.23, 520.00, 5},
		{"Fenber", "Matocaster", "EG", "Natural", 379.23, 520.00, 2},

		{"Fenber", "Squizzer", "EG", "White", 233.79, 350.00, 10},
		{"Fenber", "Squizzer", "EG", "Black", 233.79, 350.00, 9},
		{"Fenber", "Squizzer", "EG", "Cream", 233.79, 350.00, 8},
		{"Fenber", "Squizzer", "EG", "Red", 233.79, 350.00, 11},
		{"Fenber", "Squizzer", "EG", "Yellow", 233.79, 350.00, 3},
		{"Fenber", "Squizzer", "EG", "Natural", 233.79, 350.00, 3},
		{"Fenber", "Squizzer", "EG", "Azure", 233.79, 350.00, 6},

		{"Fenber", "Jazz", "EB", "White", 421.11, 599.00, 4},
		{"Fenber", "Jazz", "EB", "Black", 421.11, 599.00, 3},
		{"Fenber", "Jazz", "EB", "Red", 421.11, 599.00, 4},
		{"Fenber", "Jazz", "EB", "Yellow", 421.11, 599.00, 5},
		{"Fenber", "Jazz", "EB", "Natural", 421.11, 599.00, 4},
		{"Fenber", "Jazz", "EB", "Azure", 421.11, 599.00, 2},
		{"Fenber", "Jazz", "EB", "Amber", 421.11, 599.00, 1},
		{"Fenber", "Jazz", "EB", "Mint", 421.11, 599.00, 2},

		{"Fenber", "Matocoustic", "EAG", "Natural", 176.22, 239.00, 3},
		{"Fenber", "Matocoustic", "EAG", "Black", 176.22, 239.00, 3},
		{"Fenber", "Matocoustic", "EAG", "White", 176.22, 239.00, 3},

		{"Fenber", "Spagmaster", "EG", "Yellow", 542.28, 750.00, 3},
		{"Fenber", "Spagmaster", "EG", "Red", 542.28, 750.00, 3},

		{"Ibanjez", "RB110", "EG", "White", 112.00, 165.00, 4},
		{"Ibanjez", "RB110", "EG", "Black", 112.00, 165.00, 15},
		{"Ibanjez", "RB110", "EG", "Red", 112.00, 165.00, 7},
		{"Ibanjez", "RB110", "EG", "Mint", 112.00, 165.00, 2},
		{"Ibanjez", "RB110", "EG", "Yellow", 112.00, 165.00, 7},
		{"Ibanjez", "RB110", "EG", "Azure", 112.00, 165.00, 7},

		{"Ibanjez", "RB250", "EG", "White", 295.4, 350.00, 7},
		{"Ibanjez", "RB250", "EG", "Black", 295.4, 350.00, 10},
		{"Ibanjez", "RB250", "EG", "Red", 295.4, 350.00, 7},
		{"Ibanjez", "RB250", "EG", "Mint", 295.4, 350.00, 4},
		{"Ibanjez", "RB250", "EG", "Yellow", 295.4, 350.00, 5},
		{"Ibanjez", "RB250", "EG", "Azure", 295.4, 350.00, 2},

		{"Ibanjez", "B10J", "EB", "White", 145.32, 185.00, 11},
		{"Ibanjez", "B10J", "EB", "Black", 145.32, 185.00, 12},
		{"Ibanjez", "B10J", "EB", "Red", 145.32, 185.00, 4},
		{"Ibanjez", "B10J", "EB", "Mint", 145.32, 185.00, 2},
		{"Ibanjez", "B10J", "EB", "Natural", 145.32, 185.00, 12},

		{"Gibbon", "PLC", "EG", "Red", 1234.55, 1450.00, 5},
		{"Gibbon", "PLC", "EG", "Black", 1234.55, 1450.00, 2},
		{"Gibbon", "PLC", "EG", "White", 1234.55, 1450.00, 1},
		{"Gibbon", "PLC", "EG", "Yellow", 1234.55, 1450.00, 3},

		{"Gibbon", "PLS", "EG", "Red", 1864.51, 2050.00, 4},
		{"Gibbon", "PLS", "EG", "White", 1864.51, 2050.00, 1},
		{"Gibbon", "PLS", "EG", "Black", 1864.51, 2050.00, 1},
		{"Gibbon", "PLS", "EG", "Yellow", 1864.51, 2050.00, 1},

		{"Gibbon", "GS", "EG", "Red", 994.32, 1250.00, 5},
		{"Gibbon", "GS", "EG", "Black", 994.32, 1250.00, 4},
		{"Gibbon", "GS", "EG", "White", 994.32, 1250.00, 3},
		{"Gibbon", "GS", "EG", "Yellow", 994.32, 1250.00, 1},

		{"Gibbon", "EB4", "EB", "Red", 784.22, 999.99, 4},
		{"Gibbon", "EB4", "EB", "Black", 784.22, 999.99, 4},
		{"Gibbon", "EB4", "EB", "White", 784.22, 999.99, 4},
		{"Gibbon", "EB4", "EB", "Yellow", 784.22, 999.99, 2},

		{"Gibbon", "EB5", "EB", "Red", 983.23, 1099.00, 3},
		{"Gibbon", "EB5", "EB", "Black", 983.23, 1099.00, 3},
		{"Gibbon", "EB5", "EB", "White", 983.23, 1099.00, 3},
		{"Gibbon", "EB5", "EB", "Yellow", 983.23, 1099.00, 3},

		{"Jacksun", "X11", "EG", "Red", 94.33, 125.00, 11},
		{"Jacksun", "X11", "EG", "Sunburst", 94.33, 125.00, 11},
		{"Jacksun", "X11", "EG", "White", 94.33, 125.00, 9},
		{"Jacksun", "X11", "EG", "Black", 94.33, 125.00, 13},
		{"Jacksun", "X11", "EG", "Cream", 94.33, 125.00, 11},
		{"Jacksun", "X11", "EG", "Red", 94.33, 125.00, 4},
		{"Jacksun", "X11", "EG", "Aqua", 94.33, 125.00, 2},
		{"Jacksun", "X11", "EG", "Natural", 94.33, 125.00, 2},
		{"Jacksun", "X11", "EG", "Amber", 94.33, 125.00, 1},

		{"Jacksun", "X21", "EG", "Red", 456.32, 550.00, 7},
		{"Jacksun", "X21", "EG", "Sunburst", 456.32, 550.00, 7},
		{"Jacksun", "X21", "EG", "White", 456.32, 550.00, 7},
		{"Jacksun", "X21", "EG", "Black", 456.32, 550.00, 6},
		{"Jacksun", "X21", "EG", "Cream", 456.32, 550.00, 6},
		{"Jacksun", "X21", "EG", "Red", 456.32, 550.00, 5},
		{"Jacksun", "X21", "EG", "Aqua", 456.32, 550.00, 6},
		{"Jacksun", "X21", "EG", "Natural", 456.32, 550.00, 4},
		{"Jacksun", "X21", "EG", "Amber", 456.32, 550.00, 2},

		{"Jamaha", "AC6", "AC", "Red", 66.54, 125.00, 11},
		{"Jamaha", "AC6", "AC", "Sunburst", 66.54, 125.00, 12},
		{"Jamaha", "AC6", "AC", "White", 66.54, 125.00, 13},
		{"Jamaha", "AC6", "AC", "Black", 66.54, 125.00, 15},
		{"Jamaha", "AC6", "AC", "Natural", 66.54, 125.00, 9},
		{"Jamaha", "AC6", "AC", "Amber", 66.54, 125.00, 2},

		{"Jamaha", "AC7", "AC", "Red", 85.40, 199.00, 2},
		{"Jamaha", "AC7", "AC", "Sunburst", 85.40, 199.00, 2},
		{"Jamaha", "AC7", "AC", "White", 85.40, 199.00, 1},
		{"Jamaha", "AC7", "AC", "Black", 85.40, 199.00, 2},
		{"Jamaha", "AC7", "AC", "Natural", 85.40, 199.00, 2},
		{"Jamaha", "AC7", "AC", "Amber", 85.40, 199.00, 2},

		{"Jamaha", "AB4", "AB", "Red", 112.11, 199.00, 5},
		{"Jamaha", "AB4", "AB", "Sunburst", 112.11, 199.00, 4},
		{"Jamaha", "AB4", "AB", "White", 112.11, 199.00, 5},
		{"Jamaha", "AB4", "AB", "Black", 112.11, 199.00, 5},
		{"Jamaha", "AB4", "AB", "Natural", 112.11, 199.00, 5},
		{"Jamaha", "AB4", "AB", "Amber", 112.11, 199.00, 4},

		{"Jamaha", "AB5", "AB", "Red", 155.44, 219.00, 1},
		{"Jamaha", "AB5", "AB", "Sunburst", 155.44, 219.00, 1},
		{"Jamaha", "AB5", "AB", "White", 155.44, 219.00, 1},
		{"Jamaha", "AB5", "AB", "Black", 155.44, 219.00, 2},
		{"Jamaha", "AB5", "AB", "Natural", 155.44, 219.00, 1},
		{"Jamaha", "AB5", "AB", "Amber", 155.44, 219.00, 1},

		{"Jamaha", "AB4e", "EAB", "Red", 224.32, 299.99, 1},
		{"Jamaha", "AB4e", "EAB", "Sunburst", 224.32, 299.99, 1},
		{"Jamaha", "AB4e", "EAB", "White", 224.32, 299.99, 1},
		{"Jamaha", "AB4e", "EAB", "Black", 224.32, 299.99, 1},
		{"Jamaha", "AB4e", "EAB", "Natural", 224.32, 299.99, 1},
		{"Jamaha", "AB4e", "EAB", "Amber", 224.32, 299.99, 1},
		{"Jamaha", "AB4e", "EAB", "Red", 224.32, 299.99, 1},

		{"Buss", "Distortion", "FX", "NULL", 89.00, 110.00, 5},
		{"Buss", "DoubleDistortion", "FX", "NULL", 123.29, 140.00, 5},
		{"Buss", "Overdrive", "FX", "NULL", 144.74, 202.00, 4},
		{"Buss", "OverOverdrive", "FX", "NULL", 150.00, 160.00, 2},

		{"ECX", "BigBuff", "FX", "NULL", 55.43, 75.00, 7},
		{"ECX", "BigBuffNano", "FX", "NULL", 40.23, 56.00, 6},
		{"ECX", "MetalMan", "FX", "NULL", 92.43, 129.99, 1},
		{"ECX", "BifBuffBass", "FX", "NULL", 82.34, 105.00, 2},

		{"BlackCar", "AC100", "FX", "NULL", 44.22, 66.77, 4},
		{"BlackCar", "OG40", "FX", "NULL", 96.66, 140.00, 3},
		{"BlackCar", "OG68", "FX", "NULL", 223.22, 270.00, 1},

		{"BernieBall", "10S", "AC", "NULL", 6.45, 9.99, 30},
		{"BernieBall", "11S", "AC", "NULL", 6.49, 10.50, 25},
		{"BernieBall", "9S", "AC", "NULL", 6.49, 10.50, 10},
		{"BernieBall", "10H", "AC", "NULL", 7.45, 10.99, 20},
		{"BernieBall", "11H", "AC", "NULL", 7.49, 11.50, 15},
		{"BernieBall", "9H", "AC", "NULL", 7.49, 11.50, 5},
		{"BernieBall", "OG68", "AC", "NULL", 7.49, 11.50, 13},
	}
	return products
}

//GetProductStatements returns a list, in order, of SQL statements to be used to populate the product data.
func GetProductStatements(schema string) []string {
	products := getProducts()
	statements := make([]string, len(products))

	for i, product := range products {
		if product.Colour == "NULL" {
			statements[i] = fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT\"(NAME, TYPE, MANUF, BUY_PRICE, SELL_PRICE, RAND_WEIGHT) VALUES ('%s', (SELECT ID FROM \"%s\".\"PRODUCT_TYPE\" WHERE TYPE_NAME = '%s'), (SELECT ID FROM \"%s\".\"MANUF\" WHERE MANUF = '%s'), %.2f, %.2f, %d)", schema, product.ProductName, schema, product.ProductType, schema, product.Manufacturer, product.BuyPrice, product.SellPrice, product.Weight)
		} else {
			statements[i] = fmt.Sprintf("INSERT INTO \"%s\".\"PRODUCT\" (NAME, TYPE, MANUF, COLOUR, BUY_PRICE, SELL_PRICE, RAND_WEIGHT) VALUES ('%s', (SELECT ID FROM \"%s\".\"PRODUCT_TYPE\" WHERE TYPE_NAME = '%s'), (SELECT ID FROM \"%s\".\"MANUF\" WHERE MANUF = '%s'), (SELECT ID FROM \"%s\".\"PRODUCT_COLOUR\" WHERE COLOUR = '%s'), %.2f, %.2f, %d)", schema, product.ProductName, schema, product.ProductType, schema, product.Manufacturer, schema, product.Colour, product.BuyPrice, product.SellPrice, product.Weight)
		}

	}
	return statements
}

//func getWeightedProductIds(schema string, db1 *sql.DB) ([]randutil.Choice, error) {
//
//	products := getProducts()
//	productChoice := []randutil.Choice{}
//	var id int
//
//	//Iterate through the products in the DB to get their IDs, then assign their IDs and weight to []Ch1
//	for _, product := range products {
//		q1 := productIDQuery(schema, product)
//		r1 := db1.QueryRow(q1)
//		err := r1.Scan(&id)
//		if err != nil {
//			log.Printf("Scan Failed\n")
//			return productChoice, err
//		}
//		//If we've got this far we can fill in the choice
//		productChoice = append(productChoice, randutil.Choice{Weight: product.Weight, Item: id})
//
//	}
//	return productChoice, nil
//}

//func productIDQuery(schema string, prd product) string {
//	var s1 string
//	if prd.Colour != "NULL" {
//		s1 = fmt.Sprintf("SELECT ID FROM \"%s\".\"PRODUCT\""+
//			"WHERE "+
//			"NAME = '%s' "+
//			"AND TYPE = (SELECT ID FROM \"%s\".\"PRODUCT_TYPE\" WHERE TYPE_NAME = '%s') "+
//			"AND MANUF = (SELECT ID FROM \"%s\".\"MANUF\" WHERE MANUF = '%s') "+
//			"AND COLOUR = (SELECT ID FROM \"%s\".\"PRODUCT_COLOUR\" WHERE COLOUR = '%s');", schema, prd.ProductName, schema, prd.ProductType, schema, prd.Manufacturer, schema, prd.Colour)
//	} else {
//		s1 = fmt.Sprintf("SELECT ID FROM \"%s\".\"PRODUCT\""+
//			"WHERE "+
//			"NAME = '%s' "+
//			"AND TYPE = (SELECT ID FROM \"%s\".\"PRODUCT_TYPE\" WHERE TYPE_NAME = '%s') "+
//			"AND MANUF = (SELECT ID FROM \"%s\".\"MANUF\" WHERE MANUF = '%s')", schema, prd.ProductName, schema, prd.ProductType, schema, prd.Manufacturer)
//	}
//	return s1
//}

//{"Fenber", "Matocaster", "EG", "White", 379.23, 520.00, 10},
/*
"CREATE TABLE \"%s\".\"PRODUCT\"
(
	ID INTEGER PRIMARY KEY NOT NULL GENERATED BY DEFAULT AS IDENTITY,
	NAME NVARCHAR(64) NOT NULL,
	TYPE INTEGER NOT NULL,
	MANUF INTEGER NOT NULL,
	COLOUR INTEGER,
	BUY_PRICE DECIMAL,
	SELL_PRICE DECIMAL
)

type product struct {
	Manufacturer string
	ProductName  string
	ProductType  string
	Colour       string
	BuyPrice     float32
	SellPrice    float32
	Weight       int //used to decide how popular the product is
}
*/
