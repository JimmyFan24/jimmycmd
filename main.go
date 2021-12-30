package main

import (
	common2 "jimmycmd/internal/common"
	"jimmycmd/internal/template"
)

func main() {
	//conf := common.NewConfig("123","/tmp")
	//opt := common.NewOption(conf)
	//falg := common.NewFlagSet()
	//common.InitFlags(falg)
	//falg

	cmd := template.NewCommand(common2.NewFlagSetFactory())
	cmd.Execute()
	/*if cmd.Execute() != nil{
		os.Exit(1)
	}
	os.Exit(0)
	*/

	//fmt.Println()
}
