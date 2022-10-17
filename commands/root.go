package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/canoypa/go-twitter-test/commands/auth"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:  "tw",
		Args: cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			text := strings.Join(args, " ")
			fmt.Println(text)
			tweet(text)
		},
	}

	rootCmd.AddCommand(auth.AuthCmd())
	rootCmd.AddCommand(InitCmd())

	return rootCmd
}

func tweet(text string) {
	consumerKey := viper.GetString("consumer_key")
	consumerSecret := viper.GetString("consumer_secret")
	accessToken := viper.GetString("token")
	accessSecret := viper.GetString("secret")

	config := &oauth1.Config{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
	}
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	tweet, _, err := client.Statuses.Update(text, nil)
	cobra.CheckErr(err)

	url := strings.Join([]string{"https://twitter.com", tweet.User.ScreenName, "status", strconv.FormatInt(tweet.ID, 10)}, "/")
	fmt.Println("Your Tweet was sent: " + url)
}

func init() {
	cobra.OnInitialize(initializeConfig)
}

func initializeConfig() {
	homePath, err := os.UserHomeDir()
	cobra.CheckErr(err)

	configPath := filepath.Join(homePath, ".twcli")
	configName := "hosts"
	configType := "yaml"

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	// if config not found
	if err := viper.ReadInConfig(); err != nil {
		os.MkdirAll(configPath, 0700)
		viper.WriteConfigAs(filepath.Join(configPath, fmt.Sprintf("%s.%s", configName, configType)))
	}
}
