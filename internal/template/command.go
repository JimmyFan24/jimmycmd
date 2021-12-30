package template

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"jimmycmd/internal/common"
	"strings"
)

const (
	DefaultErrorExitCode = 1
)

var cfgFile string

type cmd *cobra.Command

type Command struct {
	usage   string
	desc    string
	command cmd
	comands []cmd
	options common.Options
}

/*func argstostring(args *Args) []string  {
	var argsString [] string
	argsString[0] = args.name
	argsString[1] = string(args.age)
	return argsString
}*/

func NewCommand(f common.FlagSetFactory) *cobra.Command {
	// Parent common to which all subcommands are added.
	cmds := &cobra.Command{
		Use:   "jimmycmd",
		Short: "Basic Command",
		Run:   RunFunc,
	}
	groups := CommandGroups{
		{
			Message: "Command One",
			Commands: []*cobra.Command{
				NewCmdInfo(f),
			},
		}, {Message: "Command Two",
			Commands: []*cobra.Command{
				NewCmdSys(f),
			}},
	}

	groups.Add(cmds)
	//persistenflags是所有子命令都可见的参数
	cmds.Flags().StringP("path", "p", "", "filepath")

	cmds.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is xxx.json)")
	cmds.PersistentFlags().BoolP("debug", "", false, "enable the debug mode")
	//pflag.Parse()
	cobra.OnInitialize(initConfig)
	//viper.BindPFlag("path",cmds.Flags().Lookup("path"))
	return cmds
}

type Pathconfig struct {
	Path string
}

var c Pathconfig

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		fmt.Printf("-----")
		viper.AddConfigPath(".")
		viper.AddConfigPath("C:\\Users\\jimmy\\GolandProjects\\jimmycmd")
		viper.SetConfigName("test")
	}

	viper.SetConfigType("json")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CMDCTL")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.ReadInConfig()
	//var c Pathconfig
	err := viper.Unmarshal(&c)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}
	//fmt.Printf("use config:%s\n", viper.ConfigFileUsed())
	//fmt.Printf("config_path:%s\n", c.Path)
}

func RunFunc(cmd *cobra.Command, args []string) {
	//path,_ :=cmd.Flags().GetString("path")
	//fmt.Printf(path)

	fmt.Printf("BaseRunFunc is running:%s\n", c.Path)

	//fmt.Printf(path)
}
func runhelp(cmd *cobra.Command, arg []string) {
	cmd.Help()
}

func NewCmdVersion(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Short:   "Print the client and server version information",
		Long:    "Print the client and server version information for the current context",
		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			//options := new(Options)
			//checkErr(options.Complete(cmd))
			//checkErr(options.Validate(cmd))
			//checkErr(options.Run(out))
			fmt.Printf("------------")
		},
	}
	cmd.Flags().BoolP("short", "", false, "Print just the version number.")
	cmd.Flags().StringP("output", "o", "", "One of 'yaml' or 'json'.")
	return cmd
}

var ErrExit = fmt.Errorf("exit")

func checkErr(prefix string, err error, handleErr func(string, int)) {
	switch {
	case err == nil:
		return
	case err == ErrExit:
		handleErr("", DefaultErrorExitCode)
		return
	default:
		handleErr("1", DefaultErrorExitCode)
	}
}
