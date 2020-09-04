package errors

// GetResultsMapper is a hashmap of all possible results when getting the status of an sms
var GetResultsMapper = map[string]map[string]string{
	"401": {
		"name":        "Authentication error",
		"description": "Authorization Required",
	},
	"414": {
		"name":        "Failed to send request",
		"description": "URL is longer than 8,190 bytes",
	},
	"503": {
		"name":        "Temporarily unavailable",
		"description": "User reached max limit on requests/sec from the single IP address",
	},
	"502": {
		"name":        "Bad gateway",
		"description": "SMS-CONSOLE completely stopped",
	},
	"555": {
		"name":        "Block IP address",
		"description": "Your IP adress has been blocked",
	},
	"666": {
		"name":        "Prior to block IP address",
		"description": "Prior to block IP address",
	},
	"574": {
		"name":        "SMS ID",
		"description": "SMS ID is invalid",
	},
	"587": {
		"name":        "Not unique SMS ID",
		"description": "SMS ID is not unique",
	},
	"514": {
		"name":        "SMS doesn't exist",
		"description": "SMS does not exist",
	},
}
