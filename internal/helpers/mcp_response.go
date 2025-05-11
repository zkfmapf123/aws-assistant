package helpers

import (
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

func ResponseMcpArray(v []string, err error) (*mcp.CallToolResult, error) {
	if err != nil {
		return mcp.NewToolResultText("올바르지 않은 유형입니다."), err
	}

	if len(v) == 0 {
		return mcp.NewToolResultText("값이 없습니다."), nil
	}

	if len(v) > 0 {
		return mcp.NewToolResultText(strings.Join(v, ",")), nil
	}

	return nil, err
}
