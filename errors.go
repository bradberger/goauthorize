package authorize

import (
    "errors"
)

// ErrorResponse is returned if the transaction could not be attempted due to some error.
type ErrorResponse struct {
	Messages ErrorResponseCodes `json:"messages"`
}

type ErrorResponseCodes struct {
	ResultCode string `json:"resultCode"`
	Message []ErrorResponseDetails `json:"message"`
}

type ErrorResponseDetails struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

func (e *ErrorResponse) Error() error {
	if e.Messages.ResultCode == "Error" {
		for _, details := range e.Messages.Message {
			return errors.New(details.Text)
		}
	}
	return nil
}
