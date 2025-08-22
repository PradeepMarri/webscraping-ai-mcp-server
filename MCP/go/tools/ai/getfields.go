package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/webscraping-ai/mcp-server/config"
	"github.com/webscraping-ai/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetfieldsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["url"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("url=%v", val))
		}
		if val, ok := args["fields"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields=%v", val))
		}
		if val, ok := args["headers"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("headers=%v", val))
		}
		if val, ok := args["timeout"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("timeout=%v", val))
		}
		if val, ok := args["js"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("js=%v", val))
		}
		if val, ok := args["js_timeout"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("js_timeout=%v", val))
		}
		if val, ok := args["wait_for"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("wait_for=%v", val))
		}
		if val, ok := args["proxy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("proxy=%v", val))
		}
		if val, ok := args["country"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("country=%v", val))
		}
		if val, ok := args["custom_proxy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("custom_proxy=%v", val))
		}
		if val, ok := args["device"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("device=%v", val))
		}
		if val, ok := args["error_on_404"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("error_on_404=%v", val))
		}
		if val, ok := args["error_on_redirect"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("error_on_redirect=%v", val))
		}
		if val, ok := args["js_script"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("js_script=%v", val))
		}
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("api_key=%s", cfg.APIKey))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/ai/fields%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			// API key already added to query string
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGetfieldsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_ai_fields",
		mcp.WithDescription("Extract structured data fields from a web page"),
		mcp.WithString("url", mcp.Required(), mcp.Description("URL of the target page.")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Object describing fields to extract from the page and their descriptions")),
		mcp.WithObject("headers", mcp.Description("HTTP headers to pass to the target page. Can be specified either via a nested query parameter (...&headers[One]=value1&headers=[Another]=value2) or as a JSON encoded object (...&headers={\"One\": \"value1\", \"Another\": \"value2\"}).")),
		mcp.WithNumber("timeout", mcp.Description("Maximum web page retrieval time in ms. Increase it in case of timeout errors (10000 by default, maximum is 30000).")),
		mcp.WithBoolean("js", mcp.Description("Execute on-page JavaScript using a headless browser (true by default).")),
		mcp.WithNumber("js_timeout", mcp.Description("Maximum JavaScript rendering time in ms. Increase it in case if you see a loading indicator instead of data on the target page.")),
		mcp.WithString("wait_for", mcp.Description("CSS selector to wait for before returning the page content. Useful for pages with dynamic content loading. Overrides js_timeout.")),
		mcp.WithString("proxy", mcp.Description("Type of proxy, use residential proxies if your site restricts traffic from datacenters (datacenter by default). Note that residential proxy requests are more expensive than datacenter, see the pricing page for details.")),
		mcp.WithString("country", mcp.Description("Country of the proxy to use (US by default).")),
		mcp.WithString("custom_proxy", mcp.Description("Your own proxy URL to use instead of our built-in proxy pool in \"http://user:password@host:port\" format (<a target=\"_blank\" href=\"https://webscraping.ai/proxies/smartproxy\">Smartproxy</a> for example).")),
		mcp.WithString("device", mcp.Description("Type of device emulation.")),
		mcp.WithBoolean("error_on_404", mcp.Description("Return error on 404 HTTP status on the target page (false by default).")),
		mcp.WithBoolean("error_on_redirect", mcp.Description("Return error on redirect on the target page (false by default).")),
		mcp.WithString("js_script", mcp.Description("Custom JavaScript code to execute on the target page.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetfieldsHandler(cfg),
	}
}
