//Excersize 7.3
// The 'func (t *tree) String() string' at the end of this file is
// the result of excersize 7.3.

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
package treesort

import (
	"fmt"
	"strconv"
)

//!+
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	fmt.Printf(root.String())

}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {

	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// 'func (s string, t *tree) String' inplements the fmt.sting interface
// it takes a string and a pointer to a linked list 'tree'
// Walks the Tree, formats and returns a string containing all the link values

func (t *tree) String() string {

	var sequence = "Values:"
	if t == nil {
		return "Empty tree\n"
	}

	//'popSeq' (Populate Sequence) is an anonymous function that, first,
	// finds the left end of the tree, then walks the tree adding each
	// value to a string 's'. Its format was barrowed from the 'appendValues'
	// function above. Instead of adding int to a []int, it adds additional
	// string to a seed string
	var popSeq func(s string, t *tree) string
	popSeq = func(s string, t *tree) string {
		if t != nil {
			s = popSeq(s, t.left)
			s = s + strconv.Itoa(t.value) + " "
			s = popSeq(s, t.right)
		}
		return s
	}

	sequence = popSeq(sequence, t)

	// Format the result string and return
	return fmt.Sprintf("%s", sequence)

}

//!-
