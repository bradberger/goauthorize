package authorize

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
)

type API struct {
    MerchantAuthentication *MerchantAuthentication
    Sandbox bool
}

type MerchantAuthentication struct {
    Name string `json:"name"`
    TransactionKey string `json:"transactionKey"`
}

func (a *API) GetEndpointURL() string {
    if a.Sandbox {
        return "https://apitest.authorize.net/xml/v1/request.api"
    }
    return "https://api.authorize.net/xml/v1/request.api"
}

func (a *API) SetCredentials(loginID, transactionKey string) {
    a.MerchantAuthentication = &MerchantAuthentication{loginID, transactionKey}
}

func (a *API) SetTestMode(t bool) {
    a.Sandbox = t
}

func (a *API) GetClient() *http.Client {
    return &http.Client{}
}

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
    log.Printf("%s\n", body)

    err = json.Unmarshal(body, &out)
    return

}
