package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagEndTime           = "end-time"
	FlagDelayed           = "delayed"
	FlagStreamPeriodsFile = "stream-periods-file"
	FlagCancellable       = "cancellable"
)

var FsStreamSend = flag.NewFlagSet("", flag.ContinueOnError)

func init() {
	FsStreamSend.String(FlagEndTime, "", "end-time ")
	FsStreamSend.Bool(FlagDelayed, false, "use to set delayed stream")
	FsStreamSend.String(FlagStreamPeriodsFile, "", "stream periods json file")
	FsStreamSend.Bool(FlagCancellable, false, "use to set cancellable stream")
}
