package band

import (
	"github.com/bandprotocol/oracle-consumer/obi"
)

type DefaultCalldata struct {
	Symbols            []string
	MinimumSourceCount uint8
}

func EncodeCalldata(symbols []string, minimumSourceCount uint8) ([]byte, error) {
	return obi.Encode(DefaultCalldata{symbols, minimumSourceCount})
}
