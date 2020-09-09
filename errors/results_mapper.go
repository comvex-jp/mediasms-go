package errors

// ResultsMapper is a hashmap of all possible results when sending an sms
var ResultsMapper = map[string]map[string]string{
	"200": {
		"name":        "success_code",
		"description": "Sent SMS",
	},
	"401": {
		"name":        "authentication_error",
		"description": "Authorization Required",
	},
	"405": {
		"name":        "method",
		"description": "405 Method not allowed",
	},
	"414": {
		"name":        "failed_to_send_request",
		"description": "URL is longer than 8,190 byte",
	},
	"502": {
		"name":        "bad_gateway",
		"description": "SMS-CONSOLE completely stopped",
	},
	"503": {
		"name":        "temporarily_unavailable",
		"description": "User reached max limit on requests/sec from the single IP address",
	},
	"514": {
		"name":        "sms_doesnt_exist",
		"description": "SMS does not exist",
	},
	"550": {
		"name":        "failed_to_send_sms",
		"description": "Failed to send SMS",
	},
	"555": {
		"name":        "block_ip_address",
		"description": "Your IP address has been blocked",
	},
	"560": {
		"name":        "mobile_number",
		"description": "Mobile number is invalid",
	},
	"562": {
		"name":        "send_date_&_time",
		"description": "Send date & time is invalid.",
	},
	"564": {
		"name":        "reply_id_is_invalid",
		"description": "Reply ID is invalid.",
	},
	"566": {
		"name":        "reply_id_doesn't_exist",
		"description": "Reply ID doesn't exist.",
	},
	"568": {
		"name":        "au_title",
		"description": "AU SMS title is invalid",
	},
	"569": {
		"name":        "softbank_title",
		"description": "Softbank SMS title is invalid",
	},
	"570": {
		"name":        "basic_sms_text_id",
		"description": "basic SMS text ID is invalid",
	},
	"571": {
		"name":        "sending_attempts_is_invalid",
		"description": "Sending attempts is invalid.",
	},
	"572": {
		"name":        "resending_interval_is_invalid",
		"description": "Resending interval is invalid.",
	},
	"573": {
		"name":        "status",
		"description": "Status is invalid",
	},
	"574": {
		"name":        "sms_id",
		"description": "SMS ID is invalid",
	},
	"575": {
		"name":        "docomo",
		"description": "Docomo is invalid",
	},
	"576": {
		"name":        "au",
		"description": "au is invalid",
	},
	"577": {
		"name":        "soft_bank",
		"description": "Soft Bank is invalid",
	},
	"579": {
		"name":        "gateway",
		"description": "Gateway is invalid",
	},
	"580": {
		"name":        "sms_title",
		"description": "SMS title is invalid",
	},
	"585": {
		"name":        "basic_sms_text",
		"description": "Basic SMS text is invalid",
	},
	"586": {
		"name":        "sms_text_2_(forSDP)",
		"description": "SMS text 2 is invalid",
	},
	"587": {
		"name":        "not_unique_sms_id",
		"description": "SMS ID is not unique",
	},
	"588": {
		"name":        "sms_text_for_docomo",
		"description": "SMS text for Docomo is invalid",
	},
	"589": {
		"name":        "sms_text_for_softbank",
		"description": "SMS text for Softbank is invalid",
	},
	"590": {
		"name":        "original_url",
		"description": "Original URL is invalid",
	},
	"591": {
		"name":        "sms_text_is_disabled",
		"description": "SMS text is disabled",
	},
	"592": {
		"name":        "time_is_disabled",
		"description": "Time is disabled.",
	},
	"593": {
		"name":        "sms",
		"description": "Wait return SMS is invalid",
	},
	"594": {
		"name":        "return_sms",
		"description": "Return SMS is invalid",
	},
	"595": {
		"name":        "2_way_sms_function_is_invalid",
		"description": "2-way SMS is disabled",
	},
	"598": {
		"name":        "docomo_title",
		"description": "Docomo SMS title is invalid",
	},
	"599": {
		"name":        "resending_is_disabled",
		"description": "Resending is disabled.",
	},
	"600": {
		"name":        "keep_chatid",
		"description": "Keep chatid is invalid",
	},
	"601": {
		"name":        "sms_title_function_is_disabled",
		"description": "SMS title function is disabled",
	},
	"602": {
		"name":        "sim_title",
		"description": "SIM SMS title is invalid.",
	},
	"603": {
		"name":        "reminder_delay",
		"description": "Reminder delay is invalid",
	},
	"604": {
		"name":        "reminder_text",
		"description": "Reminder text is invalid",
	},
	"605": {
		"name":        "type",
		"description": "Invalid type",
	},
	"606": {
		"name":        "3rd_party_api_is_disabled",
		"description": "3rd party API is disabled",
	},
	"608": {
		"name":        "registration_date",
		"description": "Invalid Registration date",
	},
	"610": {
		"name":        "hlr_is_disabled",
		"description": "HLR is disabled",
	},
	"612": {
		"name":        "original_url2",
		"description": "612 Original URL 2 is invalid",
	},
	"613": {
		"name":        "original_url3",
		"description": "613 Original URL 3 is invalid",
	},
	"614": {
		"name":        "original_url4",
		"description": "614 Original URL 4 is invalid",
	},
	"615": {
		"name":        "json_format_is_incorrect",
		"description": "615 Incorrect JSON format",
	},
	"617": {
		"name":        "memo_is_disabled",
		"description": "Memo is disabled",
	},
	"618": {
		"name":        "basic_sms_text_is_long_for_docomo_api",
		"description": "618 SMS text is long for Docomo API",
	},
	"619": {
		"name":        "basic_sms_text_is_long_for_sdp",
		"description": "619 SMS text is long for SDP",
	},
	"620": {
		"name":        "basic_sms_text_is_long_for_softbank_api",
		"description": "620 SMS text is long for Softbank API",
	},
	"621": {
		"name":        "reminder_text_is_long_for_docomo_api",
		"description": "Reminder text is long for Docomo API",
	},
	"622": {
		"name":        "reminder_text_is_long_for_sdp",
		"description": "Reminder text is long for SDP",
	},
	"623": {
		"name":        "reminder_text_is_long_for_softbank_api",
		"description": "Reminder text is long for Softbank API",
	},
	"624": {
		"name":        "duplicated_smsid",
		"description": "Duplicated SMSID",
	},
	"631": {
		"name":        "resending_parameters_can_not_be_changed",
		"description": "Resending parameters can not be changed",
	},
	"632": {
		"name":        "rakuten_title_is_invalid",
		"description": "Rakuten title is invalid",
	},
	"633": {
		"name":        "rakuten_text_is_invalid",
		"description": "Rakuten text is invalid",
	},
	"634": {
		"name":        "main_text_is_too_long_for_rakuten",
		"description": "Main text is too long for Rakuten",
	},
	"635": {
		"name":        "reminder_text_is_too_long_for_rakuten",
		"description": "Reminder text is too long for Rakuten",
	},
	"636": {
		"name":        "rakuten_is_invalid",
		"description": "Rakuten is invalid",
	},
	"637": {
		"name":        "rcs_is_disabled",
		"description": "RCS is disabled",
	},
	"639": {
		"name":        "original_url_code_is_disabled",
		"description": "639 Original URL code is disabled",
	},
	"640": {
		"name":        "original_url_code",
		"description": "640 Original URL code is invalid",
	},
	"641": {
		"name":        "original_url_code_2",
		"description": "641 Original URL code 2 is invalid",
	},
	"642": {
		"name":        "original_url_code_3",
		"description": "642 Original URL code 3 is invalid",
	},
	"643": {
		"name":        "original_url_code_4",
		"description": "643 Original URL code 4 is invalid",
	},
	"666": {
		"name":        "prior_to_block_ip_address",
		"description": "Prior to block IP address",
	},
}
