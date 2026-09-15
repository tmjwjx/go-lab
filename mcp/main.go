package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// 初始化 MCP 服务器
	s := server.NewMCPServer("AIServer", "1.0.0")

	// 注册计算器工具
	calculatorTool := mcp.NewTool("calculate",
		mcp.WithDescription("执行基本算术运算"),
		mcp.WithString("operation",
			mcp.Required(),
			mcp.Enum("add", "subtract", "multiply", "divide"),
		),
		mcp.WithNumber("x", mcp.Required()),
		mcp.WithNumber("y", mcp.Required()),
	)

	s.AddTool(calculatorTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		op := req.Params.Arguments["operation"].(string)
		x := req.Params.Arguments["x"].(float64)
		y := req.Params.Arguments["y"].(float64)

		var result float64
		switch op {
		case "add":
			result = x + y
		case "subtract":
			result = x - y
		case "multiply":
			result = x * y
		case "divide":
			if y == 0 {
				return nil, errors.New("division by zero")
			}
			result = x / y
		}
		return mcp.FormatNumberResult(result), nil
	})

	// 注册文件读取资源
	fileResource := mcp.NewResource(
		"file://readme",
		"项目说明",
		mcp.WithResourceDescription("README 文件"),
		mcp.WithMIMEType("text/markdown"),
	)

	s.AddResource(fileResource, func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		content, err := os.ReadFile("README.md")
		if err != nil {
			return nil, err
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "file://readme",
				MIMEType: "text/markdown",
				Text:     string(content),
			},
		}, nil
	})

	// 启动服务器（支持 Stdio 或 HTTP 模式）
	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
