package protobuf

import "encoding/json"

func NewAny(v interface{}) *Any {
	s, err := json.Marshal(v)
	if err != nil {
		return &Any{
			B: []byte(""),
		}
	}
	return &Any{
		B: s,
	}
}

func (x *Any) MarshalJSON() ([]byte, error) {
	return x.B, nil
}

func (x *Any) UnmarshalJSON(v []byte) error {
	x.B = v
	return nil
}
