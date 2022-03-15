package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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

func (f *flagger) IBCDenom() *flagger {
	ibcDenomFlags(f.v, f.cmd)
	return f
}

func (f *flagger) Height() *flagger {
	heightFlag(f.v, f.cmd)
	return f
}

func (f *flagger) Pagination() *flagger {
	paginationFlags(f.v, f.cmd)
	return f
}

func (f *flagger) YAML() *flagger {
	yamlFlag(f.v, f.cmd)
	return f
}

func (f *flagger) SkipConfirm() *flagger {
	skipConfirm(f.v, f.cmd)
	return f
}

// func (f *flagger) ChainsAdd

func (f *flagger) Path() *flagger {
	pathFlag(f.v, f.cmd)
	return f
}

func (f *flagger) Timeouts() *flagger {
	timeoutFlags(f.v, f.cmd)
	return f
}

func (f *flagger) JSON() *flagger {
	jsonFlag(f.v, f.cmd)
	return f
}

func (f *flagger) File() *flagger {
	fileFlag(f.v, f.cmd)
	return f
}

func (f *flagger) Timeout() *flagger {
	timeoutFlag(f.v, f.cmd)
	return f
}

func (f *flagger) URL() *flagger {
	urlFlag(f.v, f.cmd)
	return f
}

func (f *flagger) Strategy() *flagger {
	strategyFlag(f.v, f.cmd)
	return f
}

func (f *flagger) Retry() *flagger {
	retryFlag(f.v, f.cmd)
	return f
}

func (f *flagger) UpdateTime() *flagger {
	updateTimeFlags(f.v, f.cmd)
	return f
}

func (f *flagger) ClientParameters() *flagger {
	clientParameterFlags(f.v, f.cmd)
	return f
}

func (f *flagger) Override() *flagger {
	overrideFlag(f.v, f.cmd)
	return f
}

func (f *flagger) Order() *flagger {
	orderFlag(f.v, f.cmd)
	return f
}

func (f *flagger) Version() *flagger {
	versionFlag(f.v, f.cmd)
	return f
}

func (f *flagger) Port() *flagger {
	portFlag(f.v, f.cmd)
	return f
}
