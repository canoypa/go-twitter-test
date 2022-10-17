package logout

import (
	"fmt"

	"github.com/spf13/cobra"
)

func LogoutCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "logout",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("WIP")
		},
	}

	return cmd
}
