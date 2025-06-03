# MCP Playground

MCP Playground is a Go-based application that implements a simple MCP (Multi-Channel Protocol) server with Prometheus metrics integration. This project serves as a demonstration of creating and managing tools within an MCP server.

## Project Structure

```
mcp-playground
├── cmd
│   └── main.go            # Entry point of the application
├── internal
│   ├── metrics
│   │   └── prometheus.go  # Implementation for exposing Prometheus metrics
│   ├── server
│   │   └── mcp_server.go  # MCP server implementation
│   └── tools
│       ├── calculator.go  # Implementation of the calculator tool
│       └── webRequest.go  # Implementation of the web request tool
├── go.mod                  # Module definition
├── go.sum                  # Dependency checksums
└── README.md               # Project documentation
```

## Features

- **MCP Server**: A server that handles tool requests and manages their lifecycle.
- **Calculator Tool**: A tool that performs basic arithmetic operations (add, subtract, multiply, divide).
- **Web Request Tool**: A tool for performing HTTP requests.
- **Prometheus Metrics**: Exposes metrics for monitoring the application's performance.

## Getting Started

To run the application, follow these steps:

### Clone the Repository

1. Clone the repository:
   ```
   git clone <repository-url>
   cd mcp-playground
   ```

### Install Dependencies

2. Install dependencies:
   ```
   go mod tidy
   ```

### Build the Application

3. Build the application:
   ```
   go build ./cmd/main.go
   ```

This will generate an executable file named `main` in the current directory.

### Run the Application

4. Run the application:
   ```
   ./main
   ```

5. Access the Prometheus metrics at `http://localhost:9090/metrics`.

## Using VS Code to Manage MCP Servers

### Listing MCP Servers

1. Open the **Command Palette** in VS Code (`Cmd+Shift+P` on macOS).
2. Search for `MCP: List Servers` and select it.
3. You should see a list of available MCP servers, including this one.

### Starting the MCP Server

1. Open the **Command Palette** in VS Code (`Cmd+Shift+P` on macOS).
2. Search for `MCP: Start Server` and select it.
3. Choose the MCP Playground server from the list.
4. The server will start, and you can monitor its logs in the integrated terminal.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.