// Copyright (c) 2024 ALTAI Consulting, Inc and Aleksey Gershgorin. All rights reserved.
// Use of this source code is governed by MIT license that can be found in the LICENSE file.
package zigzagencode

import "testing"

func Test_int8_uint8(t *testing.T) {
	t.Log("start of int8<->uint8 encoding/decoding test")
	for n := int8(-20); n <= 20; n++ {
		encoded, err := Encode[int8, uint8](n)
		if err != nil {
			t.Errorf("no error was expected while encoding %d but the following error was received:\n %v\n", n, err)
		}
		decoded, err := Decode[uint8, int8](encoded)
		if err != nil {
			t.Errorf("no error was expected while decoding %d but the following error was received:\n %v\n", encoded, err)
		}
		if n != decoded {
			t.Errorf("decoded number %d was not equal to original number %d", decoded, n)
		}
	}
	t.Log("end of int8<->uint8 encoding/decoding test")
}

func Test_int8_uint16(t *testing.T) {
	t.Log("start of int8<->uint16 encoding/decoding test")
	for n := int8(-20); n <= 20; n++ {
		encoded, err := Encode[int8, uint16](n)
		if err != nil {
			t.Errorf("no error was expected while encoding %d but the following error was received:\n %v\n", n, err)
		}
		decoded, err := Decode[uint16, int8](encoded)
		if err != nil {
			t.Errorf("no error was expected while decoding %d but the following error was received:\n %v\n", encoded, err)
		}
		if n != decoded {
			t.Errorf("decoded number %d was not equal to original number %d", decoded, n)
		}
	}
	t.Log("end of int8<->uint16 encoding/decoding test")
}

func Test_int8_uint32(t *testing.T) {
	t.Log("start of int8<->uint32 encoding/decoding test")
	for n := int8(-20); n <= 20; n++ {
		encoded, err := Encode[int8, uint32](n)
		if err != nil {
			t.Errorf("no error was expected while encoding %d but the following error was received:\n %v\n", n, err)
		}
		decoded, err := Decode[uint32, int8](encoded)
		if err != nil {
			t.Errorf("no error was expected while decoding %d but the following error was received:\n %v\n", encoded, err)
		}
		if n != decoded {
			t.Errorf("decoded number %d was not equal to original number %d", decoded, n)
		}
	}
	t.Log("end of int8<->uint32 encoding/decoding test")
}

func Test_int8_uint64(t *testing.T) {
	t.Log("start of int8<->uint64 encoding/decoding test")
	for n := int8(-20); n <= 20; n++ {
		encoded, err := Encode[int8, uint64](n)
		if err != nil {
			t.Errorf("no error was expected while encoding %d but the following error was received:\n %v\n", n, err)
		}
		decoded, err := Decode[uint64, int8](encoded)
		if err != nil {
			t.Errorf("no error was expected while decoding %d but the following error was received:\n %v\n", encoded, err)
		}
		if n != decoded {
			t.Errorf("decoded number %d was not equal to original number %d", decoded, n)
		}
	}
	t.Log("end of int8<->uint64 encoding/decoding test")
}

func Test_int16_uint8(t *testing.T) {
	t.Log("start of int16<->uint8 encoding/decoding test")
	for n := int16(-20); n <= 20; n++ {
		encoded, err := Encode[int16, uint8](n)
		if err != nil {
			t.Errorf("no error was expected while encoding %d but the following error was received:\n %v\n", n, err)
		}
		decoded, err := Decode[uint8, int16](encoded)
		if err != nil {
			t.Errorf("no error was expected while decoding %d but the following error was received:\n %v\n", encoded, err)
		}
		if n != decoded {
			t.Errorf("decoded number %d was not equal to original number %d", decoded, n)
		}
	}
	t.Log("end of int16<->uint8 encoding/decoding test")
}

func Test_int16_uint16(t *testing.T) {
	t.Log("start of int16<->uint16 encoding/decoding test")
	for n := int16(-20); n <= 20; n++ {
		encoded, err := Encode[int16, uint16](n)
		if err != nil {
			t.Errorf("no error was expected while encoding %d but the following error was received:\n %v\n", n, err)
		}
		decoded, err := Decode[uint16, int16](encoded)
		if err != nil {
			t.Errorf("no error was expected while decoding %d but the following error was received:\n %v\n", encoded, err)
		}
		if n != decoded {
			t.Errorf("decoded number %d was not equal to original number %d", decoded, n)
		}
	}
	t.Log("end of int16<->uint16 encoding/decoding test")
}

