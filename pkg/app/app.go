package app

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	//OptionItem "jimmycmd/pkg/options"
)

type App struct {
	basename    string
	name        string
	description string
	options     CliOptions
	runFunc     RunFunc
	subCommands []*SubCommand
	args        cobra.PositionalArgs
	cmd         *cobra.Command
}
type RunFunc func(basename string) error
type Option func(*App)

var (
	cfgFile string
)

func WithOption(opt CliOptions) Option {
	return func(app *App) {
		app.options = opt
	}
}

func WithDescription(desc string) Option {
	return func(app *App) {
		app.description = desc
	}
}
func NewApp(name string, basename string, opts ...Option) *App {
	a := &App{
		basename: basename,
		name:     name,
	}
	for _, o := range opts {
		o(a)
	}

	a.buildCommand()
	return a

}
func (a *App) Run() {
	//fmt.Println(a.basename)
	//fmt.Println(a.cmd.Args)
	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}

	//flagSet := a.options.Flags()
	//fmt.Println(a.basename)

	//fmt.Println("app run func is running ")

}
func runhelp(cmd *cobra.Command, arg []string) {
	cmd.Help()
}
func (a *App) buildCommand() {
	cmd := &cobra.Command{
		Use:   FormatBaseName(a.basename),
		Short: a.name,
		Long:  a.description,
		Args:  a.args,
	}

	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Flags().SortFlags = true
	//cmd.Flags().StringVarP(&cfgFile,"config","c","","config file path")
	if a.runFunc != nil {
		cmd.RunE = a.runCommand
	}
	if a.options != nil {
		optionsFlagSets := a.options.Flags()
		fs := cmd.Flags()
		fs.AddFlagSet(optionsFlagSets)
	}

	a.cmd = cmd
}

/*func (a *App)configInit(){

	if cfgFile != ""{
		viper.SetConfigFile(cfgFile)
	}else {
		fmt.Printf("-----")
		viper.AddConfigPath(".")
		viper.AddConfigPath("C:\\Users\\jimmy\\GolandProjects\\jimmycmd")
		viper.SetConfigName("test")
	}

	viper.SetConfigType("json")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CMDCTL")
	replacer :=strings.NewReplacer(".","_")
	viper.SetEnvKeyReplacer(replacer)
	viper.ReadInConfig()
	//var c Pathconfig
	err := viper.Unmarshal(&c)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}
	//fmt.Printf("use config:%s\n", viper.ConfigFileUsed())
	//fmt.Printf("config_path:%s\n", c.Path)
}*/
func (a *App) runCommand(cmd *cobra.Command, args []string) error {

	if a.runFunc != nil {
		return a.runFunc(a.basename)
	}

	//fmt.Println("runcommand is running ")
	return nil
}

// WithDefaultValidArgs set default validation function to valid non-flag arguments.
func WithDefaultValidArgs() Option {
	return func(a *App) {
		a.args = func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		}
	}
}

// WithRunFunc is used to set the application startup callback function option.
func WithRunFunc(run RunFunc) Option {
	return func(a *App) {
		a.runFunc = run
	}
}
