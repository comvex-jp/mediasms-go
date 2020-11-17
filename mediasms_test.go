package mediasms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceMessageBodyURLs(t *testing.T) {
	messageBody := "http://google.com to https://github.com to http://stackoverflow.com and http://google.com"
	allURLs := []string{"http://google.com", "https://github.com", "http://stackoverflow.com", "http://google.com"}

	newMessageBody := ReplaceMessageBodyURLs(messageBody, allURLs)
	assert.Equal(t, "{URL} to {URL2} to {URL3} and {URL4}", newMessageBody, "they should be equal")
}
