package command

import "github.com/spf13/cobra"

// AddCommands add available commands to the speicifed command
func AddCommands(cmd *cobra.Command, cli *ManagerCli) {
	cmd.AddCommand(
		NewScreenCommand(cli),
		NewGSheetsCommand(cli),
	)
}
