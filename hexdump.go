package hexdump

import (
	"fmt"
	"unicode"
)

const charactersPerLine = 16

// DumpLine dump upto 16 bytes of data into a line of hexdump
func DumpLine(data []byte, addr int) (string, error) {
	ret := fmt.Sprintf("0x%08X:", addr)

	// hexdump data...
	llen := charactersPerLine
	if llen > len(data) {
		llen = 2 * (len(data) / 2)
	}
	i := 0
	for i = 0; i < llen; i += 2 {
		ret += fmt.Sprintf(" %02X%02X", data[i+0], data[i+1])
	}
	if i < charactersPerLine && i < len(data) {
		ret += fmt.Sprintf(" %02X  ", data[i+0])
		i += 2
	}
	for ; i < charactersPerLine; i += 2 {
		ret += "     "
	}
	ret += "  "

	// Printable data...
	llen = charactersPerLine
	if llen > len(data) {
		llen = len(data)
	}
	for i = 0; i < llen; i++ {
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

// Dump all data as lines of hexdump
func Dump(data []byte, addr int) (string, error) {
	ret := ""
	for len(data) > 0 {
		line, err := DumpLine(data, addr)
		addr += 16
		ret += line
		if err != nil {
			return ret, err
		}
		if len(data) < 16 {
			break
		}
		data = data[16:]
	}

	return ret, nil
}
