package main

import (
	"context"
	"github.com/tmc/langchaingo/llms/openai"
	"log"
)

const API_KEY = "sk-1806fac7fe2c44ae8dc58f0ef3ae8c50"

func main() {
	llm, err := openai.New(
		openai.WithModel("deepseek-reasoner"),
		openai.WithToken(API_KEY),
		openai.WithBaseURL("https://api.deepseek.com"),
	)
	if err != nil {
		log.Printf("Error creating LLM: %v", err)
	}

	ctx := context.Background()
	prompt := "你好啊，你是什么模型？"
	response, err := llm.Call(ctx, prompt)
	if err != nil {
		log.Printf("Error calling LLM: %v", err)
	} else {
		log.Printf("LLM response:\n%s", response)
	}
}
