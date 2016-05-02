package transaction

import (
    "testing"
)

func newTestResponse(responseCode string) *TransactionResponse {
    return &TransactionResponse{ResponseCode: responseCode}
}

func testResponseCodeMsg(testCode string, msg string) bool {
    resp := newTestResponse(testCode)
    respMsg := GetMessage(resp.ResponseCode)
    return respMsg == msg
}

func TestResponseCodeMsg(t *testing.T) {
    if !testResponseCodeMsg("E00029", "Payment information is required.") {
        t.Fail()
    }
    if !testResponseCodeMsg("E00103", "Customer profile creation failed. This payment method does not support profile creation.") {
        t.Fail()
    }
    if !testResponseCodeMsg("87", "Transactions of this market type cannot be processed on this system.") {
        t.Fail()
    }
    if !testResponseCodeMsg("285", "The auction buyer ID is invalid.") {
        t.Fail()
    }
    if !testResponseCodeMsg("2008", "The referenced transaction does not meet the criteria for issuing a Continued Authorization.") {
        t.Fail()
    }
    if !testResponseCodeMsg("2109", "PayPal has rejected the transaction. x_paypal_payflowcolor must be a 6 character hexadecimal value.") {
        t.Fail()
    }
}
