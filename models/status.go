package models

const (
	// StatusFailed occurs when the tranmission over the network of the carrier mobile device fails
	StatusFailed = "failed"
	// StatusDelivered occurs a success to deliver the SMS to the target number
	StatusDelivered = "delivered"
	// StatusSending occurs when the mediasms is processing the transmission
	StatusSending = "sending"
	// StatusRedirected occurs when the shortened link is clicked
	StatusRedirected = "redirected"
	// StatusNotSent occurs when it is queued and it is not yet processing
	StatusNotSent = "not_sent"
	// StatusFailedInvalidPhoneNumber occurs when The destination is registered in MediaSMS-CONSOLE as a non-sendable number. Displayed when the power is turned off or NW error occurs 10 times or more with the same number within 24 hours.
	StatusFailedInvalidPhoneNumber = "failed_invalid_phone_number"
	// StatusFailedAuOutOfServiceArea occurs when The device which has au sim is either status (Turned off device, Out of service, talking, time out, zone changed when it's called, phone app is not available to launch, memory issue on the device, there no app to show sms message on the device.)
	StatusFailedAuOutOfServiceArea = "failed_au_out_of_service_area"
	// StatusFailedAuNwFailure occurs when the mobile device belongs to au and there is a failure in the au side NW, or the mobile device belongs to au and the au side switchboard is in a congested state or malfunction.
	StatusFailedAuNwFailure = "failed_au_network_failure"
	// StatusFailedAuNwOther occurs when The mobile device does not belong to au. For example, when the SMS option is not added with a cheap SIM.
	StatusFailedAuNwOther = "failed_au_network_other"
	// StatusFailedAuNotPaid occurs when The mobile device belongs to au and is closed to charges.
	StatusFailedAuNotPaid = "failed_au_not_paid"
	// StatusProcessing occurs when each communication carrier is performing transmission processing.
	StatusProcessing = "failed_processing"
	// StatusFailedAuOther occurs when he mobile device belongs to au and has failed due to reasons other than the above. (Including rejection etc.)
	StatusFailedAuOther = "failed_au_other"
	// StatusFailedDocomoOther occurs when the mobile device belongs to Docomo and has failed due to unknown reasons
	StatusFailedDocomoOther = "failed_docomo_other"
	// StatusFailedDocomoOutOfServiceArea occurs when the mobile device belongs to Docomo and is either turned off or out of service
	StatusFailedDocomoOutOfServiceArea = "failed_docomo_out_of_service_area"
	// StatusFailedDocomoNwMalfunction occurs when The mobile device belongs to Docomo, and the NW on the Docomo side has a failure or is in a congested state.
	StatusFailedDocomoNwMalfunction = "failed_docomo_network_malfunction"
	// StatusFailedDocomoNwOther occurs when the mobile device does not belong to Docomo, or the mobile device belongs to DoCoMo and reception is rejected, or the mobile device belongs to Docomo and the processing by the mobile device has failed. When SMS is not attached as an option such as cheap SIM.
	StatusFailedDocomoNwOther = "failed_docomo_network_other"
	// StatusFailedInvalidRecords occurs when Phone number is not valid phone number such as the digit of the number is not correct.
	StatusFailedInvalidRecords = "failed_invalid_records"
	// StatusFailedSoftbankOther occurs when The mobile device belongs to Softbank and has failed to unknown reasons.
	StatusFailedSoftbankOther = "failed_softbank_other"
	// StatusFailedSoftbankOutOfServiceArea occurs when The mobile device belongs to Softbank and is either turned off or out of service.
	StatusFailedSoftbankOutOfServiceArea = "failed_softbank_out_of_service_area"
	// StatusFailedSoftbankNwMalfunction occurs when the mobile device belongs to Softbank, and a NW failure has occurred on the Softbank side.
	StatusFailedSoftbankNwMalfunction = "failed_softbank_network_malfunction"
	// StatusFailedSoftbankNwOther occurs when Mobile devices do not belong to Softbank. For example, when the SMS option is not added with a cheap SIM.
	StatusFailedSoftbankNwOther = "failed_softbank_network_other"
	// StatusFailedSoftbankReceptionRefusal occurs when the mobile device belongs to Softbank and is rejected.
	StatusFailedSoftbankReceptionRefusal = "failed_softbank_reception_refusal"
	// StatusFailedHlrAPI occurs when "Non-existent number", "Data-only SIM does not have SMS function", "PHS without SMS function", "SMS reception block from carrier" are applicable.
	StatusFailedHlrAPI = "failed_softbank_hlr_api"
	// StatusScheduled occurs when the status is earlier than the set transmission date and time.
	StatusScheduled = "scheduled"
	// StatusDelayed occurs when the pause button is clicked after uploading the CSV file.
	StatusDelayed = "delayed"
	// StatusInboundMessageReceived occurs when a reply is received
	StatusInboundMessageReceived = "inbound_message_received"
	// StatusFailedOutsideOfAllowedHours occurs when  the SMS sending permission time is set for the account, it will be displayed if you send a request outside the corresponding time
	StatusFailedOutsideOfAllowedHours = "failed_outside_of_allowed_hours"
	// StatusFailedTwoWayDisabled Refer to MediaSMS-CONSOLE Bidirectional API Specification_ver1.0.5 P9
	StatusFailedTwoWayDisabled = "failed_two_way_disabled"
	// StatusFailedDocomoOtherUnknown occurs when the SMS sending permission time is set for the account, it will be displayed if you send a request outside the corresponding time.
	StatusFailedDocomoOtherUnknown = "failed_docomo_other_unknown"
	// StatusFailedDuplicated occurs when Displayed when duplicate SMS is sent. If the number and text are the same, the check will be carried out within 1 minute, and if applicable, this result will be returned.
	StatusFailedDuplicated = "failed_duplicated"
	// StatusFailedKeepChatID Refer to MediaSMS-CONSOLE 双方向API仕様書_ver1.0.5 P10を参照
	StatusFailedKeepChatID = "failed_keep_chat_id"
	// StatusFailedDocomoPartialReception occurs when Displayed when only a part of the long text of SMS is received when using long text SMS. For example, in the long column, the format 6/7 shows that 6 out of 7 received.
	StatusFailedDocomoPartialReception = "failed_docomo_partial_reception"
	// StatusFailedNoConnectionAvailable
	StatusFailedNoConnectionAvailable = "failed_no_connection_available"
	// StatusCanceled occurs when occurs when canceled before sending.
	StatusCanceled = "canceled"
	// StatusLimitReached occurs when restricted by the transmission upper limit function.
	StatusLimitReached = "limit_reached"
	// StatusFailedNgWord occurs when It is displayed when the keyword set as the NG word on the management screen is entered, and it is in the unsent state. In addition, "NG word limit" is entered in the memo parameter and returned.
	StatusFailedNgWord = "failed_ng_word"
	// StatusFailedUnsubscribed occurs when the unsubscribe URL and it is registered in the transmission exclusion list.
	StatusFailedUnsubscribed = "unsubscribed"
	// StatusPartialDelivery occurs when displayed when a long text is sent overseas and only a part of the long text SMS is received on the terminal side. 67 characters will be delivered to the au terminal.
	StatusPartialDelivery = "partial_delivery"
	// StatusApprovalPending occurs after setting the "Simultaneous Approval Function" option on the management screen, it is displayed when the mass transmission list is waiting for approval.
	StatusApprovalPending = "approval_pending"
	// StatusFailedRakutenOther occurs when the mobile belongs to rakuten, but fails between the base station and the terminal. Or, the mobile device belongs to rakuten and reception is rejected, or the mobile device belongs to rakuten and processing by the mobile device has failed.
	StatusFailedRakutenOther = "failed_rakuten_other"
	// StatusFailedRakutenNwMalfunction occurs when the mobile device does not belong to rakuten. For example, when the SMS option is not added with a cheap SIM.
	StatusFailedRakutenNwMalfunction = "failed_rakuten_network_malfunction"
	// StatusFailedRakutenPartialReception occurs when displayed when only a part of the long text of SMS is received when using long text SMS. For example, in the long column, the format 6/7 shows that 6 out of 7 received.
	StatusFailedRakutenPartialReception = "failed_rakuten_partial_reception"
	// StatusFailedRakutenOutOfServiceArea occurs when the mobile unit belongs to Rakuten and is either turned off or out of service.
	StatusFailedRakutenOutOfServiceArea = "failed_rakuten_out_of_service_area"
)

// RetriableEvents is a list of the events that will be retried according to each provider configuration
var RetriableEvents = []string{
	StatusFailed,
	StatusFailedDocomoOutOfServiceArea,
	StatusFailedDocomoNwMalfunction,
	StatusFailedDocomoNwOther,
	StatusFailedDocomoOther,
	StatusFailedDocomoOtherUnknown,
	StatusFailedDocomoPartialReception,
	StatusFailedAuOutOfServiceArea,
	StatusFailedAuNwFailure,
	StatusFailedAuNwOther,
	StatusFailedAuNotPaid,
	StatusFailedAuOther,
	StatusFailedSoftbankOutOfServiceArea,
	StatusFailedSoftbankNwMalfunction,
	StatusFailedSoftbankNwOther,
	StatusFailedSoftbankReceptionRefusal,
	StatusFailedSoftbankOther,
	StatusPartialDelivery,
	StatusFailedRakutenOther,
	StatusFailedRakutenNwMalfunction,
	StatusFailedRakutenPartialReception,
	StatusFailedRakutenOutOfServiceArea,
}
