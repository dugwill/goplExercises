// Copyright © 2017 Douglas Will
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 189 of the Go Programming Lanugage.

//Exercise	5.17: Write a variadic function ElementsByTagName that,
//given	an HTML node tree and zero or more names, returns all the elements
//that match one of those names. Here are two example  calls

//func ElementsByTagName(doc *html.Node, name ...string) []*html.Node
//images:=ElementsByTagName(doc,"img")
//headings := ElementsByTagName(doc,	"h1",	"h2",	"h3",	"h4")

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html"
)

var depth int

func main() {

	for _, url := range os.Args[1:] {

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		doc, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		//Search the html document
		headings := ElementByTagName(doc, "h1", "h2", "h3", "h4")
		images := ElementByTagName(doc, "img")
		links := ElementByTagName(doc, "a")

		var n *html.Node
		for _, n = range headings {
			fmt.Printf("Headings: %v\n", n.Data)
		}
		for _, n = range images {
			fmt.Printf("Images: %v\n", n.Data)
		}
		for _, n = range links {
			fmt.Printf("Links: %v\n", n.Data)
		}

	}
}

// ElementsByTagName takes a html document and a list of tag names
// and returns a slice of html nodes matching the tag names
func ElementByTagName(doc *html.Node, name ...string) []*html.Node {
	defer trace("ElementByTagName")()
	var results []*html.Node // Results holds the nodes that match the target data

	//Create and load map with target search criteria
	targets := make(map[string]bool)
	for _, v := range name {
		targets[v] = true
	}

	// Create an anonymous function for walking the node tree
	// visitAllNodes takes an html node tree and recursivly iterates
	// through the entire tree.  Along the way it checks each node type
	// against the target map.  If the type exists int the map, the node
	// is appended to the results list
	var visitAllNodes func(n *html.Node)
	visitAllNodes = func(n *html.Node) {
		defer trace("visitAllNodes")()
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if targets[c.Data] {
				results = append(results, c)
			}
			visitAllNodes(c)
		}
	}

	visitAllNodes(doc)

	return results

}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}
