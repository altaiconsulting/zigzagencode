package zigzagencode

import (
	"fmt"
	"math/bits"
	"reflect"
	"unsafe"

	"golang.org/x/exp/constraints"
)

// Encode encodes signed integer n of type T to unsigned integer of type R
// Encode returns an error if requested return type does not have sufficient bits
// to hold binary representation of the encoded number
func Encode[T constraints.Signed, R constraints.Unsigned](n T) (R, error) {
	shift := unsafe.Sizeof(T(n))*8 - 1
	encoded := uint64((n << 1) ^ (n >> shift))
	rbits := unsafe.Sizeof(R(0)) * 8
	if uintptr(bits.Len64(encoded)) > rbits {
		rtype := reflect.TypeOf(R(0)).Name()
		return R(0), fmt.Errorf("cannot encode number %d into return type %q: binary representation of encoded number %d exceeds the number of bits %d in type %q", n, rtype, encoded, rbits, rtype)
	}
	return R(encoded), nil
}

// Decode decodes unsigned integer n of type T to signed integer of type R
// Decode returns an error if requested return type does nopt have sufficient bits
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
