package main

import (
	"fmt"
)

func main() {
	fmt.Println("codewars-rail-fence-cipher")

	BuildRailIndexes(21, 5)
	fmt.Println(Encrypt("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 5))
	fmt.Println(Decrypt("AIQYBHJPRXZCGKOSWDFLNTVEMU", 5))
	fmt.Println(Encrypt("WEAREDISCOVEREDFLEEATONCE", 3))
}

// BuildRailIndexes returns a 2 dimensional slice of integers representing the encoding rails
func BuildRailIndexes(slen int, rails int) {
	//rails := [][]int{}
	//rails := make([][]int, rails)

	railptr := 0
	railDir := 1 // 1 for forward, -1 for backward

	for x := 0; x < slen; x++ {
		fmt.Printf("x=%d  railptr=%d\n", x, railptr)

		if railptr == 0 && railDir == -1 {
			railDir = 1
		}

		if railptr+1 == rails {
			railDir = -1
		}

		railptr += railDir
	}
}

func Decrypt(s string, rails int) string {
	ret := ""

	slen := len(s)
	railArr := make([][]int, rails)

	railptr := 0
	railDir := 1 // 1 for forward, -1 for backward

	for x := 0; x < slen; x++ {
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

// BuildRailIndexes returns a 2 dimensional slice of integers representing the encoding rails
func Encrypt(s string, rails int) string {

	ret := ""

	slen := len(s)
	//rails := [][]int{}
	railArr := make([][]byte, rails)

	railptr := 0
	railDir := 1 // 1 for forward, -1 for backward

	for x := 0; x < slen; x++ {
		fmt.Printf("x=%d  railptr=%d\n", x, railptr)
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
		fmt.Println("rail #%d = %v\n", x, railArr[x])
		for y := 0; y < len(railArr[x]); y++ {
			ret += string(railArr[x][y])
		}
	}

	return ret

}

func printRail(s string) {
	ret := [3][]int{}

	for x := 0; x < len(s); x++ {
		if (x % 4) == 0 { // top row
			ret[0] = append(ret[0], x)
		}

		if (x % 2) != 0 { // middle row
			ret[1] = append(ret[1], x)
		}

		if ((x - 2) % 4) == 0 { // bottom row
			ret[2] = append(ret[2], x)
		}
	}

	for x := 0; x < len(ret[0]); x++ {
		fmt.Printf("%c   ", s[ret[0][x]])
	}

	fmt.Println()
	for x := 0; x < len(ret[1]); x++ {
		fmt.Printf(" %c", s[ret[1][x]])

	}

	fmt.Println()
	fmt.Printf("  ")
	for x := 0; x < len(ret[2]); x++ {
		fmt.Printf("%c   ", s[ret[2][x]])
	}

	//fmt.Printf("   Top Row : %v\n", ret[0])
	//fmt.Printf("Middle Row : %v\n", ret[1])
	//fmt.Printf("Bottom Row : %v\n", ret[2])

}

func getRails(s string) [3][]int {

	ret := [3][]int{}

	for x := 0; x < len(s); x++ {
		if (x % 4) == 0 { // top row
			ret[0] = append(ret[0], x)
		}

		if (x % 2) != 0 { // middle row
			ret[1] = append(ret[1], x)
		}

		if ((x - 2) % 4) == 0 { // bottom row
			ret[2] = append(ret[2], x)
		}
	}

	fmt.Printf("   Top Row : %v\n", ret[0])
	fmt.Printf("Middle Row : %v\n", ret[1])
	fmt.Printf("Bottom Row : %v\n", ret[2])

	return ret

}
