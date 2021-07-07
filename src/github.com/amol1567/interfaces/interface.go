/*Interface
- defines the set of opeations/methods while struct defines memory layout of a type
- store any value inside empty interface
- end with __+er & method name wihout er.
- io.Writer, io.Reader, interface{}
*/
package main

import "fmt"

func main() {
	type Empty interface {
	}
	var e Empty
	fmt.Printf("e's value is %v, type: %T\n", e, e)
	var pi *int
	fmt.Printf("pi's value: %v, Type: %T\n", pi, pi)
	e = 7
	fmt.Printf("e's value is %v, type: %T\n", e, e)
	e = "Hello, World!"
	fmt.Printf("e's value is %v, type: %T\n", e, e)

	type Writer interface {
		Write([]byte)(int, error)
	}
	type Closer interface {
		Close() error
	}
	type WriterCloser interface{
		Writer 
		Closer
	}

	type (
		Shape interface {
			Area() float64
		}
		Circle interface {
			Radius() float64
			Shape
		}
		Rect interface {
			width() float64
			Length() float64
			Shape
		}
	)
//method
type Currency float64
func (c Currency) String() string{
	return fmt.Sprintf("$%.2f", float64(c))
}
}
