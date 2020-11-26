package models

// AuConfig represents the au configuration when building the request
// where you can set the number of retries, the interval of the retry
// is automatically definited by MediaSms (Media4u)
type AuConfig struct {
	// NumberOfAttempts between 1 and 5
	NumberOfAttempts int
}

// DocomoConfig represents the au configuration when building the request
// where you can set the number of retries, if the retry is set, the appropriated
// retry interval for each attempt has to be set
type DocomoConfig struct {
	// NumberOfAttempts between 1 and 5
	NumberOfAttempts int
	// Retry interval between 1 and 9999
	RetryInterval []int
}

// SoftbankConfig represents the softbank configuration when building the request
// where you can set the number of retries, if the retry is set, the appropriated
// retry interval for each attempt has to be set
type SoftbankConfig struct {
	// NumberOfAttempts between 1 and 5
	NumberOfAttempts int
	// Retry interval between 5 and 60
	RetryInterval []int
}

// GatewayConfig represents the gateway configuration when building the request
// where you can set the number of retries, if the retry is set, the appropriated
// retry interval for each attempt has to be set
type GatewayConfig struct {
	// NumberOfAttempts between 1 and 5
	NumberOfAttempts int
	// Retry interval between 1 and 999
	RetryInterval []int
}

// SimConfig represents the au configuration when building the request
// where you can set the number of retries, the interval of the retry
// is automatically definited by MediaSms (Media4u)
type SimConfig struct {
	// NumberOfAttempts between 1 and 5
	NumberOfAttempts int
}

// RakutenConfig represents the gateway configuration when building the request
// where you can set the number of retries, if the retry is set, the appropriated
// retry interval for each attempt has to be set
type RakutenConfig struct {
	// NumberOfAttempts between 1 and 5
	NumberOfAttempts int
	// Retry interval between 2 and 4320
	RetryInterval []int
}
