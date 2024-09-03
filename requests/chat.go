package requests

import (
	GPT_cli_errors "GPT_cli/utils"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type ChatGPTModel string

// https://platform.openai.com/docs/models
const (
	GPT35Turbo         ChatGPTModel = "gpt-3.5-turbo"
	GPT35Turbo0125     ChatGPTModel = "gpt-3.5-turbo-0125"
	GPT35Turbo1106     ChatGPTModel = "gpt-3.5-turbo-1106"
	GPT35TurboInstruct ChatGPTModel = "gpt-3.5-turbo-instruct"
	GPT4               ChatGPTModel = "gpt-4"
	GPT4o              ChatGPTModel = "gpt-4o"
	DeepseekChat       ChatGPTModel = "deepseek-chat"
)

type ChatGPTModelRole string

const (
	ChatGPTModelRoleUser      ChatGPTModelRole = "user"
	ChatGPTModelRoleSystem    ChatGPTModelRole = "system"
	ChatGPTModelRoleAssistant ChatGPTModelRole = "assistant"
)

type ChatMessage struct {
	Role    ChatGPTModelRole `json:"role"`
	Content string           `json:"content"`
}

type ChatCompletionRequest struct {
	// (Required)
	// ID of the model to use.
	Model ChatGPTModel `json:"model"`

	// Required
	// The messages to generate chat completions for
	Messages []ChatMessage `json:"messages"`

	// (Optional - default: 1)
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.
	// We generally recommend altering this or top_p but not both.
	Temperature float64 `json:"temperature,omitempty"`

	// (Optional - default: 1)
	// An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.
	// We generally recommend altering this or temperature but not both.
	TopP float64 `json:"top_p,omitempty"`

	// (Optional - default: 1)
	// How many chat completion choices to generate for each input message.
	N int `json:"n,omitempty"`

	// (Optional - default: infinite)
	// The maximum number of tokens allowed for the generated answer. By default,
	// the number of tokens the model can return will be (4096 - prompt tokens).
	MaxTokens int `json:"max_tokens,omitempty"`

	// (Optional - default: 0)
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far,
	// increasing the model's likelihood to talk about new topics.
	PresencePenalty float64 `json:"presence_penalty,omitempty"`

	// (Optional - default: 0)
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far,
	// decreasing the model's likelihood to repeat the same line verbatim.
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`

	// (Optional)
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse
	User string `json:"user,omitempty"`
}

type ChatResponseChoice struct {
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"`
	FinishReason string      `json:"finish_reason"`
}

type ChatResponseUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ChatResponse struct {
	ID        string               `json:"id"`
	Object    string               `json:"object"`
	CreatedAt int64                `json:"created_at"`
	Choices   []ChatResponseChoice `json:"choices"`
	Usage     ChatResponseUsage    `json:"usage"`
}

func (c *Client) SimpleSend(ctx context.Context, message string) (*ChatResponse, error) {
	req := &ChatCompletionRequest{
		Model: DeepseekChat,
		Messages: []ChatMessage{
			{
				Role:    ChatGPTModelRoleUser,
				Content: message,
			},
		},
	}
	return c.Send(ctx, req)
}

func (c *Client) Send(ctx context.Context, req *ChatCompletionRequest) (*ChatResponse, error) {
	if err := validate(req); err != nil {
		return nil, err
	}

	reqBytes, err := json.Marshal(req)

	endpoint := "/chat/completions"
	httpReq, err := http.NewRequest("POST", c.config.BaseURL+endpoint, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)

	res, err := c.sendRequest(ctx, httpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var resp ChatResponse
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func validate(req *ChatCompletionRequest) error {
	if len(req.Messages) == 0 {
		return GPT_cli_errors.ErrNoMessages
	}
	// check model allowed
	modelAllowed := false
	allowedModels := []ChatGPTModel{
		GPT35Turbo,
		GPT35Turbo0125,
		GPT35Turbo1106,
		GPT35TurboInstruct,
		GPT35Turbo,
		GPT4o,
		DeepseekChat,
	}
	for _, model := range allowedModels {
		if req.Model == model {
			modelAllowed = true
		}
	}
	if !modelAllowed {
		return GPT_cli_errors.ErrInvalidModel
	}

	// check message available

	availRole := map[ChatGPTModelRole]bool{
		ChatGPTModelRoleUser:      true,
		ChatGPTModelRoleSystem:    true,
		ChatGPTModelRoleAssistant: true,
	}
	for _, msg := range req.Messages {
		if _, ok := availRole[msg.Role]; !ok {
			return GPT_cli_errors.ErrInvalidRole
		}
	}

	if req.Temperature < 0 || req.Temperature > 2 {
		return GPT_cli_errors.ErrInvalidTemperature
	}

	if req.PresencePenalty < -2 || req.PresencePenalty > 2 {
		return GPT_cli_errors.ErrInvalidPresencePenalty
	}

	if req.FrequencyPenalty < -2 || req.FrequencyPenalty > 2 {
		return GPT_cli_errors.ErrInvalidFrequencyPenalty
	}
	return nil
}
