package mediasms

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Client struct
type Client struct {
	Username string
	Password string
	Prefix   string
}

// SendRequest struct
type SendRequest struct {
	Smsid         string `json:"smsid"`
	Smstitle      string `json:"smstitle"`
	Smstext       string `json:"smstext"`
	Mobilenumber  string `json:"mobilenumber"`
	Originalurl   string `json:"originalurl"`
	Originalurl2  string `json:"originalurl2"`
	Originalurl3  string `json:"originalurl3"`
	Originalurl4  string `json:"originalurl4"`
	Status        string `json:"status"`
	Returnsms     string `json:"returnsms"`
	Waitreturnsms string `json:"waitreturnsms"`
	Type          string `json:"type"`
}

// APIResponse struct
type APIResponse struct {
	StatusCode  string
	Name        string
	Description string
}

// ResultCode from media4u
type ResultCode struct {
	Result string `json:"result"`
}

// WebHook from media4u
type WebHook struct {
	Mobilenumber      string
	Status            string
	Smsid             string
	Returnsms         string
	Raitreturnsms     string
	Returnsmsdatetime string
	Replyid           string
	Senderid          string
}

const smsurl = "https://www.sms-console.jp/"

// Send request to media4u
func (c Client) Send(messageID string, val SendRequest) (APIResponse, error) {
	smsID := createSMSID(c.Prefix, messageID)

	val.Smsid = smsID

	jsonValue, _ := json.Marshal(val)
	client := &http.Client{}
	req, err := http.NewRequest("POST", smsurl+"api/", bytes.NewBuffer(jsonValue))
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var sendResult ResultCode
	json.Unmarshal(body, &sendResult)

	var res = sendResultsMapper[sendResult.Result]

	results := APIResponse{
		StatusCode:  sendResult.Result,
		Name:        res["name"],
		Description: res["description"],
	}

	return results, err
}

// GetStatus of a sent sms
func (c Client) GetStatus(messageID string) (APIResponse, error) {
	smsID := createSMSID(c.Prefix, messageID)

	client := &http.Client{}
	req, err := http.NewRequest("GET", smsurl+"api5/?smsid="+smsID, nil)
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
		results := APIResponse{
			StatusCode:  "200",
			Name:        "Success",
			Description: translationMap[t[1]],
		}
		return results, err
	}

	var getResult ResultCode
	json.Unmarshal(body, &getResult)

	res := getResultsMapper[getResult.Result]

	results := APIResponse{
		StatusCode:  getResult.Result,
		Name:        res["name"],
		Description: res["description"],
	}

	return results, err
}

func createSMSID(prefix, messageID string) string {
	return prefix + messageID
}

