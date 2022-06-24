package logout

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func LogoutCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "logout",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Run: logout")
			fmt.Println("Args: ", strings.Join(args, " "))
		},
	}

	return cmd
}
