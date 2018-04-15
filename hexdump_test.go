package hexdump

import "testing"

func TestLine(t *testing.T) {
	data := []struct {
		off  byte
		addr int
		res  string
	}{
		{00, 00, "0x00000000: " +
			"0001 0203 0405 0607 0809 0A0B 0C0D 0E0F  ................\n"},
		{16, 0x10, "0x00000010: " +
			"1011 1213 1415 1617 1819 1A1B 1C1D 1E1F  ................\n"},
		{32, 0x20, "0x00000020: " +
			"2021 2223 2425 2627 2829 2A2B 2C2D 2E2F   !\"#$%&'()*+,-./\n"},
		{48, 0x30, "0x00000030: " +
			"3031 3233 3435 3637 3839 3A3B 3C3D 3E3F  0123456789:;<=>?\n"},
		{64, 0x40, "0x00000040: " +
			"4041 4243 4445 4647 4849 4A4B 4C4D 4E4F  @ABCDEFGHIJKLMNO\n"},
		{80, 0x50, "0x00000050: " +
			"5051 5253 5455 5657 5859 5A5B 5C5D 5E5F  PQRSTUVWXYZ[\\]^_\n"},
		{96, 0x60, "0x00000060: " +
			"6061 6263 6465 6667 6869 6A6B 6C6D 6E6F  `abcdefghijklmno\n"},
		{112, 0x70, "0x00000070: " +
			"7071 7273 7475 7677 7879 7A7B 7C7D 7E7F  pqrstuvwxyz{|}~.\n"},
		{128, 0x80, "0x00000080: " +
			"8081 8283 8485 8687 8889 8A8B 8C8D 8E8F  ................\n"},
	}
	for i := range data {
		off := data[i].off
		addr := data[i].addr
		res := data[i].res

		val, err := DumpLine([]byte{
			off + 0, off + 1, off + 2, off + 3,
			off + 4, off + 5, off + 6, off + 7,
			off + 8, off + 9, off + 10, off + 11,
			off + 12, off + 13, off + 14, off + 15}, addr)

		if err != nil {
			t.Errorf("err not nil")

		}
		if val != res {
			t.Errorf("data not equl: %v %v\n tst=<%s>\n got <%s>\n",
				off, addr, res, val)
		}
	}
}
