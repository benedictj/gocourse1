package basics

import "fmt"

var middleName = "Cane" // This is a package-level variable
// ':=' cannot be used outside of a function, so we use '=' instead.

func main() {
	// var age int
	// var name string = "John"
	// var name1 = "Jane"

	// count := 10
	// lastName := "Smith"
	// middleName := "Cane" // This is a local variable
	middleName := "Mayor"   // This is a local variable that shadows the package-level variable
	fmt.Println(middleName) // Output: ""

	// Default values
	// Numeric types default to 0,
	//  strings default to "",
	//  and booleans default to false.
	// Pointers, slices, maps, channels, structs and function types default to nil.

	// ----- SCOPE -----
	// Variables declared inside a function are local to that function.
	// Variables declared outside of any function are global and can be accessed from any function in the same package.

	// fmt.Println(firstName) // Output: John

}

func printname() {
	firstname := "Michael"
	fmt.Println(firstname) // Output: Michael
}
