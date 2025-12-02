package collector

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// When unmarshalling JSON from iDRAC, the "xstring" type defined here can be
// one of the following:
// - nil
// - string
// - [{"Member": "VALUE"}]
type xstring string

func (w *xstring) UnmarshalJSON(data []byte) error {
	var x any

	err := json.Unmarshal(data, &x)
	if err != nil {
		return err
	}

	if x == nil {
		*w = xstring("")
		return nil
	}

	s, ok := x.(string)
	if ok {
		*w = xstring(s)
		return nil
	}

	list := x.([]any)
	dict := list[0].(map[string]any)
	s, ok = dict["Member"].(string)
	if ok {
		*w = xstring(s)
		return nil
	}

	*w = xstring("")
	return nil
}

func (w *xstring) String() string {
	return string(*w)
}

type SafeInt int

func (i *SafeInt) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as a string first (in case it's "123")
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		val, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("cannot unmarshal %q to SafeInt: invalid string format: %w", data, err)
		}
		*i = SafeInt(val)
		return nil
	}
	// Try to unmarshal as a raw number (int or float)
	var raw interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("cannot unmarshal %q to SafeInt: unexpected type: %w", data, err)
	}
	switch v := raw.(type) {
	case float64:
		if v != float64(int(v)) { // Check if it has a fractional part
			return fmt.Errorf("cannot unmarshal %q to MyInt: value has fractional part", data)
		}
		*i = SafeInt(v)
		return nil
	case int: // This case might be hit if the JSON unmarshaler decides it's an int from the start
		*i = SafeInt(v)
		return nil
	case json.Number: // If you explicitly enable UseNumber() in the decoder
		val, err := v.Int64()
		if err != nil {
			return fmt.Errorf("cannot unmarshal %q to MyInt: invalid number format (has fractional part?): %w", data, err)
		}
		*i = SafeInt(val)
		return nil
	default:
		return fmt.Errorf("cannot unmarshal %q to MyInt: unsupported type %T", data, v)
	}
}
