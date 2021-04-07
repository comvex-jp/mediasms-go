package mediasms

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/comvex-jp/mediasms-go/translations"

	"github.com/comvex-jp/mediasms-go/errors"
	"github.com/comvex-jp/mediasms-go/models"
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

// makeRequest is a generic handler for api calls
func (c Client) makeRequest(requestMethod, url string, body interface{}) ([]byte, error) {
	httpClient := &http.Client{}

	jsonValue, _ := json.Marshal(body)

	req, _ := http.NewRequest(requestMethod, url, bytes.NewBuffer(jsonValue))
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)

	if err != nil {
		return jsonValue, err
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return responseBody, err
	}

	return responseBody, nil
}

// Send request to media4u
func (c Client) Send(messageID string, val models.BuildRequest) (models.APIResponse, error) {
	smsID := createSMSID(c.Prefix, messageID)

	val.SMSID = smsID

	results, err := c.makeRequest("POST", SMSURL+"api/", val)

	if err != nil {
		return models.APIResponse{}, err
	}

	var sendResult resultCode
	unmarshalErr := json.Unmarshal(results, &sendResult)

	if err != nil {
		return models.APIResponse{}, unmarshalErr
	}

	var res = errors.ResultsMapper[sendResult.Result]

	response := models.APIResponse{
		StatusCode:  sendResult.Result,
		Name:        res["name"],
		Description: res["description"],
	}

	return response, nil
}

// GetStatus of a sent sms
func (c Client) GetStatus(messageID string) (models.APIResponse, error) {
	smsID := createSMSID(c.Prefix, messageID)

	results, err := c.makeRequest("GET", SMSURL+"api5/?smsid="+smsID, nil)

	if err != nil {
		return models.APIResponse{}, err
	}

	s := string(results)
	t := strings.Split(s, "\n")

	status, err := strconv.Atoi(t[0])

	if err != nil {
		return models.APIResponse{}, err
	}

	if status == 200 {
		results := models.APIResponse{
			StatusCode:  "200",
			Name:        "Success",
			Description: translations.TranslationMap[status],
		}
		return results, nil
	}

	var getResult resultCode
	unmarshalErr := json.Unmarshal(results, &getResult)

	if unmarshalErr != nil {
		return models.APIResponse{}, nil
	}

	res := errors.ResultsMapper[getResult.Result]

	response := models.APIResponse{
		StatusCode:  getResult.Result,
		Name:        res["name"],
		Description: res["description"],
	}

	return response, nil
}

// ReplaceMessageBodyURLs removes urls and replaces them with mediasms acceptable values
func ReplaceMessageBodyURLs(messageBody string, allURLs []string) string {
	urlReplacements := []string{"{URL}", "{URL2}", "{URL3}", "{URL4}"}
	
	splitBody := strings.Split(messageBody, " ")
	
	urlHashmap := make(map[string]bool)
	
	for _, url := range allURLs {
		urlHashmap[url] = true
	}
	
	replacedMessageBody := []string{}
	
	replacementURLIndex := 0
	
	for _, w := range splitBody {
		if _, ok := urlHashmap[w]; ok {
			replacedMessageBody = append(replacedMessageBody, urlReplacements[replacementURLIndex])
			replacementURLIndex += 1
			continue
		}
		
		replacedMessageBody = append(replacedMessageBody, w)
	}

	return strings.Join(replacedMessageBody, " ")
}