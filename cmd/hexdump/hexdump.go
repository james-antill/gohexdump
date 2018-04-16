package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	hexdump "github.com/james-antill/gohexdump"
)

func hexdumpIO(c *hexdump.Conf, addr int, ior io.Reader, iow io.Writer) (int, error) {
	dataStore := [16]byte{}
	for {
		stop := false
		data := dataStore[:]
		n, err := io.ReadAtLeast(ior, data, 16)
		if n == 0 && err == io.EOF {
			return addr, nil
		}
		if err == io.ErrUnexpectedEOF {
			data = data[:n]
			stop = true
			err = nil
		}
		if err != nil {
			return addr, err
		}
		_, err = io.WriteString(iow, c.DumpLine(data, addr))
		if err != nil {
			return addr, err
		}
		addr += n
		if stop {
			break
		}
	}
	return addr, nil
}

func main() {

	utf8 := flag.Bool("utf8", false, "show utf8 for special characters")
	nsp := flag.Bool("no-space", false, "don't show space character")

	flag.Parse()

	c := &hexdump.Conf{}

	if *utf8 {
		c.ShowAll()
	}
	if *nsp {
		c.HidePlainSP()
	}

	addr := 0
	if len(flag.Args()) < 1 {
		hexdumpIO(c, addr, os.Stdin, os.Stdout)
		return
	}

	for _, arg := range flag.Args() {
		ior, err := os.Open(arg)
		if err != nil {
			fmt.Printf("Open(%s): %v\n", arg, err)
			os.Exit(1)
		}
		addr, err = hexdumpIO(c, addr, ior, os.Stdout)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(2)
		}
	}
}
