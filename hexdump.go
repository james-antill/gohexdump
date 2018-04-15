package hexdump

import (
	"fmt"
	"unicode"
)

const charactersPerLine = 16

// DumpLine dump 16 bytes of data into a line of hexdump
func DumpLine(data []byte, addr int) (string, error) {
	if len(data) < charactersPerLine {
		return "", fmt.Errorf("not enough data")
	}

	ret := fmt.Sprintf("0x%08X:", addr)

	// hexdump data...
	for i := 0; i < charactersPerLine; i += 2 {
		ret += fmt.Sprintf(" %02X%02X", data[i+0], data[i+1])
	}
	ret += "  "

	// Printable data...
	for i := 0; i < charactersPerLine; i++ {
		if false {
			if unicode.IsPrint(rune(data[i])) {
				ret += string(data[i])
			} else {
				ret += "."
			}
		} else {
			d := data[i]
			switch {
			case d == ' ':
				fallthrough
			case d == '.':
				fallthrough
			case d == ',':
				fallthrough
			case int(d) >= 0x21 && int(d) <= 0x7E:
				ret += string(data[i])
			default:
				ret += "."
			}
		}
	}

	ret += "\n"

	return ret, nil
}
