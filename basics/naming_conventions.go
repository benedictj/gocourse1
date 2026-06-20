package basics

import "fmt"

func basic() {
	// PascalCase: Used for types, functions, and constants - also Structs, interfaces, enums and functions should be named using PascalCase.
	// e.g. CalculateArea, UserProfile, HTTPStatusOK
	// Structs, interfaces, enums and functions should be named using PascalCase.
	// For example:
	type EmployeeGoogle struct {
		ID        int
		FirstName string
		LastName  string
	}

	type EmployeeApple struct {
		ID        int
		FirstName string
		LastName  string
	}

	// snake_case: Used for variables and package names
	// e.g. user_name(id), http_status_ok(http_request), first_name
	// Variables, package and file names should be in snake_case. As well as function parameters and local variables should be in snake_case.

	// UPPERCASE: Used for constants(constants are immutable values that are known at compile time and do not change during the execution of the program)
	// e.g. MAX_CONNECTIONS, DEFAULT_TIMEOUT, PI
	// Constants should be in UPPERCASE with underscores separating words.

	// camelCase: Used for private variables and functions
	// e.g. calculateArea, userProfile, httpStatusOK
	// Private variables and functions should be in camelCase.

	// package names should be in lowercase and should not contain underscores or camelCase. For example, "math", "http", "userprofile" are valid package names.

	const MAXENTRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)
}
