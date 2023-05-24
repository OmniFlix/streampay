package types

const (
	EventTypeEndStreamPayment    = "end_stream_payment"
	EventTypeCreateStreamPayment = "create_stream_payment"
	EventTypeStopStreamPayment   = "stop_stream_payment"
	EventTypeClaimStreamedAmount = "claim_streamed_amount"

	EventAttributePaymentId      = "payment-id"
	EventAttributeSender         = "sender"
	EventAttributeRecipient      = "recipient"
	EventAttributeClaimer        = "claimer"
	EventAttributeAmount         = "amount"
	EventAttributePaymentType    = "payment-type"
	EventAttributeEndTime        = "endtime"
	EventAttributeReturnedAmount = "returned-amount"
	EventAttributeStreamedAmount = "streamed-amount"
)
