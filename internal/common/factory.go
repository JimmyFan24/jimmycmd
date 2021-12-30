package common

import "github.com/spf13/pflag"

type FlagSetFactory struct {
	flags *pflag.FlagSet
}

func NewFlagSetFactory() FlagSetFactory {
	flags := pflag.NewFlagSet("", pflag.ContinueOnError)
	f := FlagSetFactory{
		flags: flags,
	}

	return f

}
