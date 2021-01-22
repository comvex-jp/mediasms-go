package errors

import "github.com/comvex-jp/mediasms-go/models"

// ResultsMapper is a hashmap of all possible results when sending an sms
var ResultsMapper = map[string]map[string]string{
	"200": {
		"name":        models.Success,
		"description": "Message has been succesfully sent",
	},
	"401": {
		"name":        models.AuthenticationError,
		"description": "Authorization required",
	},
	"405": {
		"name":        models.MethodNotAllowed,
		"description": "Method not allowed",
	},
	"414": {
		"name":        models.URLTooLong,
		"description": "URL is longer than 8,190 bytes",
	},
	"502": {
		"name":        models.BadGateway,
		"description": "SMS-CONSOLE completely stopped",
	},
	"503": {
		"name":        models.TemporaryUnavailable,
		"description": "User reached max limit on requests/sec from the single IP address",
	},
	"514": {
		"name":        models.SMSNotFound,
		"description": "SMS does not exist",
	},
	"550": {
		"name":        models.SendFailure,
		"description": "Failed to send SMS",
	},
	"555": {
		"name":        models.BlockedIPAddress,
		"description": "Your IP address has been blocked",
	},
	"560": {
		"name":        models.InvalidMobileNumber,
		"description": "Mobile number is invalid",
	},
	"562": {
		"name":        models.InvalidSendDateTime,
		"description": "Send date & time is invalid.",
	},
	"564": {
		"name":        models.InvalidReplyID,
		"description": "Reply ID is invalid.",
	},
	"566": {
		"name":        models.ReplyIDNotFound,
		"description": "Reply ID doesn't exist.",
	},
	"568": {
		"name":        models.InvalidAuTitle,
		"description": "AU SMS title is invalid",
	},
	"569": {
		"name":        models.InvalidSoftbankTitle,
		"description": "Softbank SMS title is invalid",
	},
	"570": {
		"name":        models.InvalidSMSTextID,
		"description": "basic SMS text ID is invalid",
	},
	"571": {
		"name":        models.InvalidSendingAttempts,
		"description": "Sending attempts is invalid.",
	},
	"572": {
		"name":        models.InvalidResendingInterval,
		"description": "Resending interval is invalid.",
	},
	"573": {
		"name":        models.InvalidStatus,
		"description": "Status is invalid",
	},
	"574": {
		"name":        models.InvalidSMSID,
		"description": "SMS ID is invalid",
	},
	"575": {
		"name":        models.InvalidDocomoTTL,
		"description": "Docomo is invalid",
	},
	"576": {
		"name":        models.InvalidAuTTL,
		"description": "au is invalid",
	},
	"577": {
		"name":        models.InvalidSoftbankTTL,
		"description": "Soft Bank is invalid",
	},
	"579": {
		"name":        models.InvalidGatewayTTL,
		"description": "Gateway is invalid",
	},
	"580": {
		"name":        models.InvalidSMSTitle,
		"description": "SMS title is invalid",
	},
	"585": {
		"name":        models.InvalidBasicSMSText,
		"description": "Basic SMS text is invalid",
	},
	"586": {
		"name":        models.InvalidSMSTextTwo,
		"description": "SMS text 2 is invalid",
	},
	"587": {
		"name":        models.SMSIDNotUnique,
		"description": "SMS ID is not unique",
	},
	"588": {
		"name":        models.InvalidDocomoSMSText,
		"description": "SMS text for Docomo is invalid",
	},
	"589": {
		"name":        models.InvalidSoftbankSMSText,
		"description": "SMS text for Softbank is invalid",
	},
	"590": {
		"name":        models.InvalidOriginalURL,
		"description": "Original URL is invalid",
	},
	"591": {
		"name":        models.DisabledSMSText,
		"description": "SMS text is disabled",
	},
	"592": {
		"name":        models.InvalidDispatchTime,
		"description": "Dispatch time is not on the allowed range.",
	},
	"593": {
		"name":        models.InvalidWaitReturnSMS,
		"description": "Wait return SMS is invalid",
	},
	"594": {
		"name":        models.InvalidReturnSMS,
		"description": "Return SMS is invalid",
	},
	"595": {
		"name":        models.InvalidTwoWaySmsFunction,
		"description": "2-way SMS is disabled",
	},
	"598": {
		"name":        models.InvalidDocomoTitle,
		"description": "Docomo SMS title is invalid",
	},
	"599": {
		"name":        models.DisabledResending,
		"description": "Resending is disabled.",
	},
	"600": {
		"name":        models.InvalidKeepChatID,
		"description": "Keep chatid is invalid",
	},
	"601": {
		"name":        models.DisabledSMSTitleFunction,
		"description": "SMS title function is disabled",
	},
	"602": {
		"name":        models.InvalidSIMTitle,
		"description": "SIM SMS title is invalid.",
	},
	"603": {
		"name":        models.InvalidReminderDelay,
		"description": "Reminder delay is invalid",
	},
	"604": {
		"name":        models.InvalidReminderText,
		"description": "Reminder text is invalid",
	},
	"605": {
		"name":        models.InvalidType,
		"description": "Invalid type",
	},
	"606": {
		"name":        models.DisabledThirdPartyAPI,
		"description": "3rd party API is disabled",
	},
	"608": {
		"name":        models.InvalidRegistrationDate,
		"description": "Invalid Registration date",
	},
	"610": {
		"name":        models.DisabledHLR,
		"description": "HLR is disabled",
	},
	"612": {
		"name":        models.InvalidOriginalURLTwo,
		"description": "612 Original URL 2 is invalid",
	},
	"613": {
		"name":        models.InvalidOriginalURLThree,
		"description": "613 Original URL 3 is invalid",
	},
	"614": {
		"name":        models.InvalidOriginalURLFour,
		"description": "614 Original URL 4 is invalid",
	},
	"615": {
		"name":        models.InvalidJSONFormat,
		"description": "615 Incorrect JSON format",
	},
	"617": {
		"name":        models.DisabledMemo,
		"description": "Memo is disabled",
	},
	"618": {
		"name":        models.DocomoBasicSMSTextTooLong,
		"description": "618 SMS text is long for Docomo API",
	},
	"619": {
		"name":        models.SDPBasicSMSTextTooLong,
		"description": "619 SMS text is long for SDP",
	},
	"620": {
		"name":        models.SoftbankSMSTextTooLong,
		"description": "620 SMS text is long for Softbank API",
	},
	"621": {
		"name":        models.DocomoReminderTextTooLong,
		"description": "Reminder text is long for Docomo API",
	},
	"622": {
		"name":        models.SDPReminderTextTooLong,
		"description": "Reminder text is long for SDP",
	},
	"623": {
		"name":        models.SoftbankReminderTextTooLong,
		"description": "Reminder text is long for Softbank API",
	},
	"624": {
		"name":        models.DuplicatedSMSId,
		"description": "Duplicated SMSID",
	},
	"631": {
		"name":        models.ResendingParametersChangeFailure,
		"description": "Resending parameters can not be changed",
	},
	"632": {
		"name":        models.InvalidRakutenTitle,
		"description": "Rakuten title is invalid",
	},
	"633": {
		"name":        models.InvalidRakutenText,
		"description": "Rakuten text is invalid",
	},
	"634": {
		"name":        models.RakutenMainTextTooLong,
		"description": "Main text is too long for Rakuten",
	},
	"635": {
		"name":        models.RakutenReminderTextTooLong,
		"description": "Reminder text is too long for Rakuten",
	},
	"636": {
		"name":        models.InvalidRakutenTTL,
		"description": "Rakuten is invalid",
	},
	"637": {
		"name":        models.DisabledRCS,
		"description": "RCS is disabled",
	},
	"639": {
		"name":        models.DisabledOriginalURLCode,
		"description": "639 Original URL code is disabled",
	},
	"640": {
		"name":        models.InvalidOriginalURLCode,
		"description": "640 Original URL code is invalid",
	},
	"641": {
		"name":        models.InvalidOriginalURLCodeTwo,
		"description": "641 Original URL code 2 is invalid",
	},
	"642": {
		"name":        models.InvalidOriginalURLCodeThree,
		"description": "642 Original URL code 3 is invalid",
	},
	"643": {
		"name":        models.InvalidOriginalURLCodeFour,
		"description": "643 Original URL code 4 is invalid",
	},
	"666": {
		"name":        models.IPAddressPriorBlock,
		"description": "Prior to block IP address",
	},
}
