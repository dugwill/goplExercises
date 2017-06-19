// Original Work
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.

// Modifications Copyright © 2017 Douglas Will
// License: https://creativecommons.org/licenses/by-sa/4.0/

// ***********  This Exersize is finished   ********/
// Added handlers for Create, update, and delete.
// I wasn't sure what was meant by 'read' since the handlers
// price and list already read the DB

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create) // New Handler for creating an new item in the DB
	http.HandleFunc("/update", db.update) // New Handler for changing an item price
	http.HandleFunc("/delete", db.delete) // New Handler for deleting an item
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

//Create a type to hold the price
type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

// Create the database to hold the items
type database map[string]dollars

// list read sthe db and prints the items to the writer
func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

// 
func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	cost, err := strconv.ParseFloat(price, 32)
	if err != nil {
		fmt.Fprintf(w, "Invalid Price %v\n", price)
	}

	db[item] = dollars(cost)
	fmt.Fprintf(w, "Item %v, successfully added.", item)

}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	newPrice := req.URL.Query().Get("price")

	if price, ok := db[item]; ok {
		cost, err := strconv.ParseFloat(newPrice, 32)
		if err != nil {
			fmt.Fprintf(w, "Invalid Price %v\n", price)
		}
		db[item] = dollars(cost)
		fmt.Fprintf(w, "Item %v, successfully updated. New Price %v\n", item, db[item])

	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}

}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	_, ok := db[item]

	if ok {
		delete(db, item)
		fmt.Fprintf(w, "Successfully removed %q\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
