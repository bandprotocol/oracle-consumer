package band

import "github.com/bandprotocol/oracle-consumer/obi"

type Responses struct {
	Symbol       string
	ResponseCode uint8
	Rate         uint64
}

func DecodeResult(result []byte) ([]Responses, error) {
	// For decode rates
	type Output struct {
		Responses []Responses
	}
	var out Output
	err := obi.Decode(result, &out)

	if err != nil {
		return nil, err
	}
	return out.Responses, nil
}
