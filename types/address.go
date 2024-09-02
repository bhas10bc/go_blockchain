package types

import (
	"encoding/hex"
	"fmt"
)

type Address [20]uint8

func (h Address) To20Slice() []byte {
	b := make([]byte, 20)
	for i := 0; i < 20; i++ {
		b[i] = h[i]
	}
	return b
}
func (a Address) String() string {
	return hex.EncodeToString(a.To20Slice())
}

func AddressFromBytes(b []byte) Address {
	if len(b) != 20 {
		msg := fmt.Sprintf("the provided bytes length %d should be 32", len(b))
		panic(msg)
	}

	var value [20]uint8
	for i := 0; i < 20; i++ {
		value[i] = b[i]
	}
	return Address(value)
}