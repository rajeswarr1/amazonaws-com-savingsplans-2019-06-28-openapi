package main

import (
	"github.com/aws-savings-plans/mcp-server/config"
	"github.com/aws-savings-plans/mcp-server/models"
	tools_describesavingsplansofferingrates "github.com/aws-savings-plans/mcp-server/tools/describesavingsplansofferingrates"
	tools_listtagsforresource "github.com/aws-savings-plans/mcp-server/tools/listtagsforresource"
	tools_tagresource "github.com/aws-savings-plans/mcp-server/tools/tagresource"
	tools_createsavingsplan "github.com/aws-savings-plans/mcp-server/tools/createsavingsplan"
	tools_describesavingsplanrates "github.com/aws-savings-plans/mcp-server/tools/describesavingsplanrates"
	tools_untagresource "github.com/aws-savings-plans/mcp-server/tools/untagresource"
	tools_describesavingsplans "github.com/aws-savings-plans/mcp-server/tools/describesavingsplans"
	tools_describesavingsplansofferings "github.com/aws-savings-plans/mcp-server/tools/describesavingsplansofferings"
	tools_deletequeuedsavingsplan "github.com/aws-savings-plans/mcp-server/tools/deletequeuedsavingsplan"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_describesavingsplansofferingrates.CreateDescribesavingsplansofferingratesTool(cfg),
		tools_listtagsforresource.CreateListtagsforresourceTool(cfg),
		tools_tagresource.CreateTagresourceTool(cfg),
		tools_createsavingsplan.CreateCreatesavingsplanTool(cfg),
		tools_describesavingsplanrates.CreateDescribesavingsplanratesTool(cfg),
		tools_untagresource.CreateUntagresourceTool(cfg),
		tools_describesavingsplans.CreateDescribesavingsplansTool(cfg),
		tools_describesavingsplansofferings.CreateDescribesavingsplansofferingsTool(cfg),
		tools_deletequeuedsavingsplan.CreateDeletequeuedsavingsplanTool(cfg),
	}
}
