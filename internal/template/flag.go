package template

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"jimmycmd/internal/common"
)

/*func NewFlag() *pflag.Flag {

	f := JFlag{
		flagset: flagset,
	}

	return f
}

func InitFlags(flagset JFlag) {
	flagset.flagset.StringVarP(&name,"name","n","jimmy","")
	flagset.flagset.IntVarP(&age,"age","a",123,"-----------")



}*/
func InitArgs(name string, age int) *common.Args {
	arg := common.NewArgs(name, age)
	return arg
}

/*func init()  {

	InitFlags(NewFlagSet())

}*/
func GetFlagBool(cmd *cobra.Command, flag string) bool {
	b, err := cmd.Flags().GetBool(flag)
	if err != nil {
		glog.Fatalf("error accessing flag %s for common %s: %v", flag, cmd.Name(), err)
	}
	return b
}

// Assumes the flag has a default value.
func GetFlagInt(cmd *cobra.Command, flag string) int {
	i, err := cmd.Flags().GetInt(flag)
	if err != nil {
		glog.Fatalf("error accessing flag %s for common %s: %v", flag, cmd.Name(), err)
	}
	return i
}
