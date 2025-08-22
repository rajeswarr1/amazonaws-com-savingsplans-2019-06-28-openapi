package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/aws-savings-plans/mcp-server/config"
	"github.com/aws-savings-plans/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreatesavingsplanHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/CreateSavingsPlan", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
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
		var result models.CreateSavingsPlanResponse
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

func CreateCreatesavingsplanTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_CreateSavingsPlan",
		mcp.WithDescription("Creates a Savings Plan."),
		mcp.WithString("clientToken", mcp.Description("Input parameter: Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")),
		mcp.WithString("commitment", mcp.Required(), mcp.Description("Input parameter: The hourly commitment, in USD. This is a value between 0.001 and 1 million. You cannot specify more than five digits after the decimal point.")),
		mcp.WithString("purchaseTime", mcp.Description("Input parameter: The time at which to purchase the Savings Plan, in UTC format (YYYY-MM-DDTHH:MM:SSZ).")),
		mcp.WithString("savingsPlanOfferingId", mcp.Required(), mcp.Description("Input parameter: The ID of the offering.")),
		mcp.WithObject("tags", mcp.Description("Input parameter: One or more tags.")),
		mcp.WithString("upfrontPaymentAmount", mcp.Description("Input parameter: The up-front payment amount. This is a whole number between 50 and 99 percent of the total value of the Savings Plan. This parameter is supported only if the payment option is <code>Partial Upfront</code>.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreatesavingsplanHandler(cfg),
	}
}
