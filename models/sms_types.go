package models

// SmsTypeAlias hide the real type of the enum
type SmsTypeAlias = string

type list struct {
	// TypeSms is used to send short messages to given phone numbers
	Sms SmsTypeAlias
	// TypeCheck TODO
	Check SmsTypeAlias
	// TypeCheckSms TODO
	CheckSms SmsTypeAlias
	// TypeHlr TODO
	Hlr SmsTypeAlias
	// TypeRcs TODO
	Rcs SmsTypeAlias
	// TypeRcsSms TODO
	RcsSms SmsTypeAlias
}

// SmsTypes are the available sms types
var SmsTypes = &list{
	Sms:      "sms",
	Check:    "check",
	CheckSms: "checksms",
	Hlr:      "hlr",
	Rcs:      "rcs",
	RcsSms:   "rcssms",
}
