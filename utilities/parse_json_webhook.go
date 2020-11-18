package utilities

import (
	"encoding/json"
	"regexp"
	"strconv"

	"github.com/comvex-jp/mediasms-go/translations"

	"github.com/comvex-jp/mediasms-go/models"
)

// ParseJSONWebHook returns a WebHook struct
func ParseJSONWebHook(body []byte) (models.WebHook, error) {
	var w models.WebHook

	err := json.Unmarshal(body, &w)

	if err != nil {
		return w, err
	}

	var yearMonthRegex = regexp.MustCompile("[年月]")

	s := yearMonthRegex.ReplaceAllString(w.ReturnSMSDatetime, `-`)

	var dayRegEx = regexp.MustCompile("[日]")

	t := dayRegEx.ReplaceAllString(s, "")

	w.ReturnSMSDatetime = t

	status, err := strconv.Atoi(w.Status)

	if err != nil {
		return w, err
	}

	w.Status = translations.TranslationMap[status]

	return w, nil
}
