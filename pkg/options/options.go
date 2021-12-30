package options

import "github.com/spf13/pflag"

type ServerOptions struct {
	Name   string
	IP     string
	Port   int
	Config string
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		Name: "Server_name",
		IP:   "127.0.0.1",
		Port: 80,
	}
}

// Validate verifies flags passed to MySQLOptions.
func (o *ServerOptions) Validate() []error {
	errs := []error{}

	return errs
}

// AddFlags adds flags related to mysql storage for a specific APIServer to the specified FlagSet.
func (o *ServerOptions) AddFlags(fs *pflag.FlagSet) (flagset *pflag.FlagSet) {
	fs.StringVarP(&o.Name, "name", "n", o.Name, ""+
		"Server Name")

	fs.StringVarP(&o.IP, "ip", "i", o.IP, ""+
		"server ip")

	fs.IntVarP(&o.Port, "port", "p", o.Port, "")
	fs.StringVarP(&o.Config, "config", "c", "", "config file path")
	return fs

}
