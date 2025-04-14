package config

import "os"

type CORSConfig struct {
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	AllowOrigins     []string
}

type SlackConfig struct {
	BotOAuthUserToken string
	ChannelName       string
}

type APIConfig struct {
	Name    string
	Port    string
	Version string
}

type Config struct {
	CORS  CORSConfig
	Slack SlackConfig
	API   APIConfig
}

var AppConfig = Config{
	CORS: CORSConfig{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOrigins:     []string{"http://127.0.0.1:3000", "http://localhost:3000", "https://finnyg.com"},
	},
	Slack: SlackConfig{
		// Requires the following scopes: channels:join, channels:read, chat:write, groups:write
		BotOAuthUserToken: os.Getenv("SLACK_BOT_OAUTH_USER_TOKEN"),
		ChannelName:       "inbox",
	},
	API: APIConfig{
		Port:    os.Getenv("API_PORT"),
		Name:    "Contact Me API",
		Version: "1.0.0",
	},
}
