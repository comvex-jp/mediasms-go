package mediasms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceMessageBodyURLs(t *testing.T) {
	messageBodyA := "http://google.com to https://github.com to http://stackoverflow.com and http://google.com/mail"
	allURLsA := []string{"http://google.com", "https://github.com", "http://stackoverflow.com", "http://google.com/mail"}

	newMessageBodyA := ReplaceMessageBodyURLs(messageBodyA, allURLsA)
	assert.Equal(t, "{URL} to {URL2} to {URL3} and {URL4}", newMessageBodyA, "they should be equal")
	
	messageBodyB := "google.com google.com/mail digima.com digima.com/らりるれる"
	allURLsB := []string{"google.com", "google.com/mail", "digima.com", "digima.com/らりるれる"}
	
	newMessageBodyB := ReplaceMessageBodyURLs(messageBodyB, allURLsB)
	assert.Equal(t, "{URL} {URL2} {URL3} {URL4}", newMessageBodyB, "they should be equal")
}
