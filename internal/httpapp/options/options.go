package options

import (
	"github.com/spf13/pflag"
	"jimmycmd/pkg/options"
)

type Options struct {
	ServerOptions *options.ServerOptions
}

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	o := Options{
		ServerOptions: options.NewServerOptions(),
	}
	return &o
}

// Flags returns flags for a specific APIServer by section name.
func (o *Options) Flags() (fs *pflag.FlagSet) {
	flagset := pflag.CommandLine
	return o.ServerOptions.AddFlags(flagset)
}

// Validate checks Options and return a slice of found errs.
func (o *Options) Validate() []error {
	var errs []error

	return errs
}
