package transaction

import (
	"regexp"

	"github.com/bradberger/goauthorize"
)

var (
	cardRegex *regexp.Regexp
)

func init() {
	cardRegex = regexp.MustCompile("[^0-9]+")
}

// Request is the high-level container for all transaction requests.
type Request struct {
	CreateTransactionRequest `json:"createTransactionRequest,omitempty"`
}

// CreateTransactionRequest holds the necessary data to create a transaction.
type CreateTransactionRequest struct {
	MerchantAuthentication *authorize.MerchantAuthentication `json:"merchantAuthentication"`
	RefID                  string                            `json:"refId,omitempty"`
	TransactionRequest     *Transaction                      `json:"transactionRequest"`
}

type Order struct {
	InvoiceNumber string `json:"invoiceNumber"`
	Description   string `json:"description"`
}

// Transaction holds data related to the transaction itself.
type Transaction struct {
	TransactionType     string                 `json:"transactionType"`
	Amount              float64                `json:"amount,string"`
	Payment             Payment                `json:"payment"`
	LineItems           map[string]LineItem    `json:"lineItems,omitempty"`
	Tax                 *Tax                   `json:"tax,omitempty"`
	Duty                *Duty                  `json:"duty,omitempty"`
	Shipping            *Shipping              `json:"shipping,omitempty"`
	PONumber            string                 `json:"poNumber,omitempty"`
	BillTo				*BillTo			       `json:"billTo,omitempty"`
	ShipTo              *ShipTo                `json:"shipTo,omitempty"`
	Order               *Order                 `json:"order,omitempty"`
	CustomerIP          string                 `json:"customerIP,omitempty"`
	TransactionSettings map[string]*Setting    `json:"transactionSettings,omitempty"`
	UserFields          UserField 		       `json:"userFields,omitempty"`
	RefTransID          string                 `json:"refRansId,omitempty"`
}

// SetType is a helper function to set the TransactionType
func (t *Transaction) SetType(transType string) {
	t.TransactionType = transType
}

// Payment is the outer container for payment information. Right now it only
// supports CreditCard data.
type Payment struct {
	CreditCard CreditCard `json:"creditCard"`
}

// CreditCard holds the credit card details.
type CreditCard struct {
	CardNumber     string `json:"cardNumber"`
	ExpirationDate string `json:"expirationDate"`
	CardCode       string `json:"cardCode"`
}

func NewCreditCard(number, expiration, code string) *CreditCard {
	return &CreditCard{cardRegex.ReplaceAllString(number, ""), expiration, code}
}

// LineItem is the struct which contains line item details.
type LineItem struct {
	ItemID      string  `json:"itemID"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    int64   `json:"quantity,string"`
	UnitPrice   float64 `json:"unitPrice,string"`
}

// Tax is an optional struct that can be supplied when creating a transaction.
type Tax struct {
	Amount      float64 `json:"amount,string"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}

// Duty is an optional struct that can be supplied when creating a transaction.
type Duty struct {
	Amount      float64 `json:"amount,string"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}

// Shipping is an optional field that can be supplied when creating a transaction.
type Shipping struct {
	Amount      float64 `json:"amount,string"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}

// BillTo is optional data that can set the bill to parameters of the transaction.
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

// ShipTo is optional data that can set the bill to parameters of the transaction.
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

// Setting holds setting data.
type Setting struct {
	SettingName  string `json:"settingName"`
	SettingValue string `json:"settingValue"`
}

// UserField holds details about the user.
type UserField struct {
	UserField []UserFieldPair `json:"userField"`

}

type UserFieldPair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// AuthOnly creates and authorize-only transaction.
func AuthOnly(api *authorize.API, refID string, amount float64, card CreditCard) (r Response, err error) {

	// refID max length is 20 characters, so make sure it's only that.
	if len(refID) > 20 {
		refID = refID[:20]
	}

	trans := Request{
		CreateTransactionRequest{
			MerchantAuthentication: api.MerchantAuthentication,
			RefID: refID,
			TransactionRequest: &Transaction{
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

func SendTransaction(api *authorize.API, refID string, t *Transaction) (r Response, err error) {

	// refID max length is 20 characters, so make sure it's only that.
	if len(refID) > 20 {
		refID = refID[:20]
	}

	trans:= Request{
		CreateTransactionRequest{
			MerchantAuthentication: api.MerchantAuthentication,
			RefID: refID,
			TransactionRequest: t,
		},
	}

	if err = api.Do(&trans, &r); err != nil {
		return
	}

	return

}

// Charge autorizes and captures the given transaction.
func Charge(api *authorize.API, refID string, amount float64, card CreditCard) (r Response, err error) {

	// refID max length is 20 characters, so make sure it's only that.
	if len(refID) > 20 {
		refID = refID[:20]
	}

	trans := Request{
		CreateTransactionRequest{
			MerchantAuthentication: api.MerchantAuthentication,
			RefID: refID,
			TransactionRequest: &Transaction{
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

// Capture captures a previously authorized transaction with the given transactionID.
func Capture(api *authorize.API, refID string, transactionID string, amt float64) (r Response, err error) {

	// refID max length is 20 characters, so make sure it's only that.
	if len(refID) > 20 {
		refID = refID[:20]
	}

	trans := Request{
		CreateTransactionRequest{
			MerchantAuthentication: api.MerchantAuthentication,
			RefID: refID,
			TransactionRequest: &Transaction{
				TransactionType: "priorAuthCaptureTransaction",
				Amount:          amt,
				RefTransID:      transactionID,
			},
		},
	}

	err = api.Do(&trans, &r)
	return

}
