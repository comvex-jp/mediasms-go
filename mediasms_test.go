package mediasms

import (
	"encoding/json"
	"testing"

	"github.com/comvex-jp/mediasms-go/models"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestReplaceMessageBodyURLs(t *testing.T) {
	messageBody := "http://google.com to https://github.com to http://stackoverflow.com and http://google.com"
	allURLs := []string{"http://google.com", "https://github.com", "http://stackoverflow.com", "http://google.com"}

	newMessageBody := ReplaceMessageBodyURLs(messageBody, allURLs)
	assert.Equal(t, "{URL} to {URL2} to {URL3} and {URL4}", newMessageBody, "they should be equal")
}

func TestMediaSMSResponder(t *testing.T) {
	client := Client{
		Username: "user_1",
		Password: "password",
		Prefix:   "PREFX",
		EnableMock: true,
	}

	response := map[string]interface{}{
		"result": "200",
	}

	request := models.BuildRequest{
		SMSID:         "",
		SMSTitle:      "",
		SMSText:       "Contact Message {URL} {URL2}",
		MobileNumber:  "07011112222",
		OriginalURL:   "http://google.com",
		OriginalURL2:  "http://yahoo.com",
		OriginalURL3:  "",
		OriginalURL4:  "",
		Status:        true,
		ReturnSMS:     true,
		WaitReturnSMS: true,
		Type:          "sms",
		Au:            models.AuConfig{NumberOfAttempts: 1},
		Docomo: models.DocomoConfig{
			NumberOfAttempts: 1,
			RetryInterval:    []int{1},
		},
		Softbank: models.SoftbankConfig{
			NumberOfAttempts: 1,
			RetryInterval:    []int{5},
		},
		Gateway: models.GatewayConfig{
			NumberOfAttempts: 1,
			RetryInterval:    []int{1},
		},
		Rakuten: models.RakutenConfig{
			NumberOfAttempts: 1,
			RetryInterval:    []int{2}},
		Sim: models.SimConfig{NumberOfAttempts: 1},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		assert.Fail(t, "Response marshalling failed: ", err)
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock responders
	httpmock.RegisterResponder("POST", SMSURL+"api/",
		httpmock.NewStringResponder(200, string(jsonResponse)))

	mResponse, err := client.Send("M1", request)
	if err != nil {
		assert.Fail(t, "Error: ", err)
	}

	assert.Equal(t, "200", mResponse.StatusCode)
	assert.Equal(t, "success", mResponse.Name)
	assert.Equal(t, "Message has been succesfully sent", mResponse.Description)
}

