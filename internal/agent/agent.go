package agent

import (
	"context"
	"fmt"

	"github.com/smerck/kaia/internal/asker"
	"github.com/smerck/kaia/internal/mcp"
)

type Agent struct {
	asker     asker.Asker
	mcpClient *mcp.MCPClient
}

func New(askerBackend string) *Agent {
	var askerImpl asker.Asker
	switch askerBackend {
	case "claude":
		askerImpl = asker.NewClaudeAsker()
	case "openai":
		fallthrough
	default:
		askerImpl = asker.NewOpenAIAsker()
	}

	return &Agent{
		asker:     askerImpl,
		mcpClient: mcp.NewMCPClient(),
	}
}

func (a *Agent) Ask(ctx context.Context, prompt, cluster string) (string, error) {
	// If a cluster is specified, gather context from MCP servers
	if cluster != "" {
		context, err := a.mcpClient.GatherClusterContext(ctx, cluster)
		if err != nil {
			// Log error but continue without context
			fmt.Printf("Warning: Could not gather cluster context: %v\n", err)
		} else {
			// Enhance the prompt with cluster context
			enhancedPrompt := fmt.Sprintf(`
Cluster Context:
%s

User Question: %s

Please provide a helpful response based on the cluster context above and your knowledge of Kubernetes troubleshooting.
`, context, prompt)
			return a.asker.Ask(ctx, enhancedPrompt, cluster)
		}
	}

	// If no cluster specified or context gathering failed, use original prompt
	return a.asker.Ask(ctx, prompt, cluster)
}
