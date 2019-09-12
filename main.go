package main

import (
	"fmt"
)

func main() {
	fmt.Println("codewars-rail-fence-cipher")

	BuildRailIndexes(21, 5)
	fmt.Println(EncryptRail2("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 5))

	/*s := "WEAREDISCOVEREDFLEEATONCE"
	fmt.Printf("s=%v len(s)=%d\n\n", s, len(s))

	toprowCnt := ((len(s) - 1) / 4) + 1
	fmt.Printf("toprowCnt=%v\n", toprowCnt)

	middlerowCnt := (len(s)) / 2
	fmt.Printf("middlerowCnt=%v\n", middlerowCnt)

	bottomrowCnt := (len(s) + 1) / 4

	fmt.Printf("toprow_cnt    = %d\n", toprowCnt)
	fmt.Printf("middlerow_cnt = %d\n", middlerowCnt)
	fmt.Printf("bottomrow_cnt = %d\n", bottomrowCnt)

	fmt.Println("-----------------------------------")

	//_ = getRails(s)
	//printRail(s)

	//fmt.Printf("EncryptRail(s) = %v\n", EncryptRail(s))

	s2 := "WECRLTEERDSOEEFEAOCAIVDEN"
	fmt.Printf("DecryptRail(s) = %v\n", DecryptRail(s2))
	*/
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

func DecryptRail2(s string rails int) string {
	ret := ""

	slen := len(s)
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

// BuildRailIndexes returns a 2 dimensional slice of integers representing the encoding rails
func EncryptRail2(s string, rails int) string {

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

// DecryptRail performs rail decryption
func DecryptRail(s string) string {
	fmt.Printf("Welcome to DecryptRail() s=%v\n", s)

	toprowCnt := ((len(s) - 1) / 4) + 1
	middlerowCnt := (len(s)) / 2
	bottomrowCnt := (len(s) + 1) / 4
	fmt.Println(toprowCnt, middlerowCnt, bottomrowCnt)

	ptrTop := 0
	ptrMiddle := toprowCnt
	ptrBottom := toprowCnt + middlerowCnt

	fmt.Println(ptrTop, ptrMiddle, ptrBottom)
	fmt.Println("checkyoself")

	newArr := make([]byte, len(s))

	for x := 0; x < toprowCnt; x++ {
		newArr[x*4] = s[x]
	}

	fmt.Println("Top Row populated: ")
	fmt.Println(string(newArr))

	/*for x := ptrMiddle; x < bot; x++ {
		newArr[x*2+1] = s[ptrMiddle]
		ptrMiddle++
	}*/

	fmt.Println("Middle Row Populated: ")
	fmt.Println(string(newArr))

	fmt.Print("test")
	return ""

}

// EncryptRail performs Rail Encryption
func EncryptRail(s string) string {
	ret := ""
	tmp := [3][]int{}

	for x := 0; x < len(s); x++ {
		if (x % 4) == 0 { // top row
			tmp[0] = append(tmp[0], x)
		}

		if (x % 2) != 0 { // middle row
			tmp[1] = append(tmp[1], x)
		}

		if ((x - 2) % 4) == 0 { // bottom row
			tmp[2] = append(tmp[2], x)
		}
	}

	for x := 0; x < len(tmp[0]); x++ {
		ret += string(s[tmp[0][x]])
		//fmt.Printf("%c   ", s[tmp[0][x]])
	}

	for x := 0; x < len(tmp[1]); x++ {
		ret += string(s[tmp[1][x]])
		//fmt.Printf(" %c", s[tmp[1][x]])

	}

	for x := 0; x < len(tmp[2]); x++ {
		ret += string(s[tmp[2][x]])
		//fmt.Printf("%c   ", s[tmp[2][x]])
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
