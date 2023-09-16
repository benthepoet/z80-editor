package main

import (
	"fmt"
	"os"
)

const missing_file = "Filename was not supplied"

func main() {
	if len(os.Args) < 1 {
		panic(missing_file)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	for {
		b := make([]byte, 1)

		if _, err := f.Read(b); err != nil {
			break
		}

		l := getInstrLen(b[0])
		if l > 1 {
			k := make([]byte, l-1)

			if count, _ := f.Read(k); count < int(l-1) && count > 0 {
				k = k[:count]
			}

			b = append(b, k...)
		}

		fmt.Printf("[% x]\n", b)
	}
}

func getInstrLen(op byte) byte {
	if op < 0x40 {
		// Upper table
		lnb := op & 0x0f

		if op == 0x00 {
			return 1
		} else if lnb == 0x00 || lnb == 0x06 || lnb == 0x0e {
			return 2
		} else if lnb == 0x01 || op == 0x22 || op == 0x32 || op == 0x2a || op == 0x3a {
			return 3
		}
		return 1
	} else if op < 0xc0 {
		// Middle table
		return 1
	} else if op == 0xcb {
		return 0
	} else if op == 0xdd {
		return 0
	} else if op == 0xed {
		return 0
	} else if op == 0xfd {
		return 0
	}
	// Bottom table
	lnb := op & 0x0f

	if lnb == 0x02 || lnb == 0x04 || lnb == 0x0a || lnb == 0x0c || op == 0xc3 || op == 0xcd {
		return 3
	} else if lnb == 0x06 || lnb == 0x0e || op == 0xd3 || op == 0xdb {
		return 2
	}
	return 1
}
