package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Account represents the Account schema from the OpenAPI specification
type Account struct {
	Email string `json:"email,omitempty"` // Your account email
	Remaining_api_calls int `json:"remaining_api_calls,omitempty"` // Remaining API credits quota
	Remaining_concurrency int `json:"remaining_concurrency,omitempty"` // Remaining concurrent requests
	Resets_at int `json:"resets_at,omitempty"` // Next billing cycle start time (UNIX timestamp)
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Body string `json:"body,omitempty"` // Target page response body
	Message string `json:"message,omitempty"` // Error description
	Status_code int `json:"status_code,omitempty"` // Target page response HTTP status code (403, 500, etc)
	Status_message string `json:"status_message,omitempty"` // Target page response HTTP status message
}
