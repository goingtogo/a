package main

import "fmt"


func main() {

	/*
	* Unsigned Int types
	*/
	var x1 uint8 = 123
	var x1s byte = 123
	var x2 uint16 = 12345
	var x3 uint32 = 123456
	var x4 uint64 = 1234567

	fmt.Println(x1,x1s,x2,x3,x4)

	/*
	* Int types
	*/

	var y1 int8 = 123
	var y2 int16 = -12345
	var y3 int32 = 123456
	var y3s rune = 123456
	var y4 int64 = -1234567

	fmt.Println(y1,y2,y3,y3s,y4)

	/*
	* Unsigned Int pointer
	*/

	var z1 uintptr = 1234567890

	fmt.Println(z1)

	/*
	* Float types
	*/

	var e1 float32 = 123.4567
	var e2 float64 = 12345.6789

	fmt.Println(e1,e2)

	/*
	* Complex Numbers
	*/

	var c1 complex64 = 12.2 + 2i;
	var c2 complex128 = 14.3 + 2i;

	fmt.Println(c1,c2)
}