# kaia (Kubernetes AI Agent)

A Go-based API server for troubleshooting EKS clusters using Anthropic's Claude and AWS APIs.

## Features
- Query Claude with troubleshooting prompts
- (Planned) Integrate with AWS EKS, CloudWatch, and Prometheus

## Getting Started

1. **Clone the repo**
2. **Set up Anthropic API key** (for Claude):
   - Export `ANTHROPIC_API_KEY` in your environment
3. **Run the API server:**
   ```sh
   go run ./cmd/kaia-api
   ```

## API

- `POST /api/v1/ask` — Ask troubleshooting questions

## Structure
- `cmd/kaia-api/` — Main entrypoint
- `api/` — HTTP handlers
- `internal/` — Core logic (Claude, AWS clients)
- `pkg/` — Shared types/utilities