func Test_int16_uint32(t *testing.T) {
	t.Log("start of int16<->uint32 encoding/decoding test")
	for n := int16(-20); n <= 20; n++ {
		encoded, err := Encode[int16, uint32](n)
		if err != nil {
			t.Errorf("no error was expected while encoding %d but the following error was received:\n %v\n", n, err)
		}
		decoded, err := Decode[uint32, int16](encoded)
		if err != nil {
			t.Errorf("no error was expected while decoding %d but the following error was received:\n %v\n", encoded, err)
		}
		if n != decoded {
			t.Errorf("decoded number %d was not equal to original number %d", decoded, n)
		}
	}
	t.Log("end of int16<->uint32 encoding/decoding test")
}

func Test_int16_uint64(t *testing.T) {
	t.Log("start of int16<->uint64 encoding/decoding test")
	for n := int16(-20); n <= 20; n++ {
		encoded, err := Encode[int16, uint64](n)
		if err != nil {
			t.Errorf("no error was expected while encoding %d but the following error was received:\n %v\n", n, err)
		}
		decoded, err := Decode[uint64, int16](encoded)
		if err != nil {
			t.Errorf("no error was expected while decoding %d but the following error was received:\n %v\n", encoded, err)
		}
		if n != decoded {
			t.Errorf("decoded number %d was not equal to original number %d", decoded, n)
		}
	}
	t.Log("end of int16<->uint64 encoding/decoding test")
}

func Test_int32_uint8(t *testing.T) {
	t.Log("start of int32<->uint8 encoding/decoding test")
	for n := int32(-20); n <= 20; n++ {
		encoded, err := Encode[int32, uint8](n)
		if err != nil {
			t.Errorf("no error was expected while encoding %d but the following error was received:\n %v\n", n, err)
		}
		decoded, err := Decode[uint8, int32](encoded)
		if err != nil {
			t.Errorf("no error was expected while decoding %d but the following error was received:\n %v\n", encoded, err)
		}
		if n != decoded {
			t.Errorf("decoded number %d was not equal to original number %d", decoded, n)
		}
	}
	t.Log("end of int32<->uint8 encoding/decoding test")
}

func Test_int32_uint16(t *testing.T) {
	t.Log("start of int32<->uint16 encoding/decoding test")
	for n := int32(-20); n <= 20; n++ {
		encoded, err := Encode[int32, uint16](n)
		if err != nil {
			t.Errorf("no error was expected while encoding %d but the following error was received:\n %v\n", n, err)
		}
		decoded, err := Decode[uint16, int32](encoded)
		if err != nil {
			t.Errorf("no error was expected while decoding %d but the following error was received:\n %v\n", encoded, err)
		}
		if n != decoded {
			t.Errorf("decoded number %d was not equal to original number %d", decoded, n)
		}
	}
	t.Log("end of int32<->uint16 encoding/decoding test")
}

func Test_int32_uint32(t *testing.T) {
	t.Log("start of int32<->uint32 encoding/decoding test")
	for n := int32(-20); n <= 20; n++ {
		encoded, err := Encode[int32, uint32](n)
		if err != nil {
			t.Errorf("no error was expected while encoding %d but the following error was received:\n %v\n", n, err)
		}
		decoded, err := Decode[uint32, int32](encoded)
		if err != nil {
			t.Errorf("no error was expected while decoding %d but the following error was received:\n %v\n", encoded, err)
		}
		if n != decoded {
			t.Errorf("decoded number %d was not equal to original number %d", decoded, n)
		}
	}
	t.Log("end of int32<->uint32 encoding/decoding test")
}

func Test_int32_uint64(t *testing.T) {
	t.Log("start of int32<->uint64 encoding/decoding test")
	for n := int32(-20); n <= 20; n++ {
		encoded, err := Encode[int32, uint64](n)
		if err != nil {
			t.Errorf("no error was expected while encoding %d but the following error was received:\n %v\n", n, err)
		}
		decoded, err := Decode[uint64, int32](encoded)
		if err != nil {
			t.Errorf("no error was expected while decoding %d but the following error was received:\n %v\n", encoded, err)
		}
		if n != decoded {
			t.Errorf("decoded number %d was not equal to original number %d", decoded, n)
		}
	}
	t.Log("end of int32<->uint64 encoding/decoding test")
}

