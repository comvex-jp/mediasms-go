package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/comvex-jp/mediasms-go/helpers"
)

// BuildRequest struct
// SMSID is a unique identifier used to retrieve status after sending, If Status is set to true the SMSID is required
// Status is a nullable boolean, if the status is set as true, the status will be sent back
type BuildRequest struct {
	SMSID         string
	SMSTitle      string
	SMSText       string
	MobileNumber  string
	OriginalURL   string
	OriginalURL2  string
	OriginalURL3  string
	OriginalURL4  string
	Status        bool
	ReturnSMS     bool
	WaitReturnSMS bool
	Type          string
	Au            AuConfig
	Docomo        DocomoConfig
	Softbank      SoftbankConfig
	Gateway       GatewayConfig
	Rakuten       RakutenConfig
	Sim           SimConfig
}

// NewBuildRequest returns an instance of BuildRequest with default values
func NewBuildRequest() BuildRequest {
	return BuildRequest{
		Au:       AuConfig{NumberOfRetries: 1},
		Docomo:   DocomoConfig{NumberOfRetries: 1, RetryInterval: []int{0}},
		Softbank: SoftbankConfig{NumberOfRetries: 1, RetryInterval: []int{5}},
		Gateway:  GatewayConfig{NumberOfRetries: 1, RetryInterval: []int{1}},
		Rakuten:  RakutenConfig{NumberOfRetries: 1, RetryInterval: []int{2}},
		Sim:      SimConfig{NumberOfRetries: 1},
		Type:     TypeSms,
		Status:   true,
	}
}

// MarshalJSON overrides the custom serialization to prepare the data for the MediaSms API
func (br BuildRequest) MarshalJSON() ([]byte, error) {
	err := br.validate()

	if err != nil {
		return nil, err
	}

	return json.Marshal(&struct {
		SMSID         string `json:"smsid,omitempty"` // SMSID is required when the ReturnSMS is true
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
		Status:        buildBooleanPayload(br.Status),
		ReturnSMS:     buildBooleanPayload(br.ReturnSMS),
		WaitReturnSMS: buildBooleanPayload(br.WaitReturnSMS),
		Type:          br.Type,
		Au:            strconv.Itoa(br.Au.NumberOfRetries),
		Docomo:        buildRetriesPayload(br.Docomo.NumberOfRetries, br.Docomo.RetryInterval),
		Softbank:      buildRetriesPayload(br.Softbank.NumberOfRetries, br.Softbank.RetryInterval),
		Gateway:       buildRetriesPayload(br.Gateway.NumberOfRetries, br.Gateway.RetryInterval),
		Rakuten:       buildRetriesPayload(br.Rakuten.NumberOfRetries, br.Rakuten.RetryInterval),
		Sim:           strconv.Itoa(br.Sim.NumberOfRetries),
	})
}

func (br BuildRequest) validate() error {
	if br.ReturnSMS && br.SMSID == "" {
		return errors.New("SMSID is required on two way sms (ReturnSMS) is enabled")
	}

	if br.Status && br.SMSID == "" {
		return errors.New("SMSID is required when Status is enabled")
	}

	if len(br.Docomo.RetryInterval) != br.Docomo.NumberOfRetries {
		return errors.New("[Docomo] Invalid amount of retries interval values, the number of retries interval has to be the same as number of retries specified")
	}

	if len(br.Softbank.RetryInterval) != br.Softbank.NumberOfRetries {
		return errors.New("[Softbank] Invalid amount of retries interval values, the number of retries interval has to be the same as number of retries specified")
	}

	if len(br.Gateway.RetryInterval) != br.Gateway.NumberOfRetries {
		return errors.New("[Gateway] Invalid amount of retries interval values, the number of retries interval has to be the same as number of retries specified")
	}

	if len(br.Rakuten.RetryInterval) != br.Rakuten.NumberOfRetries {
		return errors.New("[Rakuten] Invalid amount of retries interval values, the number of retries interval has to be the same as number of retries specified")
	}

	if br.Sim.NumberOfRetries < 1 && br.Sim.NumberOfRetries > 5 {
		return errors.New("[Sim] Invalid amount number of retries")
	}

	if br.Au.NumberOfRetries < 1 && br.Au.NumberOfRetries > 5 {
		return errors.New("[Au] Invalid amount number of retries")
	}

	if !helpers.Contains(availableSmsTypes(), br.Type) {
		return fmt.Errorf("Invalid sms type. The type must be one of the: %s", strings.Join(availableSmsTypes(), ","))
	}

	// TODO: Validate URL lenght

	return nil
}

func buildBooleanPayload(val bool) string {
	if val {
		return "1"
	}

	return "0"
}

func buildRetriesPayload(numberOfRetries int, interval []int) string {
	var sb strings.Builder

	sb.WriteString(strconv.Itoa(numberOfRetries))

	for i := 0; i < numberOfRetries; i++ {
		sb.WriteString("-")
		sb.WriteString(strconv.Itoa(interval[i]))
	}

	return sb.String()
}
