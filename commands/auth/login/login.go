package login

import (
	"fmt"

	"github.com/canoypa/go-twitter-test/utils"
	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func LoginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "login",
		Run: func(cmd *cobra.Command, args []string) {
			token, secret := getAccessToken()

			viper.Set("token", token)
			viper.Set("secret", secret)
			viper.WriteConfig()

			fmt.Println("Successfully signed in.")
		},
	}

	return cmd
}

func getAccessToken() (string, string) {
	consumerKey := viper.GetString("consumerKey")
	consumerSecret := viper.GetString("consumerSecret")

	oauthConfig := &oauth1.Config{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		CallbackURL:    "oob",
		Endpoint:       twitter.AuthenticateEndpoint,
	}

	requestToken, requestSecret, err := oauthConfig.RequestToken()
	cobra.CheckErr(err)

	authUrl, err := oauthConfig.AuthorizationURL(requestToken)

	fmt.Println("Visit this URL to get a PIN.")
	fmt.Println(authUrl)
	utils.OpenUrl(authUrl.String())

	pin := utils.Input("Enter PIN")

	accessToken, accessSecret, err := oauthConfig.AccessToken(requestToken, requestSecret, pin)
	cobra.CheckErr(err)

	return accessToken, accessSecret
}
