package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagDuration          = "duration"
	FlagDelayed           = "delayed"
	FlagStreamPeriodsFile = "stream-periods-file"
	FlagCancellable       = "cancellable"
)

var FsStreamSend = flag.NewFlagSet("", flag.ContinueOnError)

func init() {
	FsStreamSend.String(FlagDuration, "", "duration of stream")
	FsStreamSend.Bool(FlagDelayed, false, "use to set delayed stream")
	FsStreamSend.String(FlagStreamPeriodsFile, "", "stream periods json file")
	FsStreamSend.Bool(FlagCancellable, false, "use to set cancellable stream")
}
