package calldata

import (
	"consumer/obi"
)

type Calldata interface {
}

type DefaultCalldata struct {
	Symbols    []string
	Multiplier uint64
}

func EncodeCalldata(symbols []string, multiplier uint64) ([]byte, error) {
	return obi.Encode(DefaultCalldata{symbols, multiplier})
}

func DecodeCalldata(calldata []byte, symbolType string) (DefaultCalldata, error) {
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