var sendResultsMapper = map[string]map[string]string{
	"200": {
		"name":        "Success code",
		"description": "Sent SMS",
	},
	"401": {
		"name":        "Authentication error",
		"description": "Authorization Required",
	},
	"414": {
		"name":        "Failed to send SMS",
		"description": "URL is longer than 8,190 byte",
	},
	"550": {
		"name":        "Failed to send SMS",
		"description": "Failed to send SMS",
	},
	"555": {
		"name":        "Block IP address",
		"description": "Your IP address has been blocked",
	},
	"560": {
		"name":        "Mobile number",
		"description": "Mobile number is invalid",
	},
	"570": {
		"name":        "Basic SMS text ID",
		"description": "basic SMS text ID is invalid",
	},
	"580": {
		"name":        "SMS title",
		"description": "SMS title is invalid",
	},
	"585": {
		"name":        "Basic SMS text",
		"description": "Basic SMS text is invalid",
	},
	"590": {
		"name":        "Original URL",
		"description": "Original URL is invalid",
	},
	"591": {
		"name":        "SMS text is disabled",
		"description": "SMS text is disabled",
	},
	"592": {
		"name":        "Time is disabled",
		"description": "Time is disabled.",
	},
	"593": {
		"name":        "SMS",
		"description": "Wait return SMS is invalid",
	},
	"594": {
		"name":        "Return SMS",
		"description": "Return SMS is invalid",
	},
	"595": {
		"name":        "2-way SMS function is invalid",
		"description": "2-way SMS is disabled",
	},
	"573": {
		"name":        "Status",
		"description": "Status is invalid",
	},
	"574": {
		"name":        "SMS ID",
		"description": "SMS ID is invalid",
	},
	"586": {
		"name":        "SMS text 2 (forSDP)",
		"description": "SMS text 2 is invalid",
	},
	"587": {
		"name":        "SMS ID is duplicate",
		"description": "SMS ID is not unique",
	},
	"575": {
		"name":        "Docomo",
		"description": "Docomo is invalid",
	},
	"576": {
		"name":        "au",
		"description": "au is invalid",
	},
	"577": {
		"name":        "Soft Bank",
		"description": "Soft Bank is invalid",
	},
	"579": {
		"name":        "Gateway",
		"description": "Gateway is invalid",
	},
	"562": {
		"name":        "Send date & time",
		"description": "Send date & time is invalid.",
	},
	"599": {
		"name":        "Resending is disabled",
		"description": "Resending is disabled.",
	},
	"571": {
		"name":        "Sending attempts is invalid",
		"description": "Sending attempts is invalid.",
	},
	"572": {
		"name":        "Resending interval is invalid",
		"description": "Resending interval is invalid.",
	},
	"564": {
		"name":        "Reply ID is invalid",
		"description": "Reply ID is invalid.",
	},
	"566": {
		"name":        "Reply ID doesn't exist",
		"description": "Reply ID doesn't exist.",
	},
	"588": {
		"name":        "SMS text for Docomo",
		"description": "SMS text for Docomo is invalid",
	},
	"589": {
		"name":        "SMS text for Softbank",
		"description": "SMS text for Softbank is invalid",
	},
	"598": {
		"name":        "Docomo title",
		"description": "Docomo SMS title is invalid",
	},
	"568": {
		"name":        "AU title",
		"description": "AU SMS title is invalid",
	},
	"569": {
		"name":        "Softbank title",
		"description": "Softbank SMS title is invalid",
	},
	"600": {
		"name":        "Keep chatid",
		"description": "Keep chatid is invalid",
	},
	"601": {
		"name":        "SMS title function is disabled",
		"description": "SMS title function is disabled",
	},
	"602": {
		"name":        "SIM title",
		"description": "SIM SMS title is invalid.",
	},
	"603": {
		"name":        "Reminder delay",
		"description": "Reminder delay is invalid",
	},
	"604": {
		"name":        "Reminder text",
		"description": "Reminder text is invalid",
	},
	"605": {
		"name":        "Type",
		"description": "Invalid type",
	},
	"606": {
		"name":        "3rd party API is disabled",
		"description": "3rd party API is disabled",
	},
	"608": {
		"name":        "Registration date",
		"description": "Invalid Registration date",
	},
	"610": {
		"name":        "HLR is disabled",
		"description": "HLR is disabled",
	},
	"617": {
		"name":        "Memo is disabled",
		"description": "Memo is disabled",
	},
	"612": {
		"name":        "Original URL2",
		"description": "612 Original URL 2 is invalid",
	},
	"613": {
		"name":        "Original URL3",
		"description": "613 Original URL 3 is invalid",
	},
	"614": {
		"name":        "Original URL4",
		"description": "614 Original URL 4 is invalid",
	},
	"615": {
		"name":        "JSON format is incorrect",
		"description": "615 Incorrect JSON format",
	},
	"405": {
		"name":        "Method",
		"description": "405 Method not allowed",
	},
	"618": {
		"name":        "Basic SMS text is long for Docomo API",
		"description": "618 SMS text is long for Docomo API",
	},
	"619": {
		"name":        "Basic SMS text is long for SDP",
		"description": "619 SMS text is long for SDP",
	},
	"620": {
		"name":        "Basic SMS text is long for Softbank API",
		"description": "620 SMS text is long for Softbank API",
	},
	"621": {
		"name":        "Reminder text is long for Docomo API",
		"description": "Reminder text is long for Docomo API",
	},
	"622": {
		"name":        "Reminder text is long fo SDP",
		"description": "Reminder text is long for SDP",
	},
	"623": {
		"name":        "Reminder text is long for Softbank API",
		"description": "Reminder text is long for Softbank API",
	},
	"624": {
		"name":        "Duplicated SMSID",
		"description": "Duplicated SMSID",
	},
	"631": {
		"name":        "Resending parameters can not be changed",
		"description": "Resending parameters can not be changed",
	},
	"632": {
		"name":        "Rakuten title is invalid",
		"description": "Rakuten title is invalid",
	},
	"633": {
		"name":        "Rakuten text is invalid",
		"description": "Rakuten text is invalid",
	},
	"634": {
		"name":        "Main text is too long for Rakuten",
		"description": "Main text is too long for Rakuten",
	},
	"635": {
		"name":        "Reminder text is too long for Rakuten",
		"description": "Reminder text is too long for Rakuten",
	},
	"636": {
		"name":        "Rakuten is invalid",
		"description": "Rakuten is invalid",
	},
	"637": {
		"name":        "RCS is disabled",
		"description": "RCS is disabled",
	},
	"639": {
		"name":        "Original URL code is disabled",
		"description": "639 Original URL code is disabled",
	},
	"640": {
		"name":        "Original URL code",
		"description": "640 Original URL code is invalid",
	},
	"641": {
		"name":        "Original URL code 2",
		"description": "641 Original URL code 2 is invalid",
	},
	"642": {
		"name":        "Original URL code 3",
		"description": "642 Original URL code 3 is invalid",
	},
	"643": {
		"name":        "Original URL code 4",
		"description": "643 Original URL code 4 is invalid",
	},
}

