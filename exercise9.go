package main

import "fmt"

func main() {
	// %t(Boolean): the word true or false
	fmt.Printf("%t\n", true)
	// %d(Integer): base 10
	fmt.Printf("%d\n", 123)
	// %b(Integer): base 2
	fmt.Printf("%b\n", 14)
	// %c(Integer): the character represented by the corresponding Unicode code point
	fmt.Printf("%c\n", 33)
	// %x(Integer): base 16, with lower-case letters for a-f
	fmt.Printf("%x\n", 456)
	// %f(Floating-point and complex constituents): decimal point but no exponent, e.g. 123.456
	fmt.Printf("%f\n", 78.9)
	// %e(Floating-point and complex constituents): scientific notation, e.g. -1.234456e+78
	fmt.Printf("%e\n", 123400000.0)
	// %E(Floating-point and complex constituents): scientific notation, e.g. -1.234456E+78
	fmt.Printf("%E\n", 123400000.0)
	// %s(String and slice of bytes): the uninterpreted bytes of the string or slice
	fmt.Printf("%s\n", "\"string\"")
	// %q(String and slice of bytes): a double-quoted string safely escaped with Go syntax
	fmt.Printf("%q\n", "\"string\"")
	// %x(String and slice of bytes): base 16, lower-case, two characters per byte
	fmt.Printf("%x\n", "hex this")

	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// Sprintf formats according to a format specifier and returns the resulting string.
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
}
