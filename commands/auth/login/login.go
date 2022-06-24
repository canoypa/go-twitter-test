package login

import (
	"fmt"
	"os"

	"github.com/canoypa/go-twitter-test/core"
	"github.com/canoypa/go-twitter-test/utils"
	"github.com/dghubble/oauth1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func LoginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "login",
		Run: func(cmd *cobra.Command, args []string) {
			token := getAccessToken()

			viper.Set("token", token.Token)
			viper.Set("secret", token.TokenSecret)
			viper.WriteConfig()

			fmt.Println("Successfully logged in.")
		},
	}

	return cmd
}

func getAccessToken() *oauth1.Token {
	oauthConfig := core.OauthConfig()

	requestToken, requestSecret, err := oauthConfig.RequestToken()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	authUrl, err := oauthConfig.AuthorizationURL(requestToken)

	fmt.Println("Visit this URL to get a PIN.")
	fmt.Println(authUrl)
	utils.OpenUrl(authUrl.String())

	var pin string
	fmt.Print("Enter PIN: ")
	fmt.Scanln(&pin)

	accessToken, accessSecret, err := oauthConfig.AccessToken(requestToken, requestSecret, pin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	token := oauth1.NewToken(accessToken, accessSecret)

	return token
}
