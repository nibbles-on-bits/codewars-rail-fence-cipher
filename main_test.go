package main

import "testing"

type TestCase struct {
	unencrypted string
	encrypted   string
	railCount   int
}

func TestEncryptRail(t *testing.T) {
	testCases := []TestCase{
		{"WEAREDISCOVEREDFLEEATONCE", "WECRLTEERDSOEEFEAOCAIVDEN", 3},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "AIQYBHJPRXZCGKOSWDFLNTVEMU", 5},
	}

	for _, tc := range testCases {
		got := Encrypt(tc.unencrypted, tc.railCount)
		want := tc.encrypted
		if got != want {
			t.Errorf("Encrypt(%v,%d) = >>%v<<; want >>%v<<", tc.unencrypted, tc.railCount, got, want)
		}
	}

}

func TestDecrypt(t *testing.T) {
	testCases := []TestCase{
		{"WEAREDISCOVEREDFLEEATONCE", "WECRLTEERDSOEEFEAOCAIVDEN", 3},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "AIQYBHJPRXZCGKOSWDFLNTVEMU", 5},
	}

	for _, tc := range testCases {
		got := Decrypt(tc.encrypted, tc.railCount)
		want := tc.unencrypted
		if got != want {
			t.Errorf("Decrypt(%v,%d) = >>%v<<; want >>%v<<", tc.encrypted, tc.railCount, got, want)
		}
	}

}
