package models

import (
	"encoding/json"
	"testing"
)

func TestBuildRequestMarshalJson(t *testing.T) {
	br := NewBuildRequest()

	br.MobileNumber = "07012345678"
	br.SMSText = "Lorem Ipsum"

	mbr, err := json.Marshal(br)

	if err != nil {
		t.Errorf("Validation Failure %s", err.Error())
	}

	expectedJSONString := `{"smstext":"Lorem Ipsum","mobilenumber":"07012345678","returnsms":"0","waitreturnsms":"0","type":"sms","au":"1","docomo":"1-0","softbank":"1-5","gateway":"1-1","rakuten":"1-2","sim":"1"}`

	if string(mbr) != expectedJSONString {
		t.Errorf("Unexpected Serialization Value %s", string(mbr))
	}
}
