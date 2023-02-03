package json

import (
	"encoding/json"
	"io"
)

func Decode(from io.Reader, to any) error {
	return json.NewDecoder(from).Decode(to)
}
