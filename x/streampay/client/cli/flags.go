package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagDuration          = "duration"
	FlagDelayed           = "delayed"
	FlagStreamPeriodsFile = "stream-periods-file"
	FlagCancellable       = "cancellable"
	FlagPaymentFee        = "payment-fee"
)

var FsStreamSend = flag.NewFlagSet("", flag.ContinueOnError)

func init() {
	FsStreamSend.String(FlagDuration, "", "to set duration of stream payment")
	FsStreamSend.Bool(FlagDelayed, false, "to create a delayed stream payment")
	FsStreamSend.String(FlagStreamPeriodsFile, "", "stream periods json file")
	FsStreamSend.Bool(FlagCancellable, false, "to create cancellable stream payment")
	FsStreamSend.String(FlagPaymentFee, "", "payment fee")
}
