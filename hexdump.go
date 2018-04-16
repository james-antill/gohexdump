package hexdump

import (
	"fmt"
)

const charactersPerLine = 16

// DumpLine dump upto 16 bytes of data into a line of hexdump
func DumpLine(data []byte, addr int) string {
	c := Conf{}
	return c.DumpLine(data, addr)
}

// DumpLine dump upto 16 bytes of data into a line of hexdump
func (c *Conf) DumpLine(data []byte, addr int) string {
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
		d := data[i]
		switch {
		case d == '\r':
			if c.utf8CR {
				ret += utf8CR
			} else {
				ret += missing
			}
		case d == '\n':
			if c.utf8LF {
				ret += utf8LF
			} else {
				ret += missing
			}
		case d == '\t':
			if c.utf8HT {
				ret += utf8HT
			} else {
				ret += missing
			}
		case d == '\v':
			if c.utf8VT {
				ret += utf8VT
			} else {
				ret += missing
			}
		case d == ' ':
			if !c.hideSP {
				ret += " "
			} else if c.utf8SP {
				ret += utf8SP
			} else {
				ret += missing
			}

		case d == '\b':
			if c.utf8BS {
				ret += utf8BS
			} else {
				ret += missing
			}
		case int(d) == 0x7F:
			if c.utf8DEL {
				ret += utf8DEL
			} else {
				ret += missing
			}
		case int(d) == 0x07:
			if c.utf8BEL {
				ret += utf8BEL
			} else {
				ret += missing
			}

		case int(d) == 0x00:
			if c.utf8NUL {
				ret += utf8NUL
			} else {
				ret += missing
			}
		case int(d) == 0x1B:
			if c.utf8ESC {
				ret += utf8ESC
			} else {
				ret += missing
			}

		case int(d) >= 0x21 && int(d) <= 0x7E:
			ret += string(data[i])
		default:
			ret += missing
		}
	}
	ret += "\n"

	return ret
}

// Dump all data as lines of hexdump
func Dump(data []byte, addr int) string {
	ret := ""
	for len(data) > 0 {
		ret += DumpLine(data, addr)
		addr += 16
		if len(data) < 16 {
			break
		}
		data = data[16:]
	}

	return ret
}
