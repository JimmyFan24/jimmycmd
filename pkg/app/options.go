package app

import (
	"github.com/spf13/pflag"
)

type CliOptions interface {
	Flags() (fs *pflag.FlagSet)
	Validate() []error
}
type ConfigurableOptions interface {
	// ApplyFlags parsing parameters from the command line or configuration file
	// to the options instance.
	ApplyFlags() []error
}

type CompleteableOptions interface {
	Complete() error
}
type PrintableOptions interface {
	String() string
}
