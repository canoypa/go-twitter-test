package auth

import (
	"github.com/canoypa/go-twitter-test/commands/auth/login"
	"github.com/canoypa/go-twitter-test/commands/auth/logout"
	"github.com/spf13/cobra"
)

func AuthCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "auth <command>",
	}

	cmd.AddCommand(login.LoginCmd())
	cmd.AddCommand(logout.LogoutCmd())

	return cmd
}
