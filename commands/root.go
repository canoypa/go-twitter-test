package commands

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/canoypa/go-twitter-test/commands/auth"
	"github.com/canoypa/go-twitter-test/core"
	"github.com/dghubble/oauth1"
	"github.com/g8rswimmer/go-twitter/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "tw",
		Run: func(cmd *cobra.Command, args []string) {
			text, err := cmd.Flags().GetString("text")

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if text == "" {
				cmd.Help()
				return
			}

			fmt.Println(text)

			twitterTest(text)
		},
	}

	cmd.Flags().StringP("text", "t", "", "Tweet text")
	cmd.AddCommand(auth.AuthCmd())

	return cmd
}

type authorize struct {
}

func (a authorize) Add(req *http.Request) {
}

func twitterTest(text string) {
	accessToken := viper.GetString("token")
	accessSecret := viper.GetString("secret")

	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := core.OauthConfig().Client(oauth1.NoContext, token)

	client := &twitter.Client{
		Authorizer: authorize{},
		Client:     httpClient,
		Host:       "https://api.twitter.com",
	}

	request := twitter.CreateTweetRequest{
		Text: text,
	}

	client.CreateTweet(context.Background(), request)
}

func init() {
	configName := "hosts"
	configType := "yaml"
	configPath := core.GetConfigPath()

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		os.MkdirAll(configPath, 0700)
		viper.WriteConfigAs(filepath.Join(configPath, fmt.Sprintf("%s.%s", configName, configType)))
	}
}
