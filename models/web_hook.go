package models

// WebHook from media4u
type WebHook struct {
	MobileNumber      string
	Status            string
	SMSID             string
	ReturnSMS         string
	WaitReturnSMS     string
	ReturnSMSDatetime string
	ReplyID           string
	SenderID          string
}
