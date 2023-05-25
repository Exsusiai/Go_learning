package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("../protobuf_test/page_output_test_2.bin")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1024 && i < len(data); i++ {
		fmt.Printf("%02x ", data[i])
	}
	fmt.Println()
}
