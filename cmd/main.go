package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		panic("Must supply a file name")
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

func getInstrLen(i byte) byte {
	if i < 0x40 {
		// Upper table
		l := i & 0x0f

		if i == 0x00 {
			return 1
		} else if l == 0x00 || l == 0x06 || l == 0x0e {
			return 2
		} else if l == 0x01 || i == 0x22 || i == 0x32 || i == 0x2a || i == 0x3a {
			return 3
		} else {
			return 1
		}
	} else if i < 0xc0 {
		// Middle table
		return 1
	} else if i == 0xcb {
		return 0
	} else if i == 0xdd {
		return 0
	} else if i == 0xed {
		return 0
	} else if i == 0xfd {
		return 0
	} else {
		// Bottom table
		l := i & 0x0f

		if l == 0x02 || l == 0x04 || l == 0x0a || l == 0x0c || i == 0xc3 || i == 0xcd {
			return 3
		} else if l == 0x06 || l == 0x0e || i == 0xd3 || i == 0xdb {
			return 2
		} else {
			return 1
		}
	}
}
