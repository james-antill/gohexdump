package hexdump

import "testing"

func TestDumpLine(t *testing.T) {
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

		data := []byte{}
		for i := 0; i < 16; i++ {
			data = append(data, off+byte(i))
		}

		if val := DumpLine(data, addr); val != res {
			t.Errorf("data not equl: %v %v\n tst=<%s>\n got <%s>\n",
				off, addr, res, val)
		}
	}
}

func TestDump256(t *testing.T) {
	data := []struct {
		off  byte
		addr int
		res  string
	}{
		{00, 00, "0x00000000: " +
			"0001 0203 0405 0607 0809 0A0B 0C0D 0E0F  ................\n" +
			"0x00000010: " +
			"1011 1213 1415 1617 1819 1A1B 1C1D 1E1F  ................\n" +
			"0x00000020: " +
			"2021 2223 2425 2627 2829 2A2B 2C2D 2E2F   !\"#$%&'()*+,-./\n" +
			"0x00000030: " +
			"3031 3233 3435 3637 3839 3A3B 3C3D 3E3F  0123456789:;<=>?\n" +
			"0x00000040: " +
			"4041 4243 4445 4647 4849 4A4B 4C4D 4E4F  @ABCDEFGHIJKLMNO\n" +
			"0x00000050: " +
			"5051 5253 5455 5657 5859 5A5B 5C5D 5E5F  PQRSTUVWXYZ[\\]^_\n" +
			"0x00000060: " +
			"6061 6263 6465 6667 6869 6A6B 6C6D 6E6F  `abcdefghijklmno\n" +
			"0x00000070: " +
			"7071 7273 7475 7677 7879 7A7B 7C7D 7E7F  pqrstuvwxyz{|}~.\n" +
			"0x00000080: " +
			"8081 8283 8485 8687 8889 8A8B 8C8D 8E8F  ................\n" +
			"0x00000090: " +
			"9091 9293 9495 9697 9899 9A9B 9C9D 9E9F  ................\n" +
			"0x000000A0: " +
			"A0A1 A2A3 A4A5 A6A7 A8A9 AAAB ACAD AEAF  ................\n" +
			"0x000000B0: " +
			"B0B1 B2B3 B4B5 B6B7 B8B9 BABB BCBD BEBF  ................\n" +
			"0x000000C0: " +
			"C0C1 C2C3 C4C5 C6C7 C8C9 CACB CCCD CECF  ................\n" +
			"0x000000D0: " +
			"D0D1 D2D3 D4D5 D6D7 D8D9 DADB DCDD DEDF  ................\n" +
			"0x000000E0: " +
			"E0E1 E2E3 E4E5 E6E7 E8E9 EAEB ECED EEEF  ................\n" +
			"0x000000F0: " +
			"F0F1 F2F3 F4F5 F6F7 F8F9 FAFB FCFD FEFF  ................\n"},
	}
	for i := range data {
		off := data[i].off
		addr := data[i].addr
		res := data[i].res

		data := []byte{}
		for i := 0; i < 256; i++ {
			data = append(data, off+byte(i))
		}

		if val := Dump(data, addr); val != res {
			t.Errorf("data not equl: %v %v\n tst=<%s>\n got <%s>\n",
				off, addr, res, val)
		}
	}
}

func TestDump(t *testing.T) {
	data := []struct {
		off  byte
		addr int
		num  int
		res  string
	}{
		{00, 00, 20, "0x00000000: " +
			"0001 0203 0405 0607 0809 0A0B 0C0D 0E0F  ................\n" +
			"0x00000010: " +
			"1011 1213                                ....\n"},
		{0x30, 0, 33, "0x00000000: " +
			"3031 3233 3435 3637 3839 3A3B 3C3D 3E3F  0123456789:;<=>?\n" +
			"0x00000010: " +
			"4041 4243 4445 4647 4849 4A4B 4C4D 4E4F  @ABCDEFGHIJKLMNO\n" +
			"0x00000020: " +
			"50                                       P\n"},
		{0x41, 13, 3, "0x0000000D: " +
			"4142 43                                  ABC\n"},
	}
	for i := range data {
		off := data[i].off
		addr := data[i].addr
		num := data[i].num
		res := data[i].res

		data := []byte{}
		for i := 0; i < num; i++ {
			data = append(data, off+byte(i))
		}

		if val := Dump(data, addr); val != res {
			t.Errorf("data not equl: %v %v %v\n tst=<%s>\n got <%s>\n",
				off, addr, num, res, val)
		}
	}
}
