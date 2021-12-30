package template

import (
	"fmt"
	//"github.com/golang/glog"
	"github.com/spf13/cobra"
	factory "jimmycmd/internal/common"
	//hoststat "github.com/likexian/host-stat-go"
)

/*type Info struct {
	HostName  string
	IPAddress string
	OSRelease string
	CPUCore   uint64
	MemTotal  string
	MemFree   string
}*/
var (
	infoExample = `
		# Print the host information
		cmdctl info

		# Specify a server password
		cmdctl info -p newpass

		# Print details
		cmdctl info -d`
)

func NewCmdInfo(f factory.FlagSetFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "info",
		Short:   "Info command",
		Example: infoExample,
		Run: func(cmd *cobra.Command, args []string) {
			RunInfo(f, cmd, args)
		},
	}
	cmd.Flags().StringP("infoname", "i", "", "infousername")
	cmd.Flags().IntP("age", "a", 123, "infouserage")
	//cmd.Flags().StringP("passwd", "p", "", "Specify the server password.")
	//cmd.Flags().BoolP("detail", "d", false, "Print details.")
	return cmd
}
func RunInfo(f factory.FlagSetFactory, cmd *cobra.Command, args []string) error {
	name, err := cmd.Flags().GetString("infoname")
	age, err1 := cmd.Flags().GetInt("age")
	//passwd ,_:= cmd.Flags().GetString("passwd")
	//detail,_ := cmd.Flags().GetBool("detail")
	if err != nil {
		return cmd.Help()
	}
	if err1 != nil {
		return cmd.Help()
	}
	fmt.Println(name, age)
	return nil
}
