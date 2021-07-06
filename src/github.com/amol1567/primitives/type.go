/* - Boolean type
   - Numeric type
     -Integers
	 -Floating Point
	 -Complex numbers
   - Text type
*/
package main

import "fmt"

func main() {

	//boolean
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	p := 1 == 1
	q := 1 == 2
	fmt.Printf("%v, %T\n", p, p)
	fmt.Printf("%v, %T\n", q, q)

	/*Interger
	-Signed
	  -int type has varying size, but min 32 bits
	  -8 bit(int8) through 64 bit (int64)
	  int8 -128 - 127
	  int16 -32768 - 32767
	  int32
	  int64
	-Unsigned
	  -8 bit(byte and uint8) through 32 bit(uint32)
	  int8 0 - 255
	  int16 0 - 65535
	*/
	n1 := 42
	fmt.Printf("%v, %T\n", n1, n1)

	var n2 uint16 = 2
	fmt.Printf("%v, %T\n", n2, n2)

	//bitwise
	n3 := 8              // 2^3
	fmt.Println(n3 << 3) // 2^3 * 2^3
	fmt.Println(n3 >> 3) // 2^3 / 2^3

	//complex
	var n4 complex64 = 1 + 2i //complex128
	var n5 complex128 = complex(1, 2)
	fmt.Printf("%v, %T\n", n4, n4)
	fmt.Printf("%v, %T\n", n5, n5)
	fmt.Printf("%v, %T\n", real(n4), real(n4))
	fmt.Printf("%v, %T\n", imag(n4), imag(n4))

	//string UTF8
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	b := []byte(s) //byte slicing
	fmt.Printf("%v, %T\n", b, b)

	s2 := "this is also string"
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	/*rune UTF32
	  -alias for int 32
	  -Special methods normally required to process
	  -strings.Reader # ReadRune
	*/
	// r := 'a'
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)

}
