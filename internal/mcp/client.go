package mcp

import (
	"context"
	"fmt"
	"log"
)

// MCPClient represents a client for interacting with MCP servers
type MCPClient struct {
	eksClient        *EKSClient
	cloudwatchClient *CloudWatchClient
}

// NewMCPClient creates a new MCP client
func NewMCPClient() *MCPClient {
	return &MCPClient{
		eksClient:        NewEKSClient(),
		cloudwatchClient: NewCloudWatchClient(),
	}
}

// GatherClusterContext collects relevant data about a cluster for troubleshooting
func (c *MCPClient) GatherClusterContext(ctx context.Context, clusterName string) (string, error) {
	log.Printf("[MCP] Gathering context for cluster: %s", clusterName)

	var context string

	// Gather EKS data
	if eksData, err := c.eksClient.GetClusterInfo(ctx, clusterName); err != nil {
		log.Printf("[MCP] Error getting EKS data: %v", err)
	} else {
		context += fmt.Sprintf("EKS Cluster Info:\n%s\n\n", eksData)
	}

	if podData, err := c.eksClient.GetPodStatus(ctx, clusterName); err != nil {
		log.Printf("[MCP] Error getting pod data: %v", err)
	} else {
		context += fmt.Sprintf("Pod Status:\n%s\n\n", podData)
	}

	// Gather CloudWatch data
	if logData, err := c.cloudwatchClient.GetRecentLogs(ctx, clusterName); err != nil {
		log.Printf("[MCP] Error getting CloudWatch logs: %v", err)
	} else {
		context += fmt.Sprintf("Recent Logs:\n%s\n\n", logData)
	}

	if context == "" {
		context = "No cluster context available."
	}

	return context, nil
}
