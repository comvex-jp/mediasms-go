package mediasms

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Client struct
type Client struct {
	Username  string
	Password  string
	AccountID string
}

// Send request to media4u
func (c Client) Send(messageID string) {
	smsID := createSMSID(c.AccountID, messageID)

	form := url.Values{
		"smsid":         {smsID},
		"mobilenumber":  {"07043459253"},
		"status":        {"1"},
		"type":          {"sms"},
		"returnsms":     {"1"},
		"waitreturnsms": {"1"},
		"smstext":       {"Click Here 17 {URL}"},
		"originalurl":   {"https://google.com"},
	}

	// requestBody, err := json.Marshal(map[string]string{
	// 	"mobilenumber":  "07043459253",
	// 	"smstext":       "Click Here 17 {URL}",
	// 	"originalurl":   "https://google.com",
	// 	"smsid":         smsID,
	// 	"status":        "1",
	// 	"type":          "sms",
	// 	"returnsms":     "1",
	// 	"waitreturnsms": "1",
	// })

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://www.sms-console.jp/api/", strings.NewReader(form.Encode()))
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	fmt.Println(s)
}

func createSMSID(accountID, messageID string) string {
	return accountID + messageID
}
