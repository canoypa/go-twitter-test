package core

import (
	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
)

var (
	consumerKey    = "OT6LqaU9SrQGljljsV9FJJcc0"
	consumerSecret = "njbfmZ5o24bDPhAyTZP3Bt195qoj5AvPz3FSuTQ45trWzXVsmD"
)

func OauthConfig() *oauth1.Config {
	config := &oauth1.Config{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		CallbackURL:    "oob",
		Endpoint:       twitter.AuthenticateEndpoint,
	}

	return config
}
