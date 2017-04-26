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

	fmt.Printf("%v\n", cache)

	// Since the 'Lock' method fo 'sync.Mutex' is available to 'cache'
	// we can use it to lock and unlock the map.
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
