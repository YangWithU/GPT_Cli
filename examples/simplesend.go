package examples

import (
	"GPT_cli/requests"
	"context"
	"encoding/json"
	"log"
)

func main() {
	c, err := requests.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	res, err := c.SimpleSend(ctx, "Hey, Explain GoLang to me in 2 sentences.")
	if err != nil {
		log.Fatal(err)
	}

	a, _ := json.MarshalIndent(res, "", "  ")
	log.Println(string(a))

	res, err = c.Send(ctx, &requests.ChatCompletionRequest{
		Model: requests.DeepseekChat,
		Messages: []requests.ChatMessage{
			{
				Role:    requests.ChatGPTModelRoleSystem,
				Content: "Hey, Explain GoLang to me in 2 sentences.",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	a, _ = json.MarshalIndent(res, "", "  ")
	log.Println(string(a))
}
