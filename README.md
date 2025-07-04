# kaia (Kubernetes AI Agent)

A comprehensive AI-powered troubleshooting platform for Kubernetes environments, built with Go and React. kaia provides intelligent assistance for EKS cluster troubleshooting by combining LLM capabilities with real-time cluster data through MCP (Model Context Protocol) integration.

## 🚀 Features

- **Multi-LLM Support**: Switchable backends between OpenAI (GPT) and Anthropic (Claude)
- **Real-time Cluster Context**: MCP integration for EKS and CloudWatch data
- **Modern Web Interface**: React-based chat UI with markdown rendering
- **Containerized Deployment**: Docker support for both API and web components
- **Kubernetes Native**: Helm charts for easy deployment in EKS clusters
- **CLI Tool**: Command-line interface for API interaction
- **Secure Secret Management**: ExternalSecrets integration with AWS Secrets Manager

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   React Web UI  │    │   Go API Server │    │   MCP Servers   │
│   (Port 5173)   │◄──►│   (Port 8080)   │◄──►│  (EKS/CW/...)  │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                              │
                              ▼
                       ┌─────────────────┐
                       │   LLM Backends  │
                       │ (OpenAI/Claude) │
                       └─────────────────┘
```

## 🛠️ Quick Start

### Prerequisites

- Go 1.22+
- Node.js 18+
- Docker
- Kubernetes cluster (for deployment)
- API keys for OpenAI and/or Anthropic

### Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/smerck/kaia.git
   cd kaia
   ```

2. **Set up API keys**
   ```bash
   export OPENAI_API_KEY=sk-...
   export ANTHROPIC_API_KEY=sk-ant-...
   ```

3. **Configure the backend** (optional)
   ```bash
   echo '{"asker_backend": "openai"}' > config.json
   ```

4. **Start the API server**
   ```bash
   go run ./cmd/kaia-api
   ```

5. **Start the web interface**
   ```bash
   cd web
   npm install
   npm run dev
   ```

6. **Access the application**
   - Web UI: http://localhost:5173
   - API: http://localhost:8080/api/v1/ask

### Using the CLI

```bash
# Build the CLI
go build -o kaia-cli ./cmd/kaia-cli

# Ask a question
./kaia-cli --prompt "Why are my pods crashing?" --cluster "my-eks-cluster"

# Use a specific backend
./kaia-cli --prompt "What's wrong with my cluster?" --cluster "prod-cluster"
```

## 🐳 Docker Deployment

### Build Images

```bash
# Build API server
docker build -f Dockerfile.api -t kaia-api:latest .

# Build web interface
cd web
docker build -t kaia-web:latest .
```

### Run with Docker Compose

```yaml
# docker-compose.yml
version: '3.8'
services:
  api:
    image: kaia-api:latest
    environment:
      - OPENAI_API_KEY=${OPENAI_API_KEY}
      - ANTHROPIC_API_KEY=${ANTHROPIC_API_KEY}
    ports:
      - "8080:8080"

  web:
    image: kaia-web:latest
    ports:
      - "80:80"
    depends_on:
      - api
```

## ☸️ Kubernetes Deployment

### Using Helm

1. **Add the repository**
   ```bash
   helm repo add kaia https://github.com/smerck/kaia/releases
   helm repo update
   ```

2. **Install the chart**
   ```bash
   helm install kaia ./helm/kaia \
     --set api.image=kaia-api:latest \
     --set web.image=kaia-web:latest
   ```

3. **Configure secrets** (using ExternalSecrets)
   ```bash
   # Create the secret in AWS Secrets Manager
   aws secretsmanager create-secret \
     --name eks/kaia/anthropic_api_key \
     --secret-string "sk-ant-..."
   ```

### Manual Deployment

```bash
kubectl apply -f helm/kaia/templates/
```

## 🔧 Development Guide

### Project Structure

```
kaia/
├── api/                    # HTTP handlers and routes
├── cmd/
│   ├── kaia-api/          # API server entrypoint
│   └── kaia-cli/          # CLI tool
├── internal/
│   ├── agent/             # Core agent logic
│   ├── asker/             # LLM backend interfaces
│   ├── claude/            # Legacy Claude client
│   └── mcp/               # MCP client infrastructure
├── web/                   # React frontend
├── helm/                  # Kubernetes manifests
└── pkg/                   # Shared utilities
```

