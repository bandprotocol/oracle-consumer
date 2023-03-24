package band

import "github.com/bandprotocol/oracle-consumer/obi"

type Response struct {
	Symbol       string
	ResponseCode uint8
	Rate         uint64
}

func DecodeResult(result []byte) ([]Response, error) {
	// For decode rates
	type Output struct {
		Responses []Response
	}
	var out Output
	err := obi.Decode(result, &out)

	if err != nil {
		return nil, err
	}
	return out.Responses, nil
}
