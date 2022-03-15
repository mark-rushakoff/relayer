package cmd

import (
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagURL                     = "url"
	flagSkip                    = "skip"
	flagTimeout                 = "timeout"
	flagJSON                    = "json"
	flagYAML                    = "yaml"
	flagFile                    = "file"
	flagPath                    = "path"
	flagMaxTxSize               = "max-tx-size"
	flagMaxMsgLength            = "max-msgs"
	flagIBCDenoms               = "ibc-denoms"
	flagTimeoutHeightOffset     = "timeout-height-offset"
	flagTimeoutTimeOffset       = "timeout-time-offset"
	flagMaxRetries              = "max-retries"
	flagThresholdTime           = "time-threshold"
	flagUpdateAfterExpiry       = "update-after-expiry"
	flagUpdateAfterMisbehaviour = "update-after-misbehaviour"
	flagOverride                = "override"
	flagPort                    = "port"
	flagOrder                   = "unordered"
	flagVersion                 = "version"
)

// flagger provides a builder pattern to add flags to a cobra Command.
type flagger struct {
	v   *viper.Viper
	cmd *cobra.Command
}

func newFlagger(v *viper.Viper, cmd *cobra.Command) *flagger {
	return &flagger{
		v:   v,
		cmd: cmd,
	}
}

// Command terminates the builder and returns the underlying command.
//
//    return newFlagger(v, cmd).Foo().Bar().Command()
func (f *flagger) Command() *cobra.Command {
	return f.cmd
}

func (f *flagger) IBCDenom() *flagger {
	f.cmd.Flags().BoolP(flagIBCDenoms, "i", false, "Display IBC denominations for sending tokens back to other chains")
	if err := f.v.BindPFlag(flagIBCDenoms, f.cmd.Flags().Lookup(flagIBCDenoms)); err != nil {
		panic(err)
	}
	return f
}

func (f *flagger) Height() *flagger {
	f.cmd.Flags().Int64(flags.FlagHeight, 0, "Height of headers to fetch")
	if err := f.v.BindPFlag(flags.FlagHeight, f.cmd.Flags().Lookup(flags.FlagHeight)); err != nil {
		panic(err)
	}
	return f
}

func (f *flagger) Pagination() *flagger {
	f.cmd.Flags().Uint64P(flags.FlagOffset, "o", 0, "pagination offset for query")
	f.cmd.Flags().Uint64P(flags.FlagLimit, "l", 10, "pagination limit for query")
	if err := f.v.BindPFlag(flags.FlagOffset, f.cmd.Flags().Lookup(flags.FlagOffset)); err != nil {
		panic(err)
	}
	if err := f.v.BindPFlag(flags.FlagLimit, f.cmd.Flags().Lookup(flags.FlagLimit)); err != nil {
		panic(err)
	}
	return f
}

func (f *flagger) YAML() *flagger {
	f.cmd.Flags().BoolP(flagYAML, "y", false, "output using yaml")
	if err := f.v.BindPFlag(flagYAML, f.cmd.Flags().Lookup(flagYAML)); err != nil {
		panic(err)
	}
	return f
}

func (f *flagger) SkipConfirm() *flagger {
	f.cmd.Flags().BoolP(flagSkip, "y", false, "answer yes to all questions")
	if err := f.v.BindPFlag(flagSkip, f.cmd.Flags().Lookup(flagSkip)); err != nil {
		panic(err)
	}
	return f
}

func (f *flagger) Path() *flagger {
	f.cmd.Flags().StringP(flagPath, "p", "", "specify the path to relay over")
	if err := f.v.BindPFlag(flagPath, f.cmd.Flags().Lookup(flagPath)); err != nil {
		panic(err)
	}
	return f
}

func (f *flagger) Timeouts() *flagger {
	f.cmd.Flags().Uint64P(flagTimeoutHeightOffset, "y", 0, "set timeout height offset for ")
	f.cmd.Flags().DurationP(flagTimeoutTimeOffset, "c", time.Duration(0), "specify the path to relay over")
	if err := f.v.BindPFlag(flagTimeoutHeightOffset, f.cmd.Flags().Lookup(flagTimeoutHeightOffset)); err != nil {
		panic(err)
	}
	if err := f.v.BindPFlag(flagTimeoutTimeOffset, f.cmd.Flags().Lookup(flagTimeoutTimeOffset)); err != nil {
		panic(err)
	}
	return f
}

