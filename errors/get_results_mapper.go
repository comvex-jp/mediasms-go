package errors

// GetResultsMapper is a hashmap of all possible results when getting the status of an sms
var GetResultsMapper = map[string]map[string]string{
	"401": {
		"name":        "authentication_error",
		"description": "Authorization Required",
	},
	"414": {
		"name":        "failed_to_send_request",
		"description": "URL is longer than 8,190 bytes",
	},
	"503": {
		"name":        "temporarily_unavailable",
		"description": "User reached max limit on requests/sec from the single IP address",
	},
	"502": {
		"name":        "bad_gateway",
		"description": "SMS-CONSOLE completely stopped",
	},
	"555": {
		"name":        "block_ip_address",
		"description": "Your IP address has been blocked",
	},
	"666": {
		"name":        "prior_to_block_ip_address",
		"description": "Prior to block IP address",
	},
	"574": {
		"name":        "sms_id",
		"description": "SMS ID is invalid",
	},
	"587": {
		"name":        "not_unique_sms_id",
		"description": "SMS ID is not unique",
	},
	"514": {
		"name":        "sms_doesn't_exist",
		"description": "SMS does not exist",
	},
}
