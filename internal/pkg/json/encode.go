package json

import "encoding/json"

func Encode(from any) ([]byte, error) {
	return json.Marshal(from)
}
