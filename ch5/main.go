// Chapter 5. Arrays, Slices, and Maps

package main

import "fmt"

func main() {
	// Declare and print array
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	// Samples
	fmt.Println("arraySample()")
	arraySample()
	fmt.Println("arraySampleFlexCompact()")
	arraySampleFlexCompact()
	fmt.Println("sliceSample()")
	sliceSample()
	fmt.Println("sliceAppendSample()")
	sliceAppendSample()
	fmt.Println("sliceCopySample()")
	sliceCopySample()
	fmt.Println("mapSample1()")
	mapSample1()
	fmt.Println("mapSample2()")
	mapSample2()
	fmt.Println("mapSample3()")
	mapSample3()

	// Exercises
	fmt.Println("ex4()")
	ex4()
}

// arraySample1 : declare float array, calculate average of values
func arraySample() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / 5)
}

// arraySample1FlexCompact : more compact, scales w/ size of array
func arraySampleFlexCompact() {
	// Alternative:
	// x := [5]float64{ 98, 93, 77, 82, 83 }
	x := [5]float64{
		98,
		93,
		77,
		82,
		83, // underscore in last element required for easy removal/commenting
	}
	var total float64         // no need for "= 0"
	for _, value := range x { // _ used as iterator is never used inside loop
		total += value
	}
	fmt.Println(total / float64(len(x)))
}

func sliceSample() {
	// var x []float64				-->	declare slice of length zero
	// x := make([]float64, 5)		--> declare slice of length five
	x := make([]float64, 5, 10) //	--> slice of length 5 with a capacity of 10
	fmt.Println(x)

	arr := [5]float64{1, 2, 3, 4, 5}
	arrSlice := arr[0:4] // Includes index 0, not index 4
	fmt.Println(arrSlice)

	/*
		For convenience, we are also allowed to omit low, high, or even both low
		and high. arr[0:] is the same as arr[0:len(arr)], arr[:5] is the same as
		arr[0:5], and arr[:] is the same as arr[0:len(arr)].
	*/
}

func sliceAppendSample() {
	/*
		append adds elements onto the end of a slice. If there’s sufficient
		capacity in the underlying array, the element is placed after the last
		element and the length is incremented. However, if there is not sufficient
		capacity, a new array is created, all of the existing elements are copied
		over, the new element is added onto the end, and the new slice is returned.
	*/
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice1, slice2)
}

func sliceCopySample() {
	/*
		copy takes two arguments: dst and src. All of the entries in src are
		copied into dst overwriting whatever is there. If the lengths of the two
		slices are not the same, the smaller of the two will be used.
	*/
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1) // copy(dst, src)
	fmt.Println(slice1, slice2)
}

func mapSample1() {
	/*
		A map is an unordered collection of key-value pairs (maps are also
		sometimes called associative arrays, hash tables, or dictionaries). Maps
		are used to look up a value by its associated key.
	*/
	x1 := make(map[string]int) // map string to integer
	x1["key"] = 10             // initialize map
	fmt.Println(x1["key"])

	x2 := make(map[int]int) // map integer to integer
	x2[1] = 10
	fmt.Println(x2[1])

	/*
		This looks very much like an array, but there are a few differences.
		First, the length of a map (found by doing len(x)) can change as we add
		new items to it. When initially created, it has a length of 0;
		after x[1] = 10 it has a length of 1. Second, maps are not sequential.
		We have x[1], and with an array that would imply there must be an x[0],
		but maps don’t have this requirement. We can also delete items from a map
		using the built-in delete function:
	*/
	delete(x1, "key")
	delete(x2, 1)
}

func mapSample2() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	/*
		Shorter way of creating maps:
		elements := map[string]string{
			"H":  "Hydrogen",
			"He": "Helium",
			"Li": "Lithium",
			"Be": "Beryllium",
			"B":  "Boron",
			"C":  "Carbon",
			"N":  "Nitrogen",
			"O":  "Oxygen",
			"F":  "Fluorine",
			"Ne": "Neon",
		}
	*/

	// Print something that exists
	fmt.Println(elements["Li"])
	// Print something that doesn't exist
	fmt.Println(elements["Un"])

	// Will give an empty string
	name, ok := elements["Un"]
	fmt.Println(name, ok)

	/*
		Typical to see in go programs:
		First, we try to get the value from the map. Then, if it’s successful,
		we run the code inside of the block.
	*/
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}

}

func mapSample3() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}

/*
Exercises

1. How do you access the fourth element of an array or slice?
2. What is the length of a slice created using make([]int, 3, 9)?
3. Given the following array, what would x[2:5] give you?
	x := [6]string{"a","b","c","d","e","f"}
4.Write a program that finds the smallest number in this list:
x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
}
*/

// ex1
// A: xyz[3]

// ex2
// A: Length 3, Capacity 9

// ex3
// A: x[2:5] == [c d e]

// ex4 :
func ex4() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	min := x[0]
	for n := 0; n < len(x); n++ {
		if x[n] < min {
			min = x[n]
		}
	}
	fmt.Println(min)
}
