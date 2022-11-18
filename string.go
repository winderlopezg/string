// Package text supports text based on characters ASCII.
package string

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Character strings
const (
	// a string containing all ASCII lowercase letters
	lowerCase = "abcdefghijklmnopqrstuvwxyz"
	// a string containing all ASCII uppercase letters
	upperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// // a string containing all ASCII letters
	letters = lowerCase + upperCase
	// a string containing all ASCII digits
	digits = "0123456789"
	// a string containing hexdigits
	hexDigits = "0123456789abcdefABCDEF"
	// a string containing "01234567"
	octDigits = "01234567"
	// a string containing all puntuation characters
	punctuation = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~."
	// a string containing the characters space, tab, linefeed,retrun, and vertical tab
	whiteSpace = " \t\n\r\x0b\x0c"
	// a string containing all ASCII characters considered printable
	printable = digits + letters + punctuation + whiteSpace
)

// a string containing all ASCII lowercase letters
func LowerCase() string { return lowerCase }

// a string containing all ASCII uppercase letters
func UpperCase() string { return upperCase }

// a string containing all ASCII letters
func Letters() string { return letters }

// a string containing all ASCII decimal digits
func Digits() string { return digits }

// a string containing all ASCII hexadecimal digits
func HexDigits() string { return hexDigits }

// a string containing all ASCII octal digits
func OctDigits() string { return octDigits }

// a string containing all ASCII punctuation characters
func Punctuation() string { return punctuation }

// a string containing all ASCII whitespace
func WhiteSpace() string { return whiteSpace }

// a string containing all ASCII characters considered printable
func Printable() string { return printable }

// Capitalize each word
func Capitalize(str string) string {

	return cases.Title(language.Und, cases.NoLower).String(str)
}
