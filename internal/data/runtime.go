package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// defibe ab error for UnmarshalJSON() method can return if we're unable tp parse
// or convert the JSON string successfully
var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

// format for "<runtime> mins"
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	// need surrounded by double quotes for JSON valid
	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}

// implement UnmarshalJSON() method on the runtime type so that it satisfies the json.Unmarshaler interface.
// IMPORTANT: Beacuse UnmarshalJSON() needs to modify the reciever (our Runtime type), must
// use a pointer receiver for this to work correctly
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	// expect that the incoming JSON value will be a string in the format "<runtime> mins"
	// remove the surounding double quote from sting

	unquotedJSONvalue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	// split the string to isolate the part containing the number.
	parts := strings.Split(unquotedJSONvalue, " ")

	// Sanity check the parts of the string to make sure it was in the expected format
	// If t isn't, we return ErrInvalidRuntimeFormat error again.
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	// Otherwise, parse the string containing the number into an int32.
	// Again, if this fails return the ErrInvalidRuntimeFormat error.
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	// Convert the int32 to Runtime type and assign this to the receiver.
	*r = Runtime(i)

	return nil
}
