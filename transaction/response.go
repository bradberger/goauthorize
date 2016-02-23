package transaction

import (
    "errors"
)

type Response struct {
    TransactionResponse TransactionResponse `json:"transactionResponse"`
    RefID string `json:"refId"`
    Messages MessageResponse `json:"messages"`
}

type TransactionResponse struct {
    ResponseCode string `json:"responseCode"`
    AuthCode string `json:"authCode"`
    AVSResultCode string `json:"avsResultCode"`
    CVVResultCode string `json:"cvvResultCode"`
    CAVVResultCode string `json:"cavvResultCode"`
    TransID string `json:"transId"`
    RefTransID string `json:"refTransId"`
    TestRequest string `json:"testRequest"`
    TransHash string `json:"transHash"`
    AccountNumber string `json:"accountNumber"`
    AccountType string `json:"accounttype"`
    Messages []*Message
    UserFields []*UserField
    Errors []*ErrorMsg
}

type ErrorMsg struct {
    ErrorCode string `json:"errorCode"`
    ErrorText string `json:"errorText"`
}

type MessageResponse struct {
    ResultCode string `json:"resultCode"`
    Messages []*Message `json:"messages"`
}

type Message struct {
    Code string `json:"code"`
    Description string `json:"description"`
}

func (r *Response) Approved() bool {
    return r.TransactionResponse.ResponseCode == "1"
}

func (r *Response) Captured() bool {
    return r.TransactionResponse.ResponseCode == "1"
}

func (r *Response) Error() error {
    if r.Approved() {
        return nil
    }
    if len(r.TransactionResponse.Errors) > 0 {
        for _, e := range r.TransactionResponse.Errors {
            return errors.New(e.ErrorText)
        }
    }
    return errors.New(GetMessage(r.TransactionResponse.ResponseCode))
}
