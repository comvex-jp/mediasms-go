package models

// WebHook from media4u
type WebHook struct {
	MobileNumber      string `json:"mobilenumber"`
	Status            string `json:"status"`
	SMSID             string `json:"smsid"`
	ReturnSMS         string `json:"returnsms"`
	WaitReturnSMS     string `json:"waitreturnsms"`
	ReturnSMSDatetime string `json:"returnsmsdatetime"`
	ReplyID           string `json:"replyid"`
	SenderID          string `json:"senderid"`
}
