package response

import (
	"NMTMakiCustomsAll/error"
	"encoding/json"
	"net/http"
)

// Status struct to hold the status code and text.
type Status struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

// Response struct to hold the response structure.
type Response struct {
	Status *Status       `json:"status"`
	Error  *error.Error  `json:"error,omitempty"`
	Data   []interface{} `json:"data,omitempty"`
}

// NewResponse creates a new Response object with the given parameters.
func NewResponse(statusCode int, statusText string, errObj *error.Error, responseData []interface{}) *Response {
	return &Response{
		Status: &Status{
			Code: statusCode,
			Text: statusText,
		},
		Error: errObj,
		Data:  responseData,
	}
}

// SendResponse writes a Response object to the HTTP response writer.
func SendResponse(w http.ResponseWriter, statusCode int, statusText string, errObj *error.Error, responseData []interface{}) {
	// Ensure that the error object includes file and line information
	if errObj != nil && errObj.ErrorCaller == nil {
		errObj.ErrorCaller = &error.Caller{
			File: "unknown", // or set default value
			Line: 0,         // or set default value
		}
	}

	response := NewResponse(statusCode, statusText, errObj, responseData)

	// Marshal the response to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	// Set response header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
