// Copyright (c) 2024 ALTAI Consulting, Inc and Aleksey Gershgorin. All rights reserved.
// Use of this source code is governed by MIT license that can be found in the LICENSE file.

// Package zigzagencode implements zigzag encoding and decoding functions
// that return an error if the requested return type does not have sufficient
// bits to hold binary representation of the encoded/decoded number
package zigzagencode

import (
	"encoding/binary"
	"fmt"
	"math/bits"
	"reflect"
	"unsafe"

	"golang.org/x/exp/constraints"
)

func removeLeadingOnes(n uint64, extraBits int) uint64 {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, uint64(n))
	for i, byteNum := 0, extraBits>>3; i < byteNum; i++ {
		if bytes[i] == 0xFF {
			bytes[i] = 0
		}
	}
	return binary.BigEndian.Uint64(bytes)
}

// Encode encodes signed integer n of type T to unsigned integer of type R
// Encode returns an error if requested return type does not have sufficient bits
// to hold binary representation of the encoded number
func Encode[T constraints.Signed, R constraints.Unsigned](n T) (R, error) {
	shift := unsafe.Sizeof(T(n))*8 - 1
	encoded := uint64((n << 1) ^ (n >> shift))
	tbits, rbits := int(unsafe.Sizeof(T(0))*8), int(unsafe.Sizeof(R(0))*8)
	if rbits > tbits {
		tbits = rbits
	}
	extraBits := bits.Len64(encoded) - tbits
	if extraBits > 0 {
		extraBits := (extraBits >> 3) << 3
		if extraBits > 0 {
			encoded = removeLeadingOnes(encoded, extraBits)
		}
	}
	if bits.Len64(encoded) > rbits {
		rtype := reflect.TypeOf(R(0)).Name()
		return R(0), fmt.Errorf("cannot encode number %d into return type %q: binary representation of encoded number %d exceeds the number of bits %d in type %q", n, rtype, encoded, rbits, rtype)
	}
	return R(encoded), nil
}

// Decode decodes unsigned integer n of type T to signed integer of type R
// Decode returns an error if requested return type does not have sufficient bits
// to hold binary representation of the decoded number
func Decode[T constraints.Unsigned, R constraints.Signed](n T) (R, error) {
	n64 := uint64(n)
	decoded := int64((n64 >> 1) ^ (-(n64 & 1)))
	rvalue := R(decoded)
	if int64(rvalue) != decoded {
		rbits := unsafe.Sizeof(R(0)) * 8
		rtype := reflect.TypeOf(R(0)).Name()
		return R(0), fmt.Errorf("cannot decode number %d into return type %q: binary representation of decoded number %d exceeds the number of bits %d in type %q", n, rtype, decoded, rbits, rtype)
	}
	return rvalue, nil
}
