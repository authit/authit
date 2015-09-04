package commands

import (
	"github.com/spf13/cobra"

	crypto "github.com/authit/crypto/commands"
	frontend "github.com/authit/frontend/commands"
)

var Commands = []*cobra.Command{
	&cobra.Command{
		Use:   "crypto",
		Short: "Provided by authit/crypto",
	},
	&cobra.Command{
		Use:   "frontend",
		Short: "Provided by authit/frontend",
	},
}

func init() {
	ConfigDefaults(Commands...)
	Commands[0].AddCommand(crypto.Commands...)
	Commands[1].AddCommand(frontend.Commands...)
}