var translationMap = map[string]string{
	"失敗":              "Failure",
	"送信済":             "Sent",
	"送信中":             "Sending",
	"リダイレクト":         "Redirect from shortened URL to target web sites",
	"未送信":             "Not send yet",
	"403":             "Mobile numbers are in the Black list",
	"au 圏外・OFF":       "Error status from SDP",
	"au 自 NW 障害":      "Error status from SDP",
	"au 他 NW 障害":      "Error status from SDP",
	"au 未払":           "Error status from SDP",
	"処理中":             "Delivery Pending",
	"au その他失敗":        "Error status for SDP",
	"ドコモその他失敗":       "Error status for Docomo API",
	"ドコモ圏外・OFF":      "Error status for Docomo API",
	"ドコモ自 NW 障害":     "Error status for Docomo API",
	"ドコモ他 NW 障害":     "Error status for Docomo API",
	"無効":              "Invalid records",
	"ソフトバンクその他失敗":    "Error status for Softbank API",
	"ソフトバンク圏外・OFF":   "Error status for Softbank API",
	"ソフトバンク自 NW 障害":  "Error status for Softbank API",
	"ソフトバンク他 NW 障害":  "Error status for Softbank API",
	"ソフトバンク受信拒否":     "Error status for Softbank API",
	"送信不可":            "Error status for HLR API",
	"送信前":             "SMS is scheduled, but not sent yet",
	"一時停止":            "CSV file with SMS was paused",
	"SMS 送信許可時間外":     "SMS was sent out of the allowed time",
	"双方向 SMS 利用不可":    "No resourse for Reply",
	"ドコモその他不明":       "Error status for Docomo API",
	"重複不可":            "Duplicated sms",
	"送信元重複":           "Keep chatID error",
	"ドコモ一部受信":        "Error status for Docomo API",
	"-":               "他人接続判定使用なし",
	"キャンセル済":          "Cancelled SMS",
	"送信上限":            "Ovelimitted SMS",
	"NG ワード":         "Filtered SMS by stop words",
	"配信停止希望":          "opted out",
	"長文一部受信":          "Partial delivered via Gate",
	"承認待ち":            "Wait for confirmation",
	"Rakuten その他失敗":   "Error status for Rakuten",
	"Rakuten 他 NW 障害": "Error status for Rakuten",
	"Rakuten 一部受信":    "Error status for Rakuten",
	"Rakuten 圏外・OFF":  "Error status for Rakuten",
}

var getResultsMapper = map[string]map[string]string{
	"401": {
		"name":        "Authentication error",
		"description": "Authorization Required",
	},
	"414": {
		"name":        "Failed to send request",
		"description": "URL is longer than 8,190 bytes",
	},
	"503": {
		"name":        "Temporarily unavailable",
		"description": "User reached max limit on requests/sec from the single IP address",
	},
	"502": {
		"name":        "Bad gateway",
		"description": "SMS-CONSOLE completely stopped",
	},
	"555": {
		"name":        "Block IP address",
		"description": "Your IP adress has been blocked",
	},
	"666": {
		"name":        "Prior to block IP address",
		"description": "Prior to block IP address",
	},
	"574": {
		"name":        "SMS ID",
		"description": "SMS ID is invalid",
	},
	"587": {
		"name":        "Not unique SMS ID",
		"description": "SMS ID is not unique",
	},
	"514": {
		"name":        "SMS doesn't exist",
		"description": "SMS does not exist",
	},
}
