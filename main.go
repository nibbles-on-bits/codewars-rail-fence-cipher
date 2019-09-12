package main

import (
	"fmt"
)

func main() {
	fmt.Println("codewars-rail-fence-cipher please run unit tests!")
}

// Decrypt will decode a string using rail fence cipher
func Decrypt(s string, rails int) string {
	ret := ""
	railArr := make([][]int, rails)

	railptr := 0
	railDir := 1 // 1 for forward, -1 for backward

	for x := 0; x < len(s); x++ {
		railArr[railptr] = append(railArr[railptr], x)

		if railptr == 0 && railDir == -1 {
			railDir = 1
		}

		if railptr+1 == rails {
			railDir = -1
		}

		railptr += railDir
	}

	tmp := make([]byte, len(s))

	ptr := 0
	for x := 0; x < rails; x++ {
		for y := 0; y < len(railArr[x]); y++ {
			tmp[railArr[x][y]] = s[ptr]
			ptr++

		}
	}

	ret = string(tmp)
	return ret

}

// Encrypt will encode a string using rail fence cipher
func Encrypt(s string, rails int) string {

	ret := ""

	slen := len(s)
	railArr := make([][]byte, rails)

	railptr := 0
	railDir := 1 // 1 for forward, -1 for backward

	for x := 0; x < slen; x++ {
		railArr[railptr] = append(railArr[railptr], s[x])

		if railptr == 0 && railDir == -1 {
			railDir = 1
		}

		if railptr+1 == rails {
			railDir = -1
		}

		railptr += railDir
	}

	for x := 0; x < rails; x++ {
		for y := 0; y < len(railArr[x]); y++ {
			ret += string(railArr[x][y])
		}
	}

	return ret
}
