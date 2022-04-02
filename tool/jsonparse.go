package tool

// TODO
// 暂时不用
import (
	"encoding/json"
	"io"
)

type JsonParse struct {
}

func Decode(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(v)
}
