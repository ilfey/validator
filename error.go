package validator

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	errs map[string][]string
}

func (e *Error) Success() bool {
	return len(e.errs) == 0
}

func (e *Error) JSON() []byte {
	_map := make(map[string][]string)
	for k, v := range e.errs {
		for _, msg := range v {
			_map[k] = append(_map[k], fmt.Sprintf(msg, k))
		}
	}

	b, _ := json.Marshal(_map)

	return b
}
