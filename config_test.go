package main

import (
	"testing"
)

func Test_configuration_getConfig(t *testing.T) {
	type args struct {
		cfile string
	}
	tests := []struct {
		name    string
		c       configuration
		args    args
		wantErr bool
	}{
		{"No File", configuration{}, args{"/zzz/zzz/zzz/noFile.json"}, true},
		{"Good File 01", configuration{}, args{"test_data/gdCnf001.json"}, false},
		{"Good File 02", configuration{}, args{"test_data/gdCnf002.json"}, false},
		{"Malformed JSON 01", configuration{}, args{"test_data/malCnf001.json"}, true},
		{"Malformed JSON 02", configuration{}, args{"test_data/malCnf002.json"}, true},
		{"Extra Field 01", configuration{}, args{"test_data/exfCnf001.json"}, false},
		{"Missing Field 01", configuration{}, args{"test_data/mssngFld.json"}, true},
		{"Bad Configuration 01", configuration{}, args{"test_data/badCnf01.json"}, true},
		{"Bad Configuration 02", configuration{}, args{"test_data/badCnf02.json"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.getConfig(tt.args.cfile); (err != nil) != tt.wantErr {
				t.Errorf("configuration.getConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_configuration_verifyConfig(t *testing.T) {
	tests := []struct {
		name    string
		c       configuration
		wantErr bool
	}{
		{"Good configuration 01", configuration{Hostname: "localhost", Port: "30015", Username: "Sausage", Password: "QuietBadPassword", Schema: "GS", DropSchema: true, Workers: 10, Customers: 1000, Orders: 1000000, TrnxRecords: 10000, StartYear: 2010, EndYear: 2015}, false},
		{"Good configuration 02", configuration{Hostname: "192.168.0.1", Port: "30040", Username: "User", Password: "Still_a_BadPassword", Schema: "GtrShop", DropSchema: false, Workers: 64, Customers: 25000, Orders: 800000000, TrnxRecords: 10000, StartYear: 2020, EndYear: 2020}, false},
		{"Unitialised configuration", configuration{}, true},
		{"Missing Hostname", configuration{Port: "30040", Username: "User", Password: "Still_a_BadPassword", Schema: "GtrShop", DropSchema: false, Workers: 64, Customers: 25000, Orders: 800000000, StartYear: 2020, EndYear: 2020}, true},
		{"Missing Port", configuration{Hostname: "theserver.foxy.com", Username: "User", Password: "Still_a_BadPassword", Schema: "GtrShop", DropSchema: false, Workers: 64, Customers: 25000, Orders: 800000000, StartYear: 2020, EndYear: 2020}, true},
		{"Missing User", configuration{Hostname: "theserver.foxy.com", Port: "30040", Password: "Still_a_BadPassword", Schema: "GtrShop", DropSchema: false, Workers: 64, Customers: 25000, Orders: 800000000, StartYear: 2020, EndYear: 2020}, true},
		{"Missing Password", configuration{Hostname: "theserver.foxy.com", Port: "30040", Username: "User", Schema: "GtrShop", DropSchema: false, Workers: 64, Customers: 25000, Orders: 800000000, StartYear: 2020, EndYear: 2020}, true},
		{"Missing Schema", configuration{Hostname: "theserver.foxy.com", Port: "30040", Username: "User", Password: "Still_a_BadPassword", DropSchema: false, Workers: 64, Customers: 25000, Orders: 800000000, StartYear: 2020, EndYear: 2020}, true},
		{"Missing Workers", configuration{Hostname: "theserver.foxy.com", Port: "30040", Username: "User", Password: "Still_a_BadPassword", Schema: "GtrShop", DropSchema: false, Customers: 25000, Orders: 800000000, StartYear: 2020, EndYear: 2020}, true},
		{"Missing Customers", configuration{Hostname: "theserver.foxy.com", Port: "30040", Username: "User", Password: "Still_a_BadPassword", Schema: "GtrShop", DropSchema: false, Workers: 64, Orders: 800000000, StartYear: 2020, EndYear: 2020}, true},
		{"Missing Orders", configuration{Hostname: "theserver.foxy.com", Port: "30040", Username: "User", Password: "Still_a_BadPassword", Schema: "GtrShop", DropSchema: false, Workers: 64, Customers: 25000, StartYear: 2020, EndYear: 2020}, true},
		{"Missing StartYear", configuration{Hostname: "theserver.foxy.com", Port: "30040", Username: "User", Password: "Still_a_BadPassword", Schema: "GtrShop", DropSchema: false, Workers: 64, Customers: 25000, Orders: 800000000, EndYear: 2020}, true},
		{"Start and End years incorrectly set", configuration{Hostname: "theserver.foxy.com", Port: "30040", Username: "User", Password: "Still_a_BadPassword", Schema: "GtrShop", DropSchema: false, Workers: 64, Customers: 25000, Orders: 800000000, StartYear: 2021, EndYear: 1999}, true},
		{"System user configuration 01", configuration{Hostname: "localhost", Port: "30015", Username: "system", Password: "QuietBadPassword", Schema: "GS", DropSchema: true, Workers: 10, Customers: 1000, Orders: 1000000, StartYear: 2010, EndYear: 2015}, true},
		{"System user configuration 02", configuration{Hostname: "gshdb001", Port: "30040", Username: "SYSTEM", Password: "%gre456yh(((76!", Schema: "PLAY", DropSchema: true, Workers: 128, Customers: 99999, Orders: 4500000, StartYear: 2010, EndYear: 2025}, true},
		{"Start year too low", configuration{Hostname: "gshdb001", Port: "30040", Username: "Alan", Password: "%gre456yh(((76!", Schema: "PLAY", DropSchema: true, Workers: 128, Customers: 99999, Orders: 4500000, StartYear: 900, EndYear: 2025}, true},
		{"End year too high", configuration{Hostname: "gshdb001", Port: "30040", Username: "Susan", Password: "%gre456yh(((76!", Schema: "PLAY", DropSchema: true, Workers: 128, Customers: 99999, Orders: 4500000, StartYear: 1980, EndYear: 5000}, true},
		{"Start year too low and end year too high", configuration{Hostname: "gshdb001", Port: "30040", Username: "_44_4", Password: "%gre456yh(((76!", Schema: "PLAY", DropSchema: true, Workers: 128, Customers: 99999, Orders: 4500000, StartYear: 10, EndYear: 6667}, true},
		{"TrnxRecords set too low", configuration{Hostname: "localhost", Port: "30015", Username: "Sausage", Password: "QuietBadPassword", Schema: "GS", DropSchema: true, Workers: 10, Customers: 1000, Orders: 1000000, TrnxRecords: 50, StartYear: 2010, EndYear: 2015}, true},
		{"TrnxRecords set too high", configuration{Hostname: "localhost", Port: "30015", Username: "Sausage", Password: "QuietBadPassword", Schema: "GS", DropSchema: true, Workers: 10, Customers: 1000, Orders: 1000000, TrnxRecords: 50000000, StartYear: 2010, EndYear: 2015}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.verifyConfig(); (err != nil) != tt.wantErr {
				t.Errorf("configuration.verifyConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

/*

 */
