package commands

import (
	"github.com/spf13/cobra"

	crypto "github.com/authit/crypto/commands"
)

var Commands = []*cobra.Command{
	&cobra.Command{
		Use:   "crypto",
		Short: "Provided by authit/crypto",
		Run: func(cmd *cobra.Command, args []string) {
			crypto.ConfigBindFlags(cmd)
			// sample.Run()
		},
	},
}

func init() {
	ConfigDefaults(Commands...)
	Commands[0].AddCommand(crypto.Commands...)
}
