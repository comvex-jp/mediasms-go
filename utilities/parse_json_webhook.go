package utilities

import (
	"encoding/json"
	"regexp"

	"github.com/comvex-jp/mediasms-go/translations"

	"github.com/comvex-jp/mediasms-go/models"
)

// ParseJSONWebHook returns a WebHook struct
func ParseJSONWebHook(body []byte) models.WebHook {
	var w models.WebHook

	json.Unmarshal(body, &w)

	var yearMonthRegex = regexp.MustCompile("[年月]")
	s := yearMonthRegex.ReplaceAllString(w.ReturnSMSDatetime, `-`)

	var dayRegEx = regexp.MustCompile("[日]")
	t := dayRegEx.ReplaceAllString(s, "")
	w.ReturnSMSDatetime = t

	w.Status = translations.TranslationMap[w.Status]

	return w
}
