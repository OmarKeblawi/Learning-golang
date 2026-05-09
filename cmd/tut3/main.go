package main

import (
	"fmt"
)

// gas engine
type gasEngine struct {
	mpg     uint8
	gallons uint8
	//owner (just the type) or it can be ownerInfo owner
}

func (e gasEngine) milesLeft() uint32 { // method for gasEngine struct
	return uint32(e.mpg) * uint32(e.gallons)
}

// electric engine
type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e electricEngine) milesLeft() uint32 {
	return uint32(e.mpkwh) * uint32(e.kwh)
}

func canMakeIt(e engine, miles uint32) bool {
	return e.milesLeft() >= miles
}

// inheritanceee!!

type engine interface {
	milesLeft() uint32
}

func main() {
	var gaseng gasEngine = gasEngine{30, 10} // or it can be: eng := gasEngine{30, 10}

	fmt.Println(gaseng.mpg, gaseng.gallons)
	fmt.Printf("miles left: %d\n", gaseng.milesLeft())
	/*its also possible to initialize non-typed structs:
	var eng2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}

	fmt.Println(eng2.mpg, eng2.gallons)
	*/

	// methods for structs:
	var eleceng electricEngine = electricEngine{3, 100}
	fmt.Println(canMakeIt(eleceng, 301))

}
