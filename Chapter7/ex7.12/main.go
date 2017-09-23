// Original Work
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.

// Modifications Copyright © 2017 Douglas Will
// License: https://creativecommons.org/licenses/by-sa/4.0/

// ***********  This Exersize is finished   ********/
// Change the list function to print output as HTML Table
// Added a html template, itemList
// Modified list function to write the item list using template

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// ex7.12- This template creates an full web page that will show the
// items list as a web table
var itemList = template.Must(template.New("itemList").Parse(`
	
	<!DOCTYPE html>
	<html>
	<body>
	
	<h1>{{"Price List"}}</h1>
	<table>
	
	{{range $item, $price := .}}
	<tr>
	  <td>{{$item}}</td>
	  <td>{{$price}}</td>
	</tr>
	{{end}}
	</table>
	
	</body>
	</html>
	
	`))

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

// list reads the db and prints the items to the writer
// ex 7.12 - list was modified to use the html template, itemList
func (db database) list(w http.ResponseWriter, req *http.Request) {

	if err := itemList.Execute(w, db); err != nil {
		log.Fatal(err)
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
