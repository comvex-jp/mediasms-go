package models

import (
	"encoding/json"
	"strconv"
	"strings"
)

// BuildRequest struct
// TODO: Convert the return sms and wait return sms to boolean
type BuildRequest struct {
	SMSID         string
	SMSTitle      string
	SMSText       string
	MobileNumber  string
	OriginalURL   string
	OriginalURL2  string
	OriginalURL3  string
	OriginalURL4  string
	Status        string
	ReturnSMS     string
	WaitReturnSMS string
	Type          string
	Au            AuConfig
	Docomo        DocomoConfig
	Softbank      SoftbankConfig
	Gateway       GatewayConfig
	Rakuten       RakutenConfig
	Sim           SimConfig
}

// NewBuildRequest returns ...
func NewBuildRequest() BuildRequest {
	return BuildRequest{
		Au:       AuConfig{NumberOfRetries: 1},
		Docomo:   DocomoConfig{NumberOfRetries: 1, RetryInterval: []int{0}},
		Softbank: SoftbankConfig{NumberOfRetries: 1, RetryInterval: []int{5}},
		Gateway:  GatewayConfig{NumberOfRetries: 1, RetryInterval: []int{1}},
		Rakuten:  RakutenConfig{NumberOfRetries: 1, RetryInterval: []int{2}},
		Sim:      SimConfig{NumberOfRetries: 1},
	}
}

// MarshalJSON ...
func (br BuildRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		SMSID         string `json:"smsid,omitempty"`
		SMSTitle      string `json:"smstitle,omitempty"`
		SMSText       string `json:"smstext"`
		OriginalURL   string `json:"originalurl,omitempty"`
		OriginalURL2  string `json:"originalurl2,omitempty"`
		OriginalURL3  string `json:"originalurl3,omitempty"`
		OriginalURL4  string `json:"originalurl4,omitempty"`
		MobileNumber  string `json:"mobilenumber"`
		Status        string `json:"status,omitempty"`
		ReturnSMS     string `json:"returnsms,omitempty"`
		WaitReturnSMS string `json:"waitreturnsms,omitempty"`
		Type          string `json:"type,omitempty"`
		Au            string `json:"au,omitempty"`
		Docomo        string `json:"docomo,omitempty"`
		Softbank      string `json:"softbank,omitempty"`
		Gateway       string `json:"gateway,omitempty"`
		Rakuten       string `json:"rakuten,omitempty"`
		Sim           string `json:"sim,omitempty"`
	}{
		SMSID:         br.SMSID,
		SMSTitle:      br.SMSTitle,
		SMSText:       br.SMSText,
		OriginalURL:   br.OriginalURL,
		OriginalURL2:  br.OriginalURL2,
		OriginalURL3:  br.OriginalURL3,
		OriginalURL4:  br.OriginalURL4,
		MobileNumber:  br.MobileNumber,
		Status:        br.Status,
		ReturnSMS:     br.ReturnSMS,
		WaitReturnSMS: br.WaitReturnSMS,
		Type:          br.Type,
		Au:            strconv.Itoa(br.Au.NumberOfRetries),
		Docomo:        buildRetries(br.Docomo.NumberOfRetries, br.Docomo.RetryInterval),
		Softbank:      buildRetries(br.Softbank.NumberOfRetries, br.Softbank.RetryInterval),
		Gateway:       buildRetries(br.Gateway.NumberOfRetries, br.Gateway.RetryInterval),
		Rakuten:       buildRetries(br.Rakuten.NumberOfRetries, br.Rakuten.RetryInterval),
		Sim:           strconv.Itoa(br.Sim.NumberOfRetries),
	})
}

func buildRetries(numberOfRetries int, interval []int) string {
	var sb strings.Builder

	sb.WriteString(strconv.Itoa(numberOfRetries))

	for i := 0; i < numberOfRetries; i++ {
		sb.WriteString("-")
		sb.WriteString(strconv.Itoa(interval[i]))
	}

	return sb.String()
}
