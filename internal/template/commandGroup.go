package template

import (
	"github.com/spf13/cobra"
)

type CommandGroup struct {
	Message  string
	Commands []*cobra.Command
}
type CommandGroups []CommandGroup

func (cg CommandGroups) Add(c *cobra.Command) {

	for _, group := range cg {
		for _, command := range group.Commands {
			c.AddCommand(command)
		}
	}

}
