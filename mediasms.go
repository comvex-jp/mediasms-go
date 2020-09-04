package mediasms

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/comvex-jp/mediasms-go/errors"
	"github.com/comvex-jp/mediasms-go/models"
	"github.com/comvex-jp/mediasms-go/translations"
)

// Client struct
type Client struct {
	Username string
	Password string
	Prefix   string
}

// resultCode from media4u
type resultCode struct {
	Result string `json:"result"`
}

func createSMSID(prefix, messageID string) string {
	return prefix + messageID
}

// SMSURL constant for sms-console
const SMSURL = "https://www.sms-console.jp/"

// Send request to media4u
func (c Client) Send(messageID string, val models.SendRequest) (models.APIResponse, error) {
	smsID := createSMSID(c.Prefix, messageID)

	val.SMSID = smsID

	jsonValue, _ := json.Marshal(val)
	client := &http.Client{}

	req, err := http.NewRequest("POST", SMSURL+"api/", bytes.NewBuffer(jsonValue))
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var sendResult resultCode
	json.Unmarshal(body, &sendResult)

	var res = errors.ResultsMapper[sendResult.Result]

	results := models.APIResponse{
		StatusCode:  sendResult.Result,
		Name:        res["name"],
		Description: res["description"],
	}

	return results, err
}

// GetStatus of a sent sms
func (c Client) GetStatus(messageID string) (models.APIResponse, error) {
	smsID := createSMSID(c.Prefix, messageID)

	client := &http.Client{}

	req, err := http.NewRequest("GET", SMSURL+"api5/?smsid="+smsID, nil)
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	s := string(body)
	t := strings.Split(s, "\n")

	if t[0] == "200" {
		results := models.APIResponse{
			StatusCode:  "200",
			Name:        "Success",
			Description: translations.TranslationMap[t[1]],
		}
		return results, err
	}

	var getResult resultCode
	json.Unmarshal(body, &getResult)

	res := errors.ResultsMapper[getResult.Result]

	results := models.APIResponse{
		StatusCode:  getResult.Result,
		Name:        res["name"],
		Description: res["description"],
	}

	return results, err
}
