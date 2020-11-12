package models

// BuildRequest struct
type BuildRequest struct {
	SMSID         string `json:"smsid"`
	SMSTitle      string `json:"smstitle"`
	SMSText       string `json:"smstext"`
	MobileNumber  string `json:"mobilenumber"`
	OriginalURL   string `json:"originalurl"`
	OriginalURL2  string `json:"originalurl2"`
	OriginalURL3  string `json:"originalurl3"`
	OriginalURL4  string `json:"originalurl4"`
	Status        string `json:"status"`
	ReturnSMS     string `json:"returnsms"`
	WaitReturnSMS string `json:"waitreturnsms"`
	Type          string `json:"type"`
	Au            string `json:"au"`
	Docomo        string `json:"docomo"`
	Softbank      string `json:"softbank"`
	Gateway       string `json:"gateway"`
	Rakuten       string `json:"rakuten"`
}
