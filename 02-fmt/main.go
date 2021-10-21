package main

import "bytes"

func compareTwoArrayOfBytes() bool {
	bytea := []byte{}
	byteb := []byte{}
	return bytes.Compare(bytea, byteb) > 0
}
