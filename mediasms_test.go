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

	messageBodyC := "http://google.com to https://github.com to http://stackoverflow.com and http://google.com/mail http://google.com"
	allURLsC := []string{"http://google.com", "https://github.com", "http://stackoverflow.com", "http://google.com/mail"}

	newMessageBodyC := ReplaceMessageBodyURLs(messageBodyC, allURLsC)
	assert.Equal(t, "{URL} to {URL2} to {URL3} and {URL4} http://google.com", newMessageBodyC, "they should be equal")

	// English space
	messageBodyD := "のdigima.com たちつてと のdigima.com/さしすせそ のhttp://www.digima.com/test"
	allURLsD := []string{"digima.com", "digima.com/さしすせそ", "http://www.digima.com/test"}

	newMessageBodyD := ReplaceMessageBodyURLs(messageBodyD, allURLsD)
	assert.Equal(t, "の{URL} たちつてと の{URL2} の{URL3}", newMessageBodyD, "they should be equal")

	// Japanese space
	messageBodyE := "のdigima.com　たちつてと　のdigima.com/さしすせそ　のhttp://www.digima.com/test"
	allURLsE := []string{"digima.com", "digima.com/さしすせそ", "http://www.digima.com/test"}

	newMessageBodyE := ReplaceMessageBodyURLs(messageBodyE, allURLsE)
	assert.Equal(t, "の{URL}　たちつてと　の{URL2}　の{URL3}", newMessageBodyE, "they should be equal")
}
