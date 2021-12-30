package common

import "github.com/spf13/pflag"

type CliOptions interface {
	Flags() (fss pflag.FlagSet)
	Validate() []error
}
