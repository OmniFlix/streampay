package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagEndTime = "end-time"
	FlagDelayed = "delayed"
)

var (
	FsStreamSend = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsStreamSend.String(FlagEndTime, "", "end-time ")
	FsStreamSend.Bool(FlagDelayed, false, "use to set delayed stream")
}
