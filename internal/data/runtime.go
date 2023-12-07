package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

// MarshalJSON is a custom marshaler for the Runtime type
// that formats the output to a string representation of
// the runtime in minutes
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}
