package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"
)

const API_KEY = "sk-1806fac7fe2c44ae8dc58f0ef3ae8c50"

type SearchProductParams struct {
	SQL      string `json:"sql" jsonschema:"description=SQL for searching products."`
}

type SearchProductsResponse struct {
	Success  bool             `json:"success"`
	Message  string           `json:"message"`
	Products []map[string]any `json:"products"`
}

func SearchProductFunc(ctx context.Context, params *SearchProductParams) (SearchProductsResponse, error) {
	fmt.Printf("1. LLM已生成sql,传入了SearchProductFunc并准备与数据库交互: %+v\n", *params)
	// 具体的调用逻辑
	// 如：gorm.find(params.SQL, &products)
	// 假装是查到的数据
	products := []map[string]any{
		{"id": 1, "name": "手机壳1", "color": "红色", "price": 50, "comment_num": 120, "is_deleted": false},
		{"id": 2, "name": "2手机壳", "color": "红色", "price": 40, "comment_num": 130, "is_deleted": false},
	}
	resp := SearchProductsResponse{
		Success:  true,
		Message:  "查询成功",
		Products: products,
	}
	return resp, nil
}

func init() {
	// 工具定义
	searchProductTool, err := utils.InferTool("search_product", "基于用户需求进行商品搜索", SearchProductFunc)
	// 工具数组
	tools := []tool.BaseTool{
		searchProductTool,
	}
	// 工具信息数组
	toolInfos := make([]*schema.ToolInfo, 0, len(tools))
	for _, tool := range tools {
		info, err := tool.Info(ctx)
		toolInfos = append(toolInfos, info)
	}
	// LLM对象
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{...})
	// 将工具信息与LLM绑定
	err = chatModel.BindTools(toolInfos)
	...
}


//func main() {
//	template := prompt.FromMessages(schema.FString,
//		schema.SystemMessage("你是一个{role}。你需要用{style}的语气回答问题。"),
//		schema.MessagesPlaceholder("chat_history", true),
//		schema.UserMessage("问题: {question}"),
//	)
//	messages, err := template.Format(context.Background(), map[string]any{
//		"role":     "程序员鼓励师",
//		"style":    "积极、温暖且专业",
//		"question": "我的代码一直报错，感觉好沮丧，该怎么办？",
//		"chat_history": []*schema.Message{
//			schema.UserMessage("你好"),
//			schema.AssistantMessage("嘿！我是你的程序员鼓励师！有什么我可以帮你的吗？", nil),
//		},
//	})
//
//	ctx := context.Background()
//	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
//		Model:   "deepseek-reasoner",
//		APIKey:  API_KEY,
//		BaseURL: "https://api.deepseek.com",
//	})
//	if err != nil {
//		log.Printf("ChatModel创建失败: %v", err)
//	}
//
//	result, err := chatModel.Generate(ctx, messages)
//	if err != nil {
//		log.Printf("生成失败: %v", err)
//	} else {
//		log.Printf("生成结果: %s", result.Content)
//	}
//}
