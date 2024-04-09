package services

//
//const BASE_URL = "https://openkey.cloud/v1"
//
//type ChatGPT struct {
//	SecretKey string
//}
//
//func NewChatGPT(secretKey string) *ChatGPT {
//	return &ChatGPT{SecretKey: secretKey}
//}
//
//func (c *ChatGPT) Chat(ctx context.Context, extend string, in []*chat.ChatMessage, model string, maxTokens int64) (*openai.ChatCompletionStream, error) {
//	config := openai.DefaultConfig(c.SecretKey)
//	config.BaseURL = BASE_URL
//	client := openai.NewClientWithConfig(config)
//	content := fmt.Sprintf("%s\n样本数据如下:\n%s,请根据我提供的样本数据进行回答,可以适当补充一些职业", aiPlatform.AssistantContent, extend)
//	message := []openai.ChatCompletionMessage{
//		{
//			Role:    openai.ChatMessageRoleUser,
//			Content: content,
//		},
//	}
//
//	for _, v := range in {
//		message = append(message, openai.ChatCompletionMessage{
//			Role:    openai.ChatMessageRoleUser,
//			Content: v.Question,
//		})
//		if v.Content != "" {
//			message = append(message, openai.ChatCompletionMessage{
//				Role:    openai.ChatMessageRoleAssistant,
//				Content: v.Content,
//			})
//		}
//	}
//
//	logx.Info("=============", message)
//
//	stream, err := client.CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
//		Model:     model,
//		Messages:  message,
//		Stream:    true,
//		MaxTokens: int(maxTokens),
//	})
//	if err != nil {
//		return nil, err
//	}
//	return stream, nil
//}
