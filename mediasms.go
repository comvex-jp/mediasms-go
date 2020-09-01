package mediatest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Client struct
type Client struct {
	Username  string
	Password  string
	AccountID string
}

// Send request to media4u
func (c Client) Send(messageID string) (string, error) {
	smsID := createSMSID(c.AccountID, messageID)

	values := map[string]string{
		"smstext":      "Click Here 4 {URL}",
		"mobilenumber": "07043459253",
		"originalurl":  "https://google.com",
		"smsid":        smsID,
		"status":       "1",
		"type":         "sms",
	}

	jsonValue, _ := json.Marshal(values)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://www.sms-console.jp/api/", bytes.NewBuffer(jsonValue))
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Content-Type", "application/json")
	fmt.Println(req)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s, err
}

// GetStatus of a sent sms
func (c Client) GetStatus(messageID string) (string, error) {
	smsID := createSMSID(c.AccountID, messageID)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.sms-console.jp/api5/?smsid="+smsID, nil)
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Content-Type", "application/json")
	fmt.Println(req)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)

	return s, err

}

func createSMSID(accountID, messageID string) string {
	return accountID + messageID
}
