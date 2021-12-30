package common

type Options struct {
	config *Config
	//Falgs *pflag.FlagSet

}

func NewOption(conf *Config) *Options {
	opt := &Options{
		config: conf,
	}

	return opt
}
