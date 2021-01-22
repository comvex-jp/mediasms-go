package errors

// ResultsMapper is a hashmap of all possible results when sending an sms
var ResultsMapper = map[string]map[string]string{
	"200": {
		"name":        Success,
		"description": "Message has been succesfully sent",
	},
	"401": {
		"name":        AuthenticationError,
		"description": "Authorization required",
	},
	"405": {
		"name":        MethodNotAllowed,
		"description": "Method not allowed",
	},
	"414": {
		"name":        URLTooLong,
		"description": "URL is longer than 8,190 bytes",
	},
	"502": {
		"name":        BadGateway,
		"description": "SMS-CONSOLE completely stopped",
	},
	"503": {
		"name":        TemporaryUnavailable,
		"description": "User reached max limit on requests/sec from the single IP address",
	},
	"514": {
		"name":        SMSNotFound,
		"description": "SMS does not exist",
	},
	"550": {
		"name":        SendFailure,
		"description": "Failed to send SMS",
	},
	"555": {
		"name":        BlockedIPAddress,
		"description": "Your IP address has been blocked",
	},
	"560": {
		"name":        InvalidMobileNumber,
		"description": "Mobile number is invalid",
	},
	"562": {
		"name":        InvalidSendDateTime,
		"description": "Send date & time is invalid.",
	},
	"564": {
		"name":        InvalidReplyID,
		"description": "Reply ID is invalid.",
	},
	"566": {
		"name":        ReplyIDNotFound,
		"description": "Reply ID doesn't exist.",
	},
	"568": {
		"name":        InvalidAuTitle,
		"description": "AU SMS title is invalid",
	},
	"569": {
		"name":        InvalidSoftbankTitle,
		"description": "Softbank SMS title is invalid",
	},
	"570": {
		"name":        InvalidSMSTextID,
		"description": "basic SMS text ID is invalid",
	},
	"571": {
		"name":        InvalidSendingAttempts,
		"description": "Sending attempts is invalid.",
	},
	"572": {
		"name":        InvalidResendingInterval,
		"description": "Resending interval is invalid.",
	},
	"573": {
		"name":        InvalidStatus,
		"description": "Status is invalid",
	},
	"574": {
		"name":        InvalidSMSID,
		"description": "SMS ID is invalid",
	},
	"575": {
		"name":        InvalidDocomoTTL,
		"description": "Docomo is invalid",
	},
	"576": {
		"name":        InvalidAuTTL,
		"description": "au is invalid",
	},
	"577": {
		"name":        InvalidSoftbankTTL,
		"description": "Soft Bank is invalid",
	},
	"579": {
		"name":        InvalidGatewayTTL,
		"description": "Gateway is invalid",
	},
	"580": {
		"name":        InvalidSMSTitle,
		"description": "SMS title is invalid",
	},
	"585": {
		"name":        InvalidBasicSMSText,
		"description": "Basic SMS text is invalid",
	},
	"586": {
		"name":        InvalidSMSTextTwo,
		"description": "SMS text 2 is invalid",
	},
	"587": {
		"name":        SMSIDNotUnique,
		"description": "SMS ID is not unique",
	},
	"588": {
		"name":        InvalidDocomoSMSText,
		"description": "SMS text for Docomo is invalid",
	},
	"589": {
		"name":        InvalidSoftbankSMSText,
		"description": "SMS text for Softbank is invalid",
	},
	"590": {
		"name":        InvalidOriginalURL,
		"description": "Original URL is invalid",
	},
	"591": {
		"name":        DisabledSMSText,
		"description": "SMS text is disabled",
	},
	"592": {
		"name":        InvalidDispatchTime,
		"description": "Dispatch time is not on the allowed range.",
	},
	"593": {
		"name":        InvalidWaitReturnSMS,
		"description": "Wait return SMS is invalid",
	},
	"594": {
		"name":        InvalidReturnSMS,
		"description": "Return SMS is invalid",
	},
	"595": {
		"name":        InvalidTwoWaySmsFunction,
		"description": "2-way SMS is disabled",
	},
	"598": {
		"name":        InvalidDocomoTitle,
		"description": "Docomo SMS title is invalid",
	},
	"599": {
		"name":        DisabledResending,
		"description": "Resending is disabled.",
	},
	"600": {
		"name":        InvalidKeepChatID,
		"description": "Keep chatid is invalid",
	},
	"601": {
		"name":        DisabledSMSTitleFunction,
		"description": "SMS title function is disabled",
	},
	"602": {
		"name":        InvalidSIMTitle,
		"description": "SIM SMS title is invalid.",
	},
	"603": {
		"name":        InvalidReminderDelay,
		"description": "Reminder delay is invalid",
	},
	"604": {
		"name":        InvalidReminderText,
		"description": "Reminder text is invalid",
	},
	"605": {
		"name":        InvalidType,
		"description": "Invalid type",
	},
	"606": {
		"name":        DisabledThirdPartyAPI,
		"description": "3rd party API is disabled",
	},
	"608": {
		"name":        InvalidRegistrationDate,
		"description": "Invalid Registration date",
	},
	"610": {
		"name":        DisabledHLR,
		"description": "HLR is disabled",
	},
	"612": {
		"name":        InvalidOriginalURLTwo,
		"description": "612 Original URL 2 is invalid",
	},
	"613": {
		"name":        InvalidOriginalURLThree,
		"description": "613 Original URL 3 is invalid",
	},
	"614": {
		"name":        InvalidOriginalURLFour,
		"description": "614 Original URL 4 is invalid",
	},
	"615": {
		"name":        InvalidJSONFormat,
		"description": "615 Incorrect JSON format",
	},
	"617": {
		"name":        DisabledMemo,
		"description": "Memo is disabled",
	},
	"618": {
		"name":        DocomoBasicSMSTextTooLong,
		"description": "618 SMS text is long for Docomo API",
	},
	"619": {
		"name":        SDPBasicSMSTextTooLong,
		"description": "619 SMS text is long for SDP",
	},
	"620": {
		"name":        SoftbankSMSTextTooLong,
		"description": "620 SMS text is long for Softbank API",
	},
	"621": {
		"name":        DocomoReminderTextTooLong,
		"description": "Reminder text is long for Docomo API",
	},
	"622": {
		"name":        SDPReminderTextTooLong,
		"description": "Reminder text is long for SDP",
	},
	"623": {
		"name":        SoftbankReminderTextTooLong,
		"description": "Reminder text is long for Softbank API",
	},
	"624": {
		"name":        DuplicatedSMSId,
		"description": "Duplicated SMSID",
	},
	"631": {
		"name":        ResendingParametersChangeFailure,
		"description": "Resending parameters can not be changed",
	},
	"632": {
		"name":        InvalidRakutenTitle,
		"description": "Rakuten title is invalid",
	},
	"633": {
		"name":        InvalidRakutenText,
		"description": "Rakuten text is invalid",
	},
	"634": {
		"name":        RakutenMainTextTooLong,
		"description": "Main text is too long for Rakuten",
	},
	"635": {
		"name":        RakutenReminderTextTooLong,
		"description": "Reminder text is too long for Rakuten",
	},
	"636": {
		"name":        InvalidRakutenTTL,
		"description": "Rakuten is invalid",
	},
	"637": {
		"name":        DisabledRCS,
		"description": "RCS is disabled",
	},
	"639": {
		"name":        DisabledOriginalURLCode,
		"description": "639 Original URL code is disabled",
	},
	"640": {
		"name":        InvalidOriginalURLCode,
		"description": "640 Original URL code is invalid",
	},
	"641": {
		"name":        InvalidOriginalURLCodeTwo,
		"description": "641 Original URL code 2 is invalid",
	},
	"642": {
		"name":        InvalidOriginalURLCodeThree,
		"description": "642 Original URL code 3 is invalid",
	},
	"643": {
		"name":        InvalidOriginalURLCodeFour,
		"description": "643 Original URL code 4 is invalid",
	},
	"666": {
		"name":        IPAddressPriorBlock,
		"description": "Prior to block IP address",
	},
}
