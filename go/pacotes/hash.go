package main

import (
	"fmt"
	"hash/crc32"

)

func main(){
	//criar hash
	h := crc32.NewIEEE()

	//escrever dados no hash	
	h.Write([]byte("code with hash package"))

	//calculate hash
	v := h.Sum32()
	fmt.Println(v)

}