func (f *flagger) JSON() *flagger {
	f.cmd.Flags().BoolP(flagJSON, "j", false, "returns the response in json format")
	if err := f.v.BindPFlag(flagJSON, f.cmd.Flags().Lookup(flagJSON)); err != nil {
		panic(err)
	}
	return f
}

func (f *flagger) File() *flagger {
	f.cmd.Flags().StringP(flagFile, "f", "", "fetch json data from specified file")
	if err := f.v.BindPFlag(flagFile, f.cmd.Flags().Lookup(flagFile)); err != nil {
		panic(err)
	}
	return f
}

func (f *flagger) Timeout() *flagger {
	f.cmd.Flags().StringP(flagTimeout, "o", "10s", "timeout between relayer runs")
	if err := f.v.BindPFlag(flagTimeout, f.cmd.Flags().Lookup(flagTimeout)); err != nil {
		panic(err)
	}
	return f
}

func (f *flagger) URL() *flagger {
	f.cmd.Flags().StringP(flagURL, "u", "", "url to fetch data from")
	if err := f.v.BindPFlag(flagURL, f.cmd.Flags().Lookup(flagURL)); err != nil {
		panic(err)
	}
	return f
}

func (f *flagger) Strategy() *flagger {
	f.cmd.Flags().StringP(flagMaxTxSize, "s", "2", "strategy of path to generate of the messages in a relay transaction")
	f.cmd.Flags().StringP(flagMaxMsgLength, "l", "5", "maximum number of messages in a relay transaction")
	if err := f.v.BindPFlag(flagMaxTxSize, f.cmd.Flags().Lookup(flagMaxTxSize)); err != nil {
		panic(err)
	}
	if err := f.v.BindPFlag(flagMaxMsgLength, f.cmd.Flags().Lookup(flagMaxMsgLength)); err != nil {
		panic(err)
	}
	return f
}

func (f *flagger) Retry() *flagger {
	f.cmd.Flags().Uint64P(flagMaxRetries, "r", 3, "maximum retries after failed message send")
	if err := f.v.BindPFlag(flagMaxRetries, f.cmd.Flags().Lookup(flagMaxRetries)); err != nil {
		panic(err)
	}
	return f
}

func (f *flagger) UpdateTime() *flagger {
	f.cmd.Flags().Duration(flagThresholdTime, 6*time.Hour, "time before to expiry time to update client")
	if err := f.v.BindPFlag(flagThresholdTime, f.cmd.Flags().Lookup(flagThresholdTime)); err != nil {
		panic(err)
	}
	return f
}

func (f *flagger) ClientParameters() *flagger {
	f.cmd.Flags().BoolP(flagUpdateAfterExpiry, "e", true,
		"allow governance to update the client if expiry occurs")
	f.cmd.Flags().BoolP(flagUpdateAfterMisbehaviour, "m", true,
		"allow governance to update the client if misbehaviour freezing occurs")
	if err := f.v.BindPFlag(flagUpdateAfterExpiry, f.cmd.Flags().Lookup(flagUpdateAfterExpiry)); err != nil {
		panic(err)
	}
	if err := f.v.BindPFlag(flagUpdateAfterMisbehaviour, f.cmd.Flags().Lookup(flagUpdateAfterMisbehaviour)); err != nil {
		panic(err)
	}
	return f
}

func (f *flagger) Override() *flagger {
	f.cmd.Flags().Bool(flagOverride, false, "option to not reuse existing client")
	if err := f.v.BindPFlag(flagOverride, f.cmd.Flags().Lookup(flagOverride)); err != nil {
		panic(err)
	}
	return f
}

func (f *flagger) Order() *flagger {
	f.cmd.Flags().BoolP(flagOrder, "o", true, "create an unordered channel")
	if err := f.v.BindPFlag(flagOrder, f.cmd.Flags().Lookup(flagOrder)); err != nil {
		panic(err)
	}
	return f
}

func (f *flagger) Version() *flagger {
	f.cmd.Flags().StringP(flagVersion, "v", "ics20-1", "version of channel to create")
	if err := f.v.BindPFlag(flagVersion, f.cmd.Flags().Lookup(flagVersion)); err != nil {
		panic(err)
	}
	return f
}

func (f *flagger) Port() *flagger {
	f.cmd.Flags().StringP(flagPort, "p", "transfer", "port to use when generating path")
	if err := f.v.BindPFlag(flagPort, f.cmd.Flags().Lookup(flagPort)); err != nil {
		panic(err)
	}
	return f
}
