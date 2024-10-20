package error

type Caller struct {
	File string `json:"file"`
	Line int    `json:"line"`
}

type Error struct {
	ErrorCode    int     `json:"error_code"`
	ErrorStatus  string  `json:"error_status"`
	ErrorMessage string  `json:"error_message"`
	ErrorData    string  `json:"error_data"`
	ErrorCaller  *Caller `json:"error_caller"`
}

func NewError(errorCode int, errorStatus string, errorMessage string, errorData string, file string, line int) *Error {
	return &Error{
		ErrorCode:    errorCode,
		ErrorStatus:  errorStatus,
		ErrorMessage: errorMessage,
		ErrorData:    errorData,
		ErrorCaller: &Caller{
			File: file,
			Line: line,
		},
	}
}
