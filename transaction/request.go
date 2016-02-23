package transaction

import (
	"github.com/bradberger/goauthorize"
)

type Request struct {
	CreateTransactionRequest `json:"createTransactionRequest,omitempty"`
}

type CreateTransactionRequest struct {
	MerchantAuthentication *authorize.MerchantAuthentication `json:"merchantAuthentication"`
	RefID                  string                            `json:"refId,omitempty"`
	TransactionRequest     *TransactionRequest               `json:"transactionRequest"`
}

type TransactionRequest struct {
	TransactionType     string                 `json:"transactionType"`
	Amount              float64                `json:"amount,string"`
	Payment             Payment                `json:"payment"`
	LineItems           map[string]LineItem    `json:"lineItems,omitempty"`
	Tax                 *Tax                   `json:"tax,omitempty"`
	Duty                *Duty                  `json:"duty,omitempty"`
	Shipping            *Shipping              `json:"shipping,omitempty"`
	PONumber            string                 `json:"poNumber,omitempty"`
	Customer            *Customer              `json:"customer,omitempty"`
	CustomerIP          string                 `json:"customerIP,omitempty"`
	TransactionSettings map[string]*Setting    `json:"transactionSettings,omitempty"`
	UserFields          map[string][]UserField `json:"userFields,omitempty"`
	RefTransID          string                 `json:"refRansId,omitempty"`
}

func (t *TransactionRequest) SetType(transType string) {
	t.TransactionType = transType
}

type Payment struct {
	CreditCard CreditCard `json:"creditCard"`
}

type CreditCard struct {
	CardNumber     string `json:"cardNumber"`
	ExpirationDate string `json:"expirationDate"`
	CardCode       string `json:"cardCode"`
}

type LineItem struct {
	ItemID      string  `json:"itemID"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    int64   `json:"quantity,string"`
	UnitPrice   float64 `json:"unitPrice,string"`
}

type Tax struct {
	Amount      float64 `json:"amount,string"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}

type Duty struct {
	Amount      float64 `json:"amount,string"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}

type Shipping struct {
	Amount      float64 `json:"amount,string"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}

type Customer struct {
	ID string `json:"id"`
}

type BillTo struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Company   string `json:"company"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
}

type ShipTo struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Company   string `json:"company"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
}

type Setting struct {
	SettingName  string `json:"settingName"`
	SettingValue string `json:"settingValue"`
}

type UserField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func AuthOnly(api *authorize.API, refID string, amount float64, card CreditCard) (r Response, err error) {

	trans := Request{
		CreateTransactionRequest{
			MerchantAuthentication: api.MerchantAuthentication,
			RefID: refID,
			TransactionRequest: &TransactionRequest{
				TransactionType: "authOnlyTransaction",
				Amount:          amount,
				Payment:         Payment{CreditCard: card},
			},
		},
	}

	err = api.Do(&trans, &r)
	if err != nil {
		return
	}

	return

}

func Charge(api *authorize.API, refID string, amount float64, card CreditCard) (r Response, err error) {

	trans := Request{
		CreateTransactionRequest{
			MerchantAuthentication: api.MerchantAuthentication,
			RefID: refID,
			TransactionRequest: &TransactionRequest{
				TransactionType: "authCaptureTransaction",
				Amount:          amount,
				Payment:         Payment{CreditCard: card},
			},
		},
	}

	err = api.Do(&trans, &r)
	if err != nil {
		return
	}

	return

}

func Capture(api *authorize.API, refID string, transactionID string, amt float64) (r Response, err error) {

	trans := Request{
		CreateTransactionRequest{
			MerchantAuthentication: api.MerchantAuthentication,
			RefID: refID,
			TransactionRequest: &TransactionRequest{
				TransactionType: "priorAuthCaptureTransaction",
				Amount:          amt,
				RefTransID:      transactionID,
			},
		},
	}

	err = api.Do(&trans, &r)
	return

}
