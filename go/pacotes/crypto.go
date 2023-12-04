package main

import (
	"crypto/sha1"
	"fmt"
)

// //
func main() {
	//criar hash
	h := sha1.New()

	//ewcrever dados no hash
	h.Write([]byte("crypto hash"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}
