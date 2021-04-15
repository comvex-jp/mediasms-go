package mediasms

import (
	"encoding/json"
	"testing"

	"github.com/comvex-jp/mediasms-go/models"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestReplaceMessageBodyURLs(t *testing.T) {
	messageBodyA := "http://google.com to https://github.com to http://stackoverflow.com and http://google.com/mail"
	allURLsA := []string{"http://google.com", "https://github.com", "http://stackoverflow.com", "http://google.com/mail"}

	newMessageBodyA := ReplaceMessageBodyURLs(messageBodyA, allURLsA)
	assert.Equal(t, "{URL} to {URL2} to {URL3} and {URL4}", newMessageBodyA, "they should be equal")

	messageBodyB := "google.com google.com/mail digima.com digima.com/らりるれる"
	allURLsB := []string{"google.com", "google.com/mail", "digima.com", "digima.com/らりるれる"}

	newMessageBodyB := ReplaceMessageBodyURLs(messageBodyB, allURLsB)
	assert.Equal(t, "{URL} {URL2} {URL3} {URL4}", newMessageBodyB, "they should be equal")

	messageBodyC := "http://google.com to https://github.com to http://stackoverflow.com and http://google.com/mail http://google.com"
	allURLsC := []string{"http://google.com", "https://github.com", "http://stackoverflow.com", "http://google.com/mail"}

	newMessageBodyC := ReplaceMessageBodyURLs(messageBodyC, allURLsC)
	assert.Equal(t, "{URL} to {URL2} to {URL3} and {URL4} {URL}", newMessageBodyC, "they should be equal")

	// English space
	messageBodyD := "のdigima.com たちつてと のdigima.com/さしすせそ のhttp://www.digima.com/test"
	allURLsD := []string{"digima.com", "digima.com/さしすせそ", "http://www.digima.com/test"}

	newMessageBodyD := ReplaceMessageBodyURLs(messageBodyD, allURLsD)
	assert.Equal(t, "の{URL} たちつてと の{URL2} の{URL3}", newMessageBodyD, "they should be equal")

	// Japanese space
	messageBodyE := "のdigima.com　たちつてと　のdigima.com/さしすせそ　のhttp://www.digima.com/test"
	allURLsE := []string{"digima.com", "digima.com/さしすせそ", "http://www.digima.com/test"}

	newMessageBodyE := ReplaceMessageBodyURLs(messageBodyE, allURLsE)
	assert.Equal(t, "の{URL}　たちつてと　の{URL2}　の{URL3}", newMessageBodyE, "they should be equal")
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