func Test_int64_uint8(t *testing.T) {
	t.Log("start of int64<->uint8 encoding/decoding test")
	for n := int64(-20); n <= 20; n++ {
		encoded, err := Encode[int64, uint8](n)
		if err != nil {
			t.Errorf("no error was expected while encoding %d but the following error was received:\n %v\n", n, err)
		}
		decoded, err := Decode[uint8, int64](encoded)
		if err != nil {
			t.Errorf("no error was expected while decoding %d but the following error was received:\n %v\n", encoded, err)
		}
		if n != decoded {
			t.Errorf("decoded number %d was not equal to original number %d", decoded, n)
		}
	}
	t.Log("end of int64<->uint8 encoding/decoding test")
}

func Test_int64_uint16(t *testing.T) {
	t.Log("start of int64<->uint16 encoding/decoding test")
	for n := int64(-20); n <= 20; n++ {
		encoded, err := Encode[int64, uint16](n)
		if err != nil {
			t.Errorf("no error was expected while encoding %d but the following error was received:\n %v\n", n, err)
		}
		decoded, err := Decode[uint16, int64](encoded)
		if err != nil {
			t.Errorf("no error was expected while decoding %d but the following error was received:\n %v\n", encoded, err)
		}
		if n != decoded {
			t.Errorf("decoded number %d was not equal to original number %d", decoded, n)
		}
	}
	t.Log("end of int64<->uint16 encoding/decoding test")
}

func Test_int64_uint32(t *testing.T) {
	t.Log("start of int64<->uint32 encoding/decoding test")
	for n := int64(-20); n <= 20; n++ {
		encoded, err := Encode[int64, uint32](n)
		if err != nil {
			t.Errorf("no error was expected while encoding %d but the following error was received:\n %v\n", n, err)
		}
		decoded, err := Decode[uint32, int64](encoded)
		if err != nil {
			t.Errorf("no error was expected while decoding %d but the following error was received:\n %v\n", encoded, err)
		}
		if n != decoded {
			t.Errorf("decoded number %d was not equal to original number %d", decoded, n)
		}
	}
	t.Log("end of int64<->uint32 encoding/decoding test")
}

func Test_int64_uint64(t *testing.T) {
	t.Log("start of int64<->uint64 encoding/decoding test")
	for n := int64(-20); n <= 20; n++ {
		encoded, err := Encode[int64, uint64](n)
		if err != nil {
			t.Errorf("no error was expected while encoding %d but the following error was received:\n %v\n", n, err)
		}
		decoded, err := Decode[uint64, int64](encoded)
		if err != nil {
			t.Errorf("no error was expected while decoding %d but the following error was received:\n %v\n", encoded, err)
		}
		if n != decoded {
			t.Errorf("decoded number %d was not equal to original number %d", decoded, n)
		}
	}
	t.Log("end of int64<->uint64 encoding/decoding test")
}

func TestEncodeError(t *testing.T) {
	t.Log("start of encode error test")
	_, err := Encode[int32, uint8](-500)
	if err == nil {
		t.Error("encode error was expected but was not returned")
	}
	_, err = Encode[int64, uint16](-40000)
	if err == nil {
		t.Error("encode error was expected but was not returned")
	}
	t.Log("end of encode error test")
}

func TestDecodeError(t *testing.T) {
	t.Log("start of decode error test")
	_, err := Decode[uint32, int8](999)
	if err == nil {
		t.Error("decode error was expected but was not returned")
	}
	_, err = Decode[uint64, int16](79999)
	if err == nil {
		t.Error("decode error was expected but was not returned")
	}
	t.Log("start of decode error test")
}

func TestLargeNumberEncodeDecode(t *testing.T) {
	t.Log("large number encoding/decoding test")
	encoded, err := Encode[int64, uint64](-40000)
	if err != nil {
		t.Errorf("no error was expected while encoding %d but the following error was received:\n %v\n", -40000, err)
	}
	decoded, err := Decode[uint64, int64](encoded)
	if err != nil {
		t.Errorf("no error was expected while decoding %d but the following error was received:\n %v\n", encoded, err)
	}
	if decoded != -40000 {
		t.Errorf("decoded number %d was not equal to original number %d", decoded, -40000)
	}
	t.Log("large number encoding/decoding test")
}
