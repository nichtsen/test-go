package typeof

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

// TestInterface
// notation: -> means a pointer points to;
// => means a interface binds to;
func TestInterface(t *testing.T) {
	var vwtr *io.Writer      // *interface->interface
	wtr := (*io.Writer)(nil) //*interface->interface
	wtr2 := (io.Writer)(nil)
	p := (io.Writer)(&pencil{}) //interface=>*pencil->interface
	p2 := (*pencil)(nil)        //*pencil->nil
	p3 := &pencil{}             //*pencil->pencil

	typ := reflect.TypeOf(vwtr).Elem()
	fmt.Println(typ.Name()) // Writer

	typ = reflect.TypeOf(wtr).Elem()
	fmt.Println(typ.Name()) // Writer

	_ = reflect.TypeOf(wtr2) //interface=>nil is not a pointer and as a interface do not contain a vaule

	typ = reflect.TypeOf(&wtr2).Elem() //*interface->interface
	fmt.Println(typ.Name())            // Writer

	typ = reflect.TypeOf(p).Elem()
	fmt.Println(typ.Name()) // pencil

	typ = reflect.TypeOf(&p).Elem() //*interface->interface=>*pencil->pencil
	fmt.Println(typ.Name())         // Writer

	typ = reflect.TypeOf(&p2).Elem()
	fmt.Println(typ.Name()) // ""

	typ = reflect.TypeOf(p3).Elem()
	fmt.Println(typ.Name()) //pencil

	typ = reflect.TypeOf(&p3).Elem()
	fmt.Println(typ.Name()) // "" //*->*pencil
}
