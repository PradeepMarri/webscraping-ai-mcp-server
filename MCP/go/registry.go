package main

import (
	"github.com/webscraping-ai/mcp-server/config"
	"github.com/webscraping-ai/mcp-server/models"
	tools_ai "github.com/webscraping-ai/mcp-server/tools/ai"
	tools_selected_html "github.com/webscraping-ai/mcp-server/tools/selected_html"
	tools_text "github.com/webscraping-ai/mcp-server/tools/text"
	tools_account "github.com/webscraping-ai/mcp-server/tools/account"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_ai.CreateGetfieldsTool(cfg),
		tools_selected_html.CreateGetselectedmultipleTool(cfg),
		tools_text.CreateGettextTool(cfg),
		tools_account.CreateAccountTool(cfg),
	}
}
