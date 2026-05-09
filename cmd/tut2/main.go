package main

import (
	"fmt"
	"time"
)

func main() {
	//Arrays
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	//Slices
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("Length of intSlice: %v with capacity %v\n",
		len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("Length of intSlice: %v with capacity %v\n",
		len(intSlice), cap(intSlice))
	var intSlice2 []int32 = []int32{8, 9, 10}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)
	fmt.Printf("Length of intSlice: %v with capacity %v\n",
		len(intSlice), cap(intSlice))
	var intSlice3 []int32 = make([]int32, 5, 20)
	fmt.Println(intSlice3)
	fmt.Printf("Length of intSlice3: %v with capacity %v\n",
		len(intSlice3), cap(intSlice3))

	var Mymap map[string]uint8 = make(map[string]uint8)
	Mymap["key1"] = 1
	Mymap["key2"] = 2
	fmt.Println(Mymap)
	var myMap2 = map[string]uint8{
		"adam": 1,
		"eve":  2,
	}
	fmt.Println(myMap2)
	fmt.Println(myMap2["adam"])
	fmt.Println(myMap2["James"])
	var age, ok = myMap2["James"]
	if ok {
		fmt.Printf("James is %d years old\n", age)
	} else {
		fmt.Println("Invalid name")
	}
	//delete(myMap2, "eve")
	fmt.Println(myMap2)

	for name, age := range myMap2 {
		fmt.Println(name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %d Value: %d\n", i, v)
	}

	fmt.Println("Using while loop in go")

	var i int = 0
	for i < len(intArr) {
		fmt.Printf("Index: %d Value: %d\n", i, intArr[i])
		i++
	}

	fmt.Println("Using c for loop in go")

	for i := 0; i < len(intArr); i++ {
		fmt.Printf("Index: %d Value: %d\n", i, intArr[i])
	}

	//Time the loop
	var slice []int
	var duration time.Duration = timeLoop(slice, 1000000)
	fmt.Printf("Time taken to execute loop: %v\n", duration)
	fmt.Printf("Slice: %v\n", slice)

}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for i := 0; i < n; i++ {
		slice = append(slice, i)
	}
	slice = []int{1}
	fmt.Printf("Slice: %v\n", slice)
	var t1 = time.Now()
	return t1.Sub(t0)
	// return time.Since(t0)
}
