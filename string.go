package string

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	lowerCase   = "abcdefghijklmnopqrstuvwxyz"
	upperCase   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letters     = lowerCase + upperCase
	digits      = "0123456789"
	hexDigits   = "0123456789abcdefABCDEF"
	octDigits   = "01234567"
	punctuation = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~."
	whiteSpace  = " \t\n\r\x0b\x0c"
	printable   = digits + letters + punctuation + whiteSpace
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
