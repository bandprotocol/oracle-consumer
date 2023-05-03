package band

import "github.com/bandprotocol/oracle-consumer/obi"

type Calldata struct {
	Symbols            []string
	MinimumSourceCount uint8
}

func EncodeCalldata(symbols []string, minimumSourceCount uint8) ([]byte, error) {
	return obi.Encode(Calldata{symbols, minimumSourceCount})
}
