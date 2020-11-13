package models

const (
	// TypeSms TODO
	TypeSms = "sms"
	// TypeCheck TODO
	TypeCheck = "check"
	// TypeCheckSms TODO
	TypeCheckSms = "checksms"
	// TypeHlr TODO
	TypeHlr = "hlr"
	// TypeRcs TODO
	TypeRcs = "rcs"
	// TypeRcsSms TODO
	TypeRcsSms = "rcssms"
)

// AvailableSmsTypes contains the list of the sms types
func availableSmsTypes() []string {
	return []string{TypeSms, TypeCheck, TypeCheckSms, TypeHlr, TypeRcs, TypeRcsSms}
}
