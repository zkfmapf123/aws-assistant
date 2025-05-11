package handlers

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/zkfmapf123/aws-questions/internal/helpers"
	"github.com/zkfmapf123/aws-questions/internal/vendor"
	"github.com/zkfmapf123/aws-questions/models"
)

func getParameter(req mcp.CallToolRequest, value string) string {
	arg, exists := req.Params.Arguments[value]
	if !exists {
		return ""
	}

	str, ok := arg.(string)
	if !ok {
		return ""
	}

	return str
}

func AssistantHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	cfg, err := vendor.NewAWSConfig()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("AWS 설정 에러: %v", err)), err
	}

	// mcp requests
	mcpRequest := models.MCPRequestParmas{
		Resource: getParameter(req, "resource"),
		Command:  getParameter(req, "command"),
		Item:     getParameter(req, "item"),
	}

	// interface 강제
	assistantMapper := map[string]models.Commands{
		"s3":       vendor.NewS3Client(cfg),
		"iam-user": vendor.NewIAMUserClien(cfg),
	}

	// mcp response use commands
	switch mcpRequest.Command {
	case "list":
		return helpers.ResponseMcpArray(assistantMapper[mcpRequest.Resource].List())

	case "search":
		return helpers.ResponseMcpArray(assistantMapper[mcpRequest.Resource].Search(mcpRequest.Item))
	}

	return mcp.NewToolResultText("지원하지 않는 명령어입니다."), nil
}
