package routes

import (
	"context"
	"fmt"

	huma "github.com/danielgtaylor/huma/v2"

	config "contact-me/src/config"
	Infra "contact-me/src/infrastructure"
	RestApi "contact-me/src/infrastructure/rest_api"
	Slack "contact-me/src/services/slack"
)

type ContactRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Message     string `json:"message"`
}

func PrepareForSlack(contactRequest ContactRequest) string {
	return fmt.Sprintf(`
## A visitor has requested to be contacted! :tada:
- Name: %s
- Email: %s
- Phone Number: %s
- Message: %s
	`, contactRequest.Name, contactRequest.Email, contactRequest.PhoneNumber, contactRequest.Message)
}

func Register(infrastructure *Infra.Infrastructure, slackClient Slack.Client, config config.Config) {
	api := infrastructure.Server.Api

	RestApi.UseRoute(
		api,
		huma.Operation{
			Method:      "POST",
			Path:        "/contact-requests",
			Description: "Send a message",
			Tags:        []string{"Contact Request"},
		},
		func(ctx context.Context, request *RestApi.RequestBody[ContactRequest]) (*RestApi.NonNullableResponse[string], error) {
			channelId, channelSearchError := slackClient.FindChannelIdByName(config.Slack.ChannelName)

			if channelSearchError != nil {
				response := RestApi.NonNullableJsonResponse("")

				return &response, channelSearchError
			}

			err := slackClient.SendMessage(channelId, PrepareForSlack(request.Body), true)

			if err != nil {
				fmt.Println(err.Error())
			}

			response := RestApi.NonNullableJsonResponse("Ok")

			return &response, err
		},
	)
}
