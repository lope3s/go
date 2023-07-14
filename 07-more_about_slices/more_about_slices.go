package main

import (
	"fmt"
	"strings"
)

func main() {
    fmt.Printf("Slices\n\n")

	array := [6]int{1, 2, 3, 4, 5, 6}
    fmt.Printf("array: %v\n", array)

    slice := array[3:6]
    fmt.Printf("slice := array[3:6], slice: %v\n\n", slice)

    fmt.Println("A slice has both a length and a capacity.")
    fmt.Println("The capacity of the slice is the number of elements in the underlying array, counting from the first element in the slice.")
    fmt.Printf("The length of the slice is equal to the number of elements whithin the slice.\n\n")

    fmt.Println("So:")
	fmt.Printf("slice capacity: %d\n", cap(slice)) // 3
	fmt.Printf("slice lenght: %d\n\n", len(slice))   // 3

    fmt.Println("But if we change slice to array[0:3]:")
    slice = array[0:3]
	fmt.Printf("slice capacity: %d\n", cap(slice)) // 6
	fmt.Printf("slice lenght: %d\n\n", len(slice))   // 3

    fmt.Printf("Whait, what?\n\n")

	fmt.Printf("Slices doesn't store any data, it just describes a section of an underlying array.\n\n")

	rootSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("rootSlice: %v\n", rootSlice) // [1 2 3 4 5 6 7 8 9 10]

	rootSlice = rootSlice[:0]
    fmt.Printf("rootSlice = rootSlice[:0], rootSlice: %v\n", rootSlice) // []

	rootSlice = rootSlice[:4]
    fmt.Printf("rootSlice = rootSlice[:4], rootSlice: %v\n\n", rootSlice) // [1 2 3 4] 

    fmt.Printf("This is only possible because we're not changing the undelying array itself.\n\n")

    fmt.Printf("%s\n\n", strings.Repeat("-", 80))

	fmt.Printf("Changing the elements of a slice modifies the corresponding elements of its underlying array.\n\n")

    fmt.Printf("array: %v\n", array) // [1 2 3 4 5 6]

    fmt.Printf("slice: %v\n\n", slice) // [1 2 3]

	siblingSlice := array[0:3]
	fmt.Printf("siblingSlice (this slice has a reference to the same array as slice): %v\n\n", siblingSlice) // [1 2 3]

	slice[0] = 0
    // this indexing reffer to the slice 0 index, in this case it corresponds to the index 3 from the array.
    fmt.Printf("slice[0] = 0\n\n")

    fmt.Printf("array: %v\n\n", array) // [0 2 3 4 5 6]
	fmt.Printf("Other slices that share the same underlying array will see those changes.\n\n")
	fmt.Printf("siblingSlice: %v\n\n", siblingSlice) // [0 2 3]

    fmt.Printf("%s\n\n", strings.Repeat("-", 80))

	sliceLiteral := []int{1, 2, 3}
	fmt.Printf("sliceLiteral := []int{1, 2, 3}\n\n")
	fmt.Printf("sliceLiteral: %v\n", sliceLiteral) // [1 2 3]
	fmt.Printf("sliceLiteral capacity: %d\n", cap(sliceLiteral))
	fmt.Printf("sliceLiteral length: %d\n", len(sliceLiteral))

    fmt.Printf("%s\n\n", strings.Repeat("-", 80))

    fmt.Printf("A nil slice has a length and capacity of 0 and has no underlying array.\n\n")
	var emptySlice []int
    fmt.Printf("var emptySlice []int\n\n")
	fmt.Printf("emptySlice == nil: %t\n\n", emptySlice == nil) // nil

    fmt.Printf("With that trying to slice or add a value to this emptySlice will panic with a runtime error:\n\n")
	fmt.Printf("emptySlice[0] = 0 -> index out of range.\n")
    fmt.Printf("fmt.Println(emptySlice[0]) -> bound out of range.\n\n")

    fmt.Printf("%s\n\n", strings.Repeat("-", 80))

    fmt.Printf("There are 3 ways to create a slice in go, slicing an existing array, using a slice literal and with the make fn.\n\n")

	fmt.Printf("The make function allocates a zeroed array and returns a slice that refers to that array.\n")
    makeSlice := make([]int, 0, 10)
    fmt.Printf("makeSlice := make([]int, 0, 10) == %v\n\n", makeSlice) // []
    fmt.Printf("where 0 == len(slice) and 10 == cap(slice)\n\n")

    fmt.Printf("%s\n\n", strings.Repeat("-", 80))

	fmt.Printf("To append new elements to a slice we use the append fn.\n\n")
    makeSlice = append(makeSlice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
    fmt.Printf("makeSlice = append(makeSlice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20): %v\n\n", makeSlice)
	fmt.Printf("append return a new slice containing all the elements of the original slice plus the provided values.\n")
	fmt.Printf("If the backing array of s is too small, a bigger array will be allocated.\n\n")

    fmt.Printf("makeSlice capacity: %d\n", cap(makeSlice))
    fmt.Printf("makeSlice length: %d\n", len(makeSlice))
}
