package string

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

func LowerCase() string { return lowerCase }

func UpperCase() string { return upperCase }

func Letters() string { return letters }

func Digits() string { return digits }

func HexDigits() string { return hexDigits }

func OctDigits() string { return octDigits }

func Punctuation() string { return punctuation }

func WhiteSpace() string { return whiteSpace }

func Printable() string { return printable }
