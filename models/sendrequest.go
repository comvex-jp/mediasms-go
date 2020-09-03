package models

// SendRequest struct
type SendRequest struct {
	Smsid         string `json:"smsid"`
	Smstitle      string `json:"smstitle"`
	Smstext       string `json:"smstext"`
	Mobilenumber  string `json:"mobilenumber"`
	Originalurl   string `json:"originalurl"`
	Originalurl2  string `json:"originalurl2"`
	Originalurl3  string `json:"originalurl3"`
	Originalurl4  string `json:"originalurl4"`
	Status        string `json:"status"`
	Returnsms     string `json:"returnsms"`
	Waitreturnsms string `json:"waitreturnsms"`
	Type          string `json:"type"`
}
