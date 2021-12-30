package template

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	factory "jimmycmd/internal/common"
)

func NewCmdSys(f factory.FlagSetFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sys",
		Short:   "sys command",
		Example: infoExample,
		Run: func(cmd *cobra.Command, args []string) {
			RunSys(f, cmd, args)
		},
	}
	cmd.Flags().StringP("sysname", "i", "sssss", "sysusername")
	cmd.Flags().IntP("sysage", "a", 123, "sysuserage")
	return cmd

}
func RunSys(f factory.FlagSetFactory, cmd *cobra.Command, args []string) error {
	/*name, err :=cmd.Flags().GetString("name")
	if err!= nil{
		glog.Fatalf("error accessing flag %s for command %s: %v", name, cmd.Name(), err)
		return err
	}*/
	name, err := cmd.Flags().GetString("sysname")
	age, err1 := cmd.Flags().GetInt("sysage")
	if err != nil {
		glog.Fatalf("error accessing flag %s for command %s: %v", name, cmd.Name(), err)
		return err
	}
	if err1 != nil {
		glog.Fatalf("error accessing flag %s for command %s: %v", name, cmd.Name(), err)
		return err1
	}
	fmt.Println(name, age)
	//fmt.Println("this is Sys command")
	return nil
}
