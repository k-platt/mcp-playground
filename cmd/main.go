package main

import (
	"fmt"
	"net/http"

	"github.com/k-platt/mcp-playground/internal/metrics" // Adjust the import path as necessary
	"github.com/k-platt/mcp-playground/internal/tools"   // Import the tools package
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"mcp-playground",
		"1.0.0",
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)

	// Register the calculator tool
	tools.RegisterCalculatorTool(s)
	tools.RegisterWebRequestTool(s)

	go func() {
		http.Handle("/metrics", metrics.Handler())
		if err := http.ListenAndServe(":9090", nil); err != nil {
			fmt.Printf("Metrics server error: %v\n", err)
		}
	}()

	// Start the MCP server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
