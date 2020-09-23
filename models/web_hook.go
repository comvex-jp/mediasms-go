package models

// WebHook from media4u
type WebHook struct {
	MobileNumber      string `json:"mobilenumber"`
	Status            string `json:"status"`
	SMSID             string `json:"smsid"`
	SMSText           string `json:"smstext"`
	Clicks            string `json:"clicks"`
	Clicks2           string `json:"clicks2"`
	Clicks3           string `json:"clicks3"`
	Clicks4           string `json:"clicks4"`
	OriginalURL       string `json:"originalurl"`
	OriginalURL2      string `json:"originalurl2"`
	OriginalURL3      string `json:"originalurl3"`
	OriginalURL4      string `json:"originalurl4"`
	Type              string `json:"type"`
	ReturnSMS         string `json:"returnsms"`
	WaitReturnSMS     string `json:"waitreturnsms"`
	ReturnSMSDatetime string `json:"returnsmsdatetime"`
	ReplyID           string `json:"replyid"`
	SenderID          string `json:"senderid"`
}
