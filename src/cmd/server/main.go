package main

import (
	Infra "contact-me/src/infrastructure"
	Routes "contact-me/src/rest-api"
	Slack "contact-me/src/services/slack"
	"fmt"

	"contact-me/src/config"
)

func main() {
	fmt.Println("Starting API server...")
	config := config.AppConfig
	infrastructure, err := Infra.UseInfrastructure(config)

	defer infrastructure.Close()

	if err != nil {
		fmt.Println(err.Error())
		panic("Unable to create infrastructure: " + err.Error())
	}

	slackClient := Slack.CreateClient(config.Slack.BotOAuthUserToken)

	Routes.Register(infrastructure, slackClient, config)

	// TODO: Probably need to specify server port here
	infrastructure.Server.Router.Run()
}
