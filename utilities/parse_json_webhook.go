package utilities

import (
	"encoding/json"

	"github.com/comvex-jp/mediasms-go/models"
)

// ParseJSONWebHook returns a WebHook struct
func ParseJSONWebHook(body []byte) models.WebHook {
	var w models.WebHook
	json.Unmarshal(body, &w)
	return w
}
