package band

import "github.com/bandprotocol/oracle-consumer/obi"

type Response struct {
	Symbol       string
	ResponseCode uint8
	Rate         uint64
}

func DecodeResult(result []byte) ([]Response, error) {
	var out []Response
	if err := obi.Decode(result, &out); err != nil {
		return nil, err
	}
	return out, nil
}
