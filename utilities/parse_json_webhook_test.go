package utilities

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var reqJSON = map[string]string{
	"mobilenumber":      "11122223333",
	"status":            "送信済",
	"smsid":             "1234567890",
	"smstext":           "test",
	"clicks":            "1",
	"clicks2":           "0",
	"clicks3":           "0",
	"clicks4":           "0",
	"originalurl":       "http://google.com",
	"originalurl2":      "",
	"originalurl3":      "",
	"originalurl4":      "",
	"type":              "sms",
	"returnsms":         "1",
	"waitreturnsms":     "1",
	"returnsmsdatetime": "2020年09月01日 10:00:00",
	"replyid":           "1",
	"senderid":          "222211110000"}

func TestParseJSONWebHook(t *testing.T) {
	request, _ := json.Marshal(reqJSON)
	w := ParseJSONWebHook(request)

	assert.Equal(t, "11122223333", w.MobileNumber, "they should be equal")
	assert.Equal(t, "http://google.com", w.OriginalURL, "they should be equal")
	assert.Equal(t, "2020-09-01 10:00:00", w.ReturnSMSDatetime, "they should be equal")
	assert.Equal(t, "222211110000", w.SenderID, "they should be equal")
}
