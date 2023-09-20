package errorcode

type ErrorCode int

// Define costume error code
const (
	ErrNotFound ErrorCode = iota + 1
	ErrParams
	ErrInternal
	ErrUnknown
)

var ErrorMessages = map[ErrorCode]string{
	ErrNotFound: "data not found",
	ErrParams:   "params is not suitable",
	ErrInternal: "internal server error",
	ErrUnknown:  "error unknown",
}

func (e ErrorCode) String() string {
	msg, ok := ErrorMessages[e]
	if !ok {
		return "Unknown error"
	}

	return msg
}
