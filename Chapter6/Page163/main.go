// Original Work gopl.io Example on Page 163
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Modifications Copyright © 2017 Douglas Will
// License: https://creativecommons.org/licenses/by-sa/4.0/

package main

import (
	"fmt"
	"sync"
)

// This var statement groups the struct and the map together in
// a single var 'cache'. The methods of 'sync.mutex' are
// promoted to 'cache' through the unnamed struct type
var cache = struct {
	sync.Mutex // guards mapping
	mapping    map[string]string
}{
	mapping: make(map[string]string),
}

func main() {

<<<<<<< HEAD
	fmt.Printf("%v\n", cache)

	// Since the 'Lock' method fo 'sync.Mutex' is available to 'cache'
	// we can use it to lock and unlock the map.
=======
	// Since the 'Lock' and 'Unlock' methods for 'sync.Mutex' are
	// available to 'cache' we can use it to lock and unlock the map.
>>>>>>> a18b2bdb42a5580b7450bb942ce5eb2115947802
	cache.Lock()
	cache.mapping["bob"] = "jones"
	cache.mapping["joe"] = "smith"
	cache.Unlock()

	fmt.Printf("%v\n", Lookup("joe"))

}

// Since the 'cache' is a package level var, we can use the
// methods in the package funcs as well
func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
