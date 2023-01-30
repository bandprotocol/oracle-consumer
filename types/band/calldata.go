package band

import (
	"consumer/obi"
)

type Calldata interface {
}

type DefaultCalldata struct {
	Symbols            []string
	MinimumSourceCount uint8
}

func EncodeCalldata(symbols []string, minimumSourceCount uint8) ([]byte, error) {
	return obi.Encode(DefaultCalldata{symbols, minimumSourceCount})
}

func DecodeCalldata(calldata []byte) (DefaultCalldata, error) {
	decodedDefaultCalldata, err := decodeDefaultCalldata(calldata)
	if err != nil {
		return DefaultCalldata{}, err
	}
	return decodedDefaultCalldata, nil
}

func decodeDefaultCalldata(calldata []byte) (DefaultCalldata, error) {
	var decodedCalldata DefaultCalldata
	err := obi.Decode(calldata, &decodedCalldata)
	if err != nil {
		return DefaultCalldata{}, err
	}

	return decodedCalldata, nil
}
