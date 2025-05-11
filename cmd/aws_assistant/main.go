package main

import (
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zkfmapf123/aws-questions/internal/handlers"
	"github.com/zkfmapf123/aws-questions/models"
)

var (
	APP_NAME = "aws-assistant"
	VERSION  = "1.0.0"
)

func main() {
	s := server.NewMCPServer(
		APP_NAME,
		VERSION,
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	s.AddTool(initToolChaining(), handlers.AssistantHandler)

	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error : %v\n", err)
	}
}

func initToolChaining() mcp.Tool {
	return mcp.NewTool("aws-assistant",
		mcp.WithDescription("나만의 AWS 비서"),
		mcp.WithString("resource", mcp.Required(), mcp.Description("AWS Resource"), mcp.Enum(models.AWS_ASSISTANT_RESOURCES...)),
		mcp.WithString("command", mcp.Required(), mcp.Description(fmt.Sprintf("동작의 대한 설명 예시) %s", strings.Join(models.AWS_ASSISTANT_COMMANDS, ",")))),
	)
}
