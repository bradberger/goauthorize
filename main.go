// Package authorize provides Go bindings for the Authorize.net JSON api. Right
// now it's quite incomplete, but useful for auth and charge actions. Feel free
// to jump in and help out to add more functionality.
package authorize

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// API is the high level structure that holds authentication information.
type API struct {
	MerchantAuthentication *MerchantAuthentication
	Sandbox                bool
}

// MerchantAuthentication is the required struct for authenticating with the Authorize.net API.
type MerchantAuthentication struct {
	Name           string `json:"name"`
	TransactionKey string `json:"transactionKey"`
}

// GetEndpointURL returns the endpoint URL depending on the Sandbox flag.
func (a *API) GetEndpointURL() string {
	if a.Sandbox {
		return "https://apitest.authorize.net/xml/v1/request.api"
	}
	return "https://api.authorize.net/xml/v1/request.api"
}

// SetCredentials sets the credentials for connecting with the Authorize.net API.
func (a *API) SetCredentials(loginID, transactionKey string) {
	a.MerchantAuthentication = &MerchantAuthentication{loginID, transactionKey}
}

// SetTestMode toggles test mode according to supplied value.
func (a *API) SetTestMode(t bool) {
	a.Sandbox = t
}

// GetClient returns an http.Client for connecting with the API.
func (a *API) GetClient() *http.Client {
	return &http.Client{}
}

// Do sends the request to the API. It's a high-level command, and better to use
// other methods if they exist for the action you're looking to execute.
func (a *API) Do(data interface{}, out interface{}) (err error) {

	url := a.GetEndpointURL()
	payload, err := json.Marshal(data)
	if err != nil {
		return
	}

	log.Printf("Sending to %s: %s\n", url, payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	client := a.GetClient()

	r, err := client.Do(req)
	if err != nil {
		return
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	// Handle BOM encoding.
	body = bytes.TrimPrefix(body, []byte{239, 187, 191})
	log.Printf("API response body: %s\n", body)

	// First, see if it's a generic error response. If so, return it.
	errResponse := ErrorResponse{}
	if err = json.Unmarshal(body, &errResponse); err == nil {
		if errResponse.Error() != nil {
			return errResponse.Error()
		}
	}

	err = json.Unmarshal(body, &out)
	return

}
