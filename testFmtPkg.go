package main

import (
	"fmt"
)

func main() {

	// Declaring some const variables
	const name, dept = "GeeksforGeeks", "CS"

	// Calling Sprintf() function
	s := fmt.Sprintf("%s is a %s Portal.\n", name, dept)

	// Calling WriteString() function to write the
	// contents of the string "s" to "os.Stdout"
	//io.WriteString(os.Stdout, s)
	fmt.Print(s)

}
