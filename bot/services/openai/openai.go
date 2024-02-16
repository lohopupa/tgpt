package services

import (
	"bot/config"
	"context"
	openai "github.com/sashabaranov/go-openai"
)

type OpenAiClient struct {
	client openai.Client
}

func CreateClient(config config.OpenAIConfig) (OpenAiClient, error) {
	openAIConfig := openai.DefaultConfig(config.Token)
	openAIConfig.BaseURL = config.BaseURL

	client := openai.NewClientWithConfig(openAIConfig)

	return OpenAiClient{
		client: *client,
	}, nil

}

func (client OpenAiClient) Query(query string, history []string) (string, error) {
	resp, err := client.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
