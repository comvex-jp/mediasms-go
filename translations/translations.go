package translations

import "github.com/comvex-jp/mediasms-go/models"

// TranslationMap converts the int status value to a readable code
var TranslationMap = map[int]string{
	0:  models.StatusFailed,
	1:  models.StatusDelivered,
	2:  models.StatusSending,
	3:  models.StatusRedirected,
	4:  models.StatusNotSent,
	5:  models.StatusFailedInvalidPhoneNumber,
	7:  models.StatusFailedAuOutOfServiceArea,
	8:  models.StatusFailedAuNwFailure,
	9:  models.StatusFailedAuNwOther,
	10: models.StatusFailedAuNotPaid,
	11: models.StatusProcessing,
	12: models.StatusFailedAuOther,
	13: models.StatusFailedDocomoOther,
	14: models.StatusFailedDocomoOutOfServiceArea,
	15: models.StatusFailedDocomoNwMalfunction,
	16: models.StatusFailedDocomoOther,
	17: models.StatusFailedInvalidRecords,
	18: models.StatusFailedSoftbankOther,
	19: models.StatusFailedSoftbankOutOfServiceArea,
	20: models.StatusFailedSoftbankNwMalfunction,
	21: models.StatusFailedSoftbankNwOther,
	22: models.StatusFailedSoftbankReceptionRefusal,
	23: models.StatusFailedHlrAPI,
	24: models.StatusScheduled,
	25: models.StatusDelayed,
	26: models.StatusInboundMessageReceived,
	27: models.StatusFailedOutsideOfAllowedHours,
	28: models.StatusFailedTwoWayDisabled,
	29: models.StatusFailedDocomoOtherUnknown,
	30: models.StatusFailedDuplicated,
	31: models.StatusFailedKeepChatID,
	32: models.StatusFailedDocomoPartialReception,
	33: models.StatusFailedNoConnectionAvailable,
	34: models.StatusCanceled,
	35: models.StatusLimitReached,
	36: models.StatusFailedNgWord,
	37: models.StatusFailedUnsubscribed,
	38: models.StatusPartialDelivery,
	39: models.StatusApprovalPending,
	40: models.StatusFailedRakutenOther,
	41: models.StatusFailedRakutenNwMalfunction,
	42: models.StatusFailedRakutenPartialReception,
	43: models.StatusFailedRakutenOutOfServiceArea,
}