### Adding New LLM Backends

1. **Implement the Asker interface**
   ```go
   // internal/asker/newbackend.go
   type NewBackendAsker struct{}
   
   func (a *NewBackendAsker) Ask(ctx context.Context, prompt, cluster string) (string, error) {
       // Implementation here
   }
   ```

2. **Add to the agent factory**
   ```go
   // internal/agent/agent.go
   case "newbackend":
       askerImpl = asker.NewNewBackendAsker()
   ```

3. **Update configuration**
   ```json
   {"asker_backend": "newbackend"}
   ```

### Extending MCP Integration

1. **Add new data sources**
   ```go
   // internal/mcp/prometheus.go
   type PrometheusClient struct{}
   
   func (c *PrometheusClient) GetMetrics(ctx context.Context, cluster string) (string, error) {
       // Implementation here
   }
   ```

2. **Update the main MCP client**
   ```go
   // internal/mcp/client.go
   func (c *MCPClient) GatherClusterContext(ctx context.Context, clusterName string) (string, error) {
       // Add new data gathering calls
   }
   ```

### Testing

```bash
# Run all tests
go test ./...

# Run specific test
go test ./api -v

# Test the web interface
cd web && npm test
```

## 🔒 Security Considerations

### API Key Management

- **Never commit API keys** to version control
- **Use ExternalSecrets** for Kubernetes deployments
- **Rotate keys regularly** and monitor usage
- **Implement rate limiting** for API endpoints

### Network Security

- **Use HTTPS** in production
- **Implement proper CORS** policies
- **Restrict MCP server access** to trusted networks
- **Use service mesh** for internal communication

### Data Privacy

- **Log sensitive data carefully** (avoid logging API keys)
- **Implement data retention** policies
- **Consider data residency** requirements
- **Audit access** to cluster data

## 🚀 Production Deployment

### Recommended Setup

1. **Use a production Kubernetes cluster** (EKS, GKE, AKS)
2. **Implement proper monitoring** (Prometheus, Grafana)
3. **Set up logging** (ELK stack, CloudWatch)
4. **Configure auto-scaling** for API servers
5. **Use a CDN** for web assets
6. **Implement proper backup** strategies

### Performance Optimization

- **Enable caching** for frequently accessed data
- **Use connection pooling** for database connections
- **Implement request batching** for MCP calls
- **Optimize container images** (multi-stage builds)
- **Use resource limits** in Kubernetes

### Monitoring and Observability

```yaml
# Example Prometheus metrics
api_requests_total{endpoint="/api/v1/ask",status="200"}
api_request_duration_seconds{endpoint="/api/v1/ask"}
mcp_data_gathering_duration_seconds{source="eks"}
llm_response_time_seconds{backend="openai"}
```

## 🤝 Contributing

1. **Fork the repository**
2. **Create a feature branch** (`git checkout -b feature/amazing-feature`)
3. **Commit your changes** (`git commit -m 'Add amazing feature'`)
4. **Push to the branch** (`git push origin feature/amazing-feature`)
5. **Open a Pull Request**

### Development Guidelines

- **Follow Go conventions** (gofmt, golint)
- **Write tests** for new functionality
- **Update documentation** for API changes
- **Use conventional commits** for commit messages
- **Add examples** for new features

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- **Anthropic** for Claude API
- **OpenAI** for GPT API
- **Microsoft** for MCP specification
- **Kubernetes** community for inspiration
- **React** and **Vite** for the frontend framework

## 📞 Support

- **Issues**: [GitHub Issues](https://github.com/smerck/kaia/issues)
- **Discussions**: [GitHub Discussions](https://github.com/smerck/kaia/discussions)
- **Documentation**: [Wiki](https://github.com/smerck/kaia/wiki)

## 🔮 Roadmap

- [ ] **Prometheus MCP integration**
- [ ] **Custom MCP server support**
- [ ] **Multi-cluster management**
- [ ] **Advanced analytics dashboard**
- [ ] **Plugin system for custom data sources**
- [ ] **Integration with popular monitoring tools**
- [ ] **Mobile application**
- [ ] **Voice interface support**

---

**Made with ❤️ for the Kubernetes community**
