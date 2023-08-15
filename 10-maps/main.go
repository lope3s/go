package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

// A map maps keys to values.

// The key can be of any type for which the equality operator is defined, such as integers, 
// floating point and complex numbers, strings, pointers, interfaces (as long as the dynamic type supports equality), 
// structs and arrays. 
// Slices cannot be used as map keys, because equality is not defined on them. 
// Like slices, maps hold references to an underlying data structure. 
// If you pass a map to a function that changes the contents of the map, the changes will be visible in the caller.

var m map[string]Vertex

// The zero value of a map is nill.
// A nill map has no keys, nor can keys be added.

func main() {
    // The make fn returns a map of the given type, initialized and ready for use.
    m = make(map[string]Vertex)

    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }

    fmt.Println(m["Bell Labs"])

    // Map literals are like struct literals, but the keys are required.
    m2 := map[string]Vertex{
        "Google": {
            37.42202, -122.08408,
        },
    }
    fmt.Println(m2)

    m2["Google"] = Vertex{0, 0}

    fmt.Printf("m2['Google'] updated: %v\n", m2)

    google := m2["Google"]

    fmt.Printf("Retrieving an element: %v\n", google)

    delete(m2, "Google")

    fmt.Printf("deleting Google from m2: %v\n", m2)

    // If the key exists in the map, ok will be true, false otherwise.
    // If the element doesn't exist in the map, ele will be a zero value for the type of the entries in the map.
    // In this case {0 0}
    ele, ok := m2["Google"]

    fmt.Printf("Testing if a key is present with two-value assignment, ele: %v, ok: %t\n", ele, ok)
}
