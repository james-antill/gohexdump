package hexdump

const utf8CR = "␍"
const utf8LF = "␊" // ⏎

const utf8HT = "␉"
const utf8VT = "␋"

const utf8SP = "␠"
const utf8BS = "␈"
const utf8DEL = "␡"

const utf8BEL = "␇" // ⍾ 🛎 ?

const utf8NUL = "␀"
const utf8ESC = "␛"

const missing = "."

// Conf configure hexdump output
type Conf struct {
	utf8CR bool
	utf8LF bool

	utf8HT bool
	utf8VT bool

	utf8SP  bool
	utf8BS  bool
	utf8DEL bool

	utf8BEL bool
	utf8NUL bool
	utf8ESC bool

	hideSP bool
}

// ShowCR display a utf8 character for \r byte.
func (c *Conf) ShowCR() *Conf { c.utf8CR = true; return c }

// HideCR don't display a utf8 character for \r byte.
func (c *Conf) HideCR() *Conf { c.utf8CR = false; return c }

// ShowLF display a utf8 character for \n byte.
func (c *Conf) ShowLF() *Conf { c.utf8LF = true; return c }

// HideLF don't display a utf8 character for \n byte.
func (c *Conf) HideLF() *Conf { c.utf8LF = false; return c }

// ShowHT display a utf8 character for \t byte.
func (c *Conf) ShowHT() *Conf { c.utf8HT = true; return c }

// HideHT don't display a utf8 character for \t byte.
func (c *Conf) HideHT() *Conf { c.utf8HT = false; return c }

// ShowVT display a utf8 character for \v byte.
func (c *Conf) ShowVT() *Conf { c.utf8VT = true; return c }

// HideVT don't display a utf8 character for \v byte.
func (c *Conf) HideVT() *Conf { c.utf8VT = false; return c }

// ShowSP display a utf8 character for " " byte.
func (c *Conf) ShowSP() *Conf { c.utf8SP = true; return c }

// HideSP don't display a utf8 character for " " byte.
func (c *Conf) HideSP() *Conf { c.utf8SP = false; return c }

// ShowPlainSP display a normal character for " " byte.
func (c *Conf) ShowPlainSP() *Conf { c.hideSP = false; return c }

// HidePlainSP don't display a normal character for " " byte.
func (c *Conf) HidePlainSP() *Conf { c.hideSP = true; return c }

// ShowBS display a utf8 character for \b byte.
func (c *Conf) ShowBS() *Conf { c.utf8BS = true; return c }

// HideBS don't display a utf8 character for \b byte.
func (c *Conf) HideBS() *Conf { c.utf8BS = false; return c }

// ShowDEL display a utf8 character for DEL byte.
func (c *Conf) ShowDEL() *Conf { c.utf8DEL = true; return c }

// HideDEL don't display a utf8 character for DEL byte.
func (c *Conf) HideDEL() *Conf { c.utf8DEL = false; return c }

// ShowBEL display a utf8 character for BEL byte.
func (c *Conf) ShowBEL() *Conf { c.utf8BEL = true; return c }

// HideBEL don't display a utf8 character for BEL byte.
func (c *Conf) HideBEL() *Conf { c.utf8BEL = false; return c }

// ShowESC display a utf8 character for ESC byte.
func (c *Conf) ShowESC() *Conf { c.utf8ESC = true; return c }

// HideESC don't display a utf8 character for ESC byte.
func (c *Conf) HideESC() *Conf { c.utf8ESC = false; return c }

// ShowNUL display a utf8 character for \0 byte.
func (c *Conf) ShowNUL() *Conf { c.utf8NUL = true; return c }

// HideNUL don't display a utf8 character for \0 byte.
func (c *Conf) HideNUL() *Conf { c.utf8NUL = false; return c }

// ShowAll display utf8 characters for all specials.
func (c *Conf) ShowAll() *Conf {
	return c.ShowCR().ShowLF().ShowHT().ShowVT().ShowSP().ShowBS().ShowDEL().ShowBEL().ShowESC().ShowESC()
}

// HideAll display utf8 characters for all specials.
func (c *Conf) HideAll() *Conf {
	return c.HideCR().HideLF().HideHT().HideVT().HideSP().HideBS().HideDEL().HideBEL().HideESC().HideESC()
}
