package main

import "testing"

func TestEncryptRail(t *testing.T) {
	testCases := [][2]string{
		{"WEAREDISCOVEREDFLEEATONCE", "WECRLTEERDSOEEFEAOCAIVDEN"},
	}

	for _, a := range testCases {
		got := Encrypt(a[0])
		if got != a[1] {
			t.Errorf("EncryptRail(%v) = >>%v<<; want >>%v<<", a[0], got, a[1])
		}
	}

}

func TestDecryptRail(t *testing.T) {
	testCases := [][2]string{
		{"WEAREDISCOVEREDFLEEATONCE", "WECRLTEERDSOEEFEAOCAIVDEN"},
	}
	for _, a := range testCases {
		got := DecryptRail(a[1])
		if got != a[0] {
			t.Errorf("DecryptRail(%v) = >>%v<<; want >>%v<<", a[1], got, a[0])
		}
	}

}
