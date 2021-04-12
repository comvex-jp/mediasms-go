package mediasms

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"unicode"

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

const (
	URLRegex                = `(?:https?:\/\/)?[-\w@:%_+.~#?&=]{2,256}\.[a-z]{2,4}(?:[-\w@:%_\x{30A1}-\x{30FC}\p{Hiragana}\p{Han}}+.~#?&\/=,()\][|]*)?`
	InvalidURLRegex         = `(?:[\w]+)(?:https?:\/\/)[-\w@:%_+.~#?&=]{2,256}\.[a-z]{2,4}(?:[-\w@:%_\x{30A1}-\x{30FC}\p{Hiragana}\p{Han}}+.~#?&\/=,()\][|]*)?`
	PlaceholderBracketRegex = `\*\|(.*?)\|\*`
)

// ExtractURLsFromText finds and returns all urls
func ExtractURLsFromText(t string) []string {
	r := regexp.MustCompile(URLRegex)

	return r.FindAllString(t, -1)
}

// ProcessAndExtractURLs removes placeholder brackets, invalidURLs and returns only valid URLs
func ProcessAndExtractURLs(t string) []string {
	placeholderOmitted := replaceElementsinPlaceholderBrackets(t, "")

	invalidURLsOmittied := replaceInvalidURLsInText(placeholderOmitted, "")

	return ExtractURLsFromText(invalidURLsOmittied)
}

// replaceElementsinPlaceholderBrackets find and replace anything between *||*
func replaceElementsinPlaceholderBrackets(t, replacedWith string) string {
	r := regexp.MustCompile(PlaceholderBracketRegex)

	return r.ReplaceAllString(t, replacedWith)
}

// replaceInvalidURLsInText find and replace invalid URLs, e.g. aahttp://... 01http://...
func replaceInvalidURLsInText(t, replaceWith string) string {
	r := regexp.MustCompile(InvalidURLRegex)

	return r.ReplaceAllString(t, replaceWith)
}

func split(r rune) bool {
	return unicode.IsSpace(r)
}

// ReplaceMessageBodyURLs removes urls and replaces them with mediasms acceptable values
func ReplaceMessageBodyURLs(messageBody string, allURLs []string) string {
	urlReplacements := []string{"{URL}", "{URL2}", "{URL3}", "{URL4}"}

	splitBody := strings.FieldsFunc(messageBody, split)

	urlHashmap := make(map[string]bool)

	for _, url := range allURLs {
		urlHashmap[url] = true
	}

	replacedMessageBody := []string{}

	replacementURLIndex := 0

	for _, w := range splitBody {
		extractedURL := ProcessAndExtractURLs(w)

		if len(ProcessAndExtractURLs(w)) == 1 {

			if _, ok := urlHashmap[extractedURL[0]]; ok && replacementURLIndex != len(urlReplacements) {

				replacement := strings.Replace(w, extractedURL[0], urlReplacements[replacementURLIndex], -1)
				replacedMessageBody = append(replacedMessageBody, replacement)

				replacementURLIndex += 1
				continue
			}
		}

		replacedMessageBody = append(replacedMessageBody, w)
	}

	return strings.Join(replacedMessageBody, " ")
}
