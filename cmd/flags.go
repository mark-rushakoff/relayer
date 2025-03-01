package cmd

import (
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
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

func ibcDenomFlags(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().BoolP(flagIBCDenoms, "i", false, "Display IBC denominations for sending tokens back to other chains")
	if err := viper.BindPFlag(flagIBCDenoms, cmd.Flags().Lookup(flagIBCDenoms)); err != nil {
		panic(err)
	}
	return cmd
}

func heightFlag(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().Int64(flags.FlagHeight, 0, "Height of headers to fetch")
	if err := viper.BindPFlag(flags.FlagHeight, cmd.Flags().Lookup(flags.FlagHeight)); err != nil {
		panic(err)
	}
	return cmd
}

func paginationFlags(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().Uint64P(flags.FlagOffset, "o", 0, "pagination offset for query")
	cmd.Flags().Uint64P(flags.FlagLimit, "l", 10, "pagination limit for query")
	if err := viper.BindPFlag(flags.FlagOffset, cmd.Flags().Lookup(flags.FlagOffset)); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag(flags.FlagLimit, cmd.Flags().Lookup(flags.FlagLimit)); err != nil {
		panic(err)
	}
	return cmd
}

func yamlFlag(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().BoolP(flagYAML, "y", false, "output using yaml")
	if err := viper.BindPFlag(flagYAML, cmd.Flags().Lookup(flagYAML)); err != nil {
		panic(err)
	}
	return cmd
}

func skipConfirm(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().BoolP(flagSkip, "y", false, "output using yaml")
	if err := viper.BindPFlag(flagSkip, cmd.Flags().Lookup(flagSkip)); err != nil {
		panic(err)
	}
	return cmd
}

func chainsAddFlags(cmd *cobra.Command) *cobra.Command {
	fileFlag(cmd)
	urlFlag(cmd)
	return cmd
}

func pathFlag(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringP(flagPath, "p", "", "specify the path to relay over")
	if err := viper.BindPFlag(flagPath, cmd.Flags().Lookup(flagPath)); err != nil {
		panic(err)
	}
	return cmd
}

func timeoutFlags(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().Uint64P(flagTimeoutHeightOffset, "y", 0, "set timeout height offset for ")
	cmd.Flags().DurationP(flagTimeoutTimeOffset, "c", time.Duration(0), "specify the path to relay over")
	if err := viper.BindPFlag(flagTimeoutHeightOffset, cmd.Flags().Lookup(flagTimeoutHeightOffset)); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag(flagTimeoutTimeOffset, cmd.Flags().Lookup(flagTimeoutTimeOffset)); err != nil {
		panic(err)
	}
	return cmd
}

func jsonFlag(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().BoolP(flagJSON, "j", false, "returns the response in json format")
	if err := viper.BindPFlag(flagJSON, cmd.Flags().Lookup(flagJSON)); err != nil {
		panic(err)
	}
	return cmd
}

func fileFlag(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringP(flagFile, "f", "", "fetch json data from specified file")
	if err := viper.BindPFlag(flagFile, cmd.Flags().Lookup(flagFile)); err != nil {
		panic(err)
	}
	return cmd
}

func timeoutFlag(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringP(flagTimeout, "o", "10s", "timeout between relayer runs")
	if err := viper.BindPFlag(flagTimeout, cmd.Flags().Lookup(flagTimeout)); err != nil {
		panic(err)
	}
	return cmd
}

func getTimeout(cmd *cobra.Command) (time.Duration, error) {
	to, err := cmd.Flags().GetString(flagTimeout)
	if err != nil {
		return 0, err
	}
	return time.ParseDuration(to)
}

func urlFlag(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringP(flagURL, "u", "", "url to fetch data from")
	if err := viper.BindPFlag(flagURL, cmd.Flags().Lookup(flagURL)); err != nil {
		panic(err)
	}
	return cmd
}

func strategyFlag(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringP(flagMaxTxSize, "s", "2", "strategy of path to generate of the messages in a relay transaction")
	cmd.Flags().StringP(flagMaxMsgLength, "l", "5", "maximum number of messages in a relay transaction")
	if err := viper.BindPFlag(flagMaxTxSize, cmd.Flags().Lookup(flagMaxTxSize)); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag(flagMaxMsgLength, cmd.Flags().Lookup(flagMaxMsgLength)); err != nil {
		panic(err)
	}
	return cmd
}

func getAddInputs(cmd *cobra.Command) (file string, url string, err error) {
	file, err = cmd.Flags().GetString(flagFile)
	if err != nil {
		return
	}

	url, err = cmd.Flags().GetString(flagURL)
	if err != nil {
		return
	}

	if file != "" && url != "" {
		return "", "", errMultipleAddFlags
	}

	return
}

func retryFlag(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().Uint64P(flagMaxRetries, "r", 3, "maximum retries after failed message send")
	if err := viper.BindPFlag(flagMaxRetries, cmd.Flags().Lookup(flagMaxRetries)); err != nil {
		panic(err)
	}
	return cmd
}

func updateTimeFlags(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().Duration(flagThresholdTime, 6*time.Hour, "time before to expiry time to update client")
	if err := viper.BindPFlag(flagThresholdTime, cmd.Flags().Lookup(flagThresholdTime)); err != nil {
		panic(err)
	}
	return cmd
}

func clientParameterFlags(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().BoolP(flagUpdateAfterExpiry, "e", true,
		"allow governance to update the client if expiry occurs")
	cmd.Flags().BoolP(flagUpdateAfterMisbehaviour, "m", true,
		"allow governance to update the client if misbehaviour freezing occurs")
	if err := viper.BindPFlag(flagUpdateAfterExpiry, cmd.Flags().Lookup(flagUpdateAfterExpiry)); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag(flagUpdateAfterMisbehaviour, cmd.Flags().Lookup(flagUpdateAfterMisbehaviour)); err != nil {
		panic(err)
	}
	return cmd
}

func overrideFlag(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().Bool(flagOverride, false, "option to not reuse existing client")
	if err := viper.BindPFlag(flagOverride, cmd.Flags().Lookup(flagOverride)); err != nil {
		panic(err)
	}
	return cmd
}

func orderFlag(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().BoolP(flagOrder, "o", true, "create an unordered channel")
	if err := viper.BindPFlag(flagOrder, cmd.Flags().Lookup(flagOrder)); err != nil {
		panic(err)
	}
	return cmd
}

func versionFlag(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringP(flagVersion, "v", "ics20-1", "version of channel to create")
	if err := viper.BindPFlag(flagVersion, cmd.Flags().Lookup(flagVersion)); err != nil {
		panic(err)
	}
	return cmd
}

func portFlag(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringP(flagPort, "p", "transfer", "port to use when generating path")
	if err := viper.BindPFlag(flagPort, cmd.Flags().Lookup(flagPort)); err != nil {
		panic(err)
	}
	return cmd
}
