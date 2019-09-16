package main

import "fmt"

func main() {
		mapNew := make(map[string]int)
		mapNew["key1"] = 10
		mapNew["key2"] = 20
		mapNew["key3"] = 30

		fmt.Println(mapNew["key1"])
		fmt.Println(mapNew)

		delete(mapNew, "key3")
		fmt.Println(mapNew)

		// Try to find 'key3', return true of false in ok
		if key, ok := mapNew["key3"]; ok {
				fmt.Println(key, ok)
		} else {
				fmt.Println("Key not found")
		}

		// Syntax sugar by create map
		elements := map[string]string{
				"H": "Hydrogen",
				"He": "Helium",
		}
		fmt.Println(elements)

		// Multi-maps
		newElements := map[string]map[string]string {
				"H": map[string]string{
				"name":"Hydrogen",
				"state":"gas",
				},
				"He": map[string]string{
					"name":"Helium",
					"state":"gas",
				},
		}

		if el, ok := newElements["He"]; ok {
        fmt.Println(el["name"], el["state"])
    }
}