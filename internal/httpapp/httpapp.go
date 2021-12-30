package httpapp

import (
	"fmt"
	"github.com/spf13/viper"
	"jimmycmd/internal/httpapp/options"
	"jimmycmd/pkg/app"
)

func NewApp(name string) *app.App {
	opt := options.NewOptions()
	application := app.NewApp(
		"httpapp", "httpapp_basename",
		app.WithRunFunc(run(opt)),
		app.WithDefaultValidArgs(),
		app.WithDescription("httpappdescription"),
		app.WithOption(opt))
	return application
}
func run(opt *options.Options) app.RunFunc {

	return func(basename string) error {

		if opt.ServerOptions.Config != "" {
			conf := configInit(opt.ServerOptions.Config)
			fmt.Println(conf.ServerName)
		}
		return nil
	}
}

var config Config

func configInit(conf string) (c Config) {

	viper.SetConfigFile(conf)
	viper.ReadInConfig()
	//var c Pathconfig
	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}
	return config

}
