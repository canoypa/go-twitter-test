package commands

import (
	"github.com/canoypa/go-twitter-test/commands/auth/login"
	"github.com/canoypa/go-twitter-test/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func InitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "init",
		Run: func(cmd *cobra.Command, args []string) {
			consumerKey := utils.Input("Enter ConsumerKey")
			consumerSecret := utils.Input("Enter ConsumerSecret")

			viper.Set("consumer_key", consumerKey)
			viper.Set("consumer_secret", consumerSecret)

			continueSignIn := utils.Confirm("Continue SignIn?", true)

			if continueSignIn {
				login.LoginCmd().Execute()
			}
		},
	}

	return cmd
}
