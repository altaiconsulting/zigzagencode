Package zigzagencode implements zigzag encoding and decoding functions that return an error if the requested return type does not have sufficient bits to hold binary representation of the encoded/decoded number.

Package zigzagencode requires Go version 1.22 or above

## Getting zigzagencode
```go get github.com/altaiconsulting/zigzagencode```

## Examples
```
package main

import (
	"altaiconsulting/zigzagencode"
	"fmt"
)

func main() {
	var n8 int8
	var n16 int16
	var n32 int32
	var n64 int64
	var encoded8 uint8
	var encoded16 uint16
	var encoded32 uint32
	var encoded64 uint64
	var err error

	n8 = -1
	encoded8, err = zigzagencode.Encode[int8, uint8](n8)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Number %d is encoded into %d\n", n8, encoded8)
	}
	n8, err = zigzagencode.Decode[uint8, int8](encoded8)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Number %d is decoded into %d\n", encoded8, n8)
	}

	n8 = -3
	encoded32, err = zigzagencode.Encode[int8, uint32](n8)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Number %d is encoded into %d\n", n8, encoded32)
	}
	n8, err = zigzagencode.Decode[uint32, int8](encoded32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Number %d is decoded into %d\n", encoded32, n8)
	}

	n32 = -3
	encoded8, err = zigzagencode.Encode[int32, uint8](n32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Number %d is encoded into %d\n", n32, encoded8)
	}
	n32, err = zigzagencode.Decode[uint8, int32](encoded8)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Number %d is decoded into %d\n", encoded8, n32)
	}

	n32 = -500
	encoded8, err = zigzagencode.Encode[int32, uint8](n32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Number %d is encoded into %d\n", n32, encoded8)
	}

	encoded32, err = zigzagencode.Encode[int32, uint32](n32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Number %d is encoded into %d\n", n32, encoded32)
	}
	n32, err = zigzagencode.Decode[uint32, int32](encoded32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Number %d is decoded into %d\n", encoded32, n32)
	}

	n32 = -20
	encoded32, err = zigzagencode.Encode[int32, uint32](n32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Number %d is encoded into %d\n", n32, encoded32)
	}
	n32, err = zigzagencode.Decode[uint32, int32](encoded32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Number %d is decoded into %d\n", encoded32, n32)
	}

	encoded32 = 999
	n8, err = zigzagencode.Decode[uint32, int8](encoded32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Number %d is decoded into %d\n", encoded32, n8)
	}

	n64 = -40000
	encoded64, err = zigzagencode.Encode[int64, uint64](n64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Number %d is encoded into %d\n", n64, encoded64)
	}
	n64, err = zigzagencode.Decode[uint64, int64](encoded64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Number %d is decoded into %d\n", encoded64, n64)
	}

	n64 = -40000
	encoded16, err = zigzagencode.Encode[int64, uint16](n64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Number %d is encoded into %d\n", n64, encoded16)
	}

	encoded64 = 79999
	n16, err = zigzagencode.Decode[uint64, int16](encoded64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Number %d is decoded into %d\n", encoded64, n16)
	}
}
```