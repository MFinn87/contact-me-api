package slack

import (
	"contact-me/src/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	apiKey string
}

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ListChannelsResponse struct {
	Ok       bool      `json:"ok"`
	Channels []Channel `json:"channels"`
}

type CreateSimpleMessageRequest struct {
	Channel      string `json:"channel"`
	Text         string `json:"text"`
	MarkdownText string `json:"markdown_text"`
}

type CreateSimpleMessageResponse struct {
	Ok bool `json:"ok"`
}

func (self *Client) FindChannelIdByName(name string) (string, error) {
	client := &http.Client{}
	data := strings.NewReader("")

	request, requestErr := http.NewRequest("GET", "https://slack.com/api/conversations.list", data)

	if requestErr != nil {
		return "", requestErr
	}

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", self.apiKey))

	response, responseErr := client.Do(request)

	if responseErr != nil {
		return "", responseErr
	}

	defer response.Body.Close()
	body, responseReadErr := io.ReadAll(response.Body)

	if responseReadErr != nil {
		return "", responseReadErr
	}

	slackResponse := ListChannelsResponse{}

	unmarshalErr := json.Unmarshal(body, &slackResponse)

	if unmarshalErr != nil {
		return "", unmarshalErr
	}

	if slackResponse.Ok != true {
		fmt.Println(string(body))

		return "", errors.New("Failed to retrieve Slack Channels")
	}

	channel := utils.Find[Channel](slackResponse.Channels, func(channel Channel, index int32) bool {
		if channel.Name == name {
			return true
		}

		return false
	})

	channelId := ""

	if channel != nil {
		channelId = channel.Id
	}

	return channelId, nil
}

func (self *Client) SendMessage(channelId string, content string, isMarkdown bool) error {
	client := &http.Client{}
	messageRequest := &CreateSimpleMessageRequest{
		Channel:      channelId,
		Text:         "",
		MarkdownText: "",
	}

	// According to Slack docs, MUST be one or the other
	if isMarkdown {
		messageRequest.MarkdownText = content
	} else {
		messageRequest.Text = content
	}

	binaryData, jsonMarshalErr := json.Marshal(messageRequest)

	if jsonMarshalErr != nil {
		return jsonMarshalErr
	}

	textData := string(binaryData)

	data := strings.NewReader(textData)
	request, requestErr := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", data)

	if requestErr != nil {
		return requestErr
	}

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", self.apiKey))

	response, responseErr := client.Do(request)

	if responseErr != nil {
		return responseErr
	}

	defer response.Body.Close()
	body, responseReadErr := io.ReadAll(response.Body)

	if responseReadErr != nil {
		return responseReadErr
	}

	slackResponse := CreateSimpleMessageResponse{}

	unmarshalErr := json.Unmarshal(body, &slackResponse)

	if unmarshalErr != nil {
		return unmarshalErr
	}

	if slackResponse.Ok != true {
		fmt.Println(string(body))

		return errors.New("Failed to retrieve Slack Channels")
	}

	return nil
}

func CreateClient(apiKey string) Client {
	return Client{
		apiKey: apiKey,
	}
}
