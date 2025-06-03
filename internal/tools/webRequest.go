package tools

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterWebRequestTool(s *server.MCPServer) {
	webRequestTool := mcp.NewTool("web_request",
		mcp.WithDescription("Perform basic web requests"),
		mcp.WithString("curl",
			mcp.Required(),
			mcp.Description("The operation to perform (get,post)"),
			mcp.Enum("GET", "POST"),
		),
		mcp.WithString("url",
			mcp.Required(),
			mcp.Description("address of the web request"),
		),
		mcp.WithString("method",
			mcp.Required(),
			mcp.Description("method of the web request (GET, POST, etc.)"),
		),
	)

	// Add the calculator handler
	s.AddTool(webRequestTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Using helper functions for type-safe argument access
		// op, err := request.RequireString("curl")
		// if err != nil {
		// 	return mcp.NewToolResultError(err.Error()), nil
		// }

		x, err := request.RequireString("url")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		y, err := request.RequireString("method")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		var result string
		switch y {
		case "GET":
			result, err = performGetRequest(x)
		case "POST":
			// Perform POST request logic here
		default:
			//Perform GET request logic here
			return mcp.NewToolResultError("unsupported method"), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf(result)), nil
	})
}

func performGetRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error performing GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GET request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}
