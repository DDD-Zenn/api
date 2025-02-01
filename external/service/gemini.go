package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DDD-Zenn/api/external/serviceIF"
	"github.com/google/generative-ai-go/genai"
)

type geminisvc struct {
	Client *genai.Client
}

func NewGeminiService(clt *genai.Client) serviceIF.Gemini {
	return &geminisvc{
		Client: clt,
	}
}

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []string `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
	UsageMetadata struct {
		PromptTokenCount     int `json:"promptTokenCount"`
		CandidatesTokenCount int `json:"candidatesTokenCount"`
		TotalTokenCount      int `json:"totalTokenCount"`
	} `json:"usageMetadata"`
	ModelVersion string `json:"modelVersion"`
}

func (svc *geminisvc) GenerateResponse(tasks string) (summaryResult string, err error) {
	ctx := context.Background()

	role := "あなたはユーザー専用のCopilotです"
	instruction := "以下の質問に対して回答してください"
	prompt := role + "\n" + instruction + tasks
	model := svc.Client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		return "", fmt.Errorf("failed to marshal response: %w", err)
	}

	fmt.Printf("DEBUG: response JSON = %s\n", respBytes)

	var geminiResp GeminiResponse
	err = json.Unmarshal(respBytes, &geminiResp)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(geminiResp.Candidates) > 0 && len(geminiResp.Candidates[0].Content.Parts) > 0 {
		summaryResult = geminiResp.Candidates[0].Content.Parts[0]
	} else {
		return "", fmt.Errorf("no candidates or parts found")
	}

	return summaryResult, nil
}
