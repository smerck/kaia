# Contributing to kaia

Thank you for your interest in contributing to kaia! This document provides guidelines and information for contributors.

## 🤝 How to Contribute

### Reporting Issues

Before creating an issue, please:

1. **Search existing issues** to avoid duplicates
2. **Use the issue templates** when available
3. **Provide detailed information**:
   - Environment details (OS, Go version, Node version)
   - Steps to reproduce
   - Expected vs actual behavior
   - Logs and error messages

### Feature Requests

When requesting features:

1. **Describe the use case** clearly
2. **Explain the expected benefits**
3. **Consider implementation complexity**
4. **Check if it aligns with project goals**

### Code Contributions

#### Getting Started

1. **Fork the repository**
2. **Clone your fork**
   ```bash
   git clone https://github.com/your-username/kaia.git
   cd kaia
   ```
3. **Create a feature branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```
4. **Set up development environment**
   ```bash
   # Install Go dependencies
   go mod download
   
   # Install Node.js dependencies
   cd web && npm install
   ```

#### Development Workflow

1. **Make your changes**
2. **Write tests** for new functionality
3. **Update documentation** as needed
4. **Run tests locally**
   ```bash
   # Go tests
   go test ./...
   
   # Web tests
   cd web && npm test
   ```
5. **Commit your changes**
   ```bash
   git add .
   git commit -m "feat: add new feature description"
   ```
6. **Push to your fork**
   ```bash
   git push origin feature/your-feature-name
   ```
7. **Create a Pull Request**

## 📋 Development Guidelines

### Code Style

#### Go Code

- **Follow Go conventions** and use `gofmt`
- **Use meaningful variable names**
- **Add comments** for exported functions
- **Keep functions small** and focused
- **Handle errors** explicitly

```go
// Good example
func GetClusterInfo(ctx context.Context, clusterName string) (*ClusterInfo, error) {
    if clusterName == "" {
        return nil, errors.New("cluster name is required")
    }
    
    // Implementation here
    return info, nil
}
```

#### React/TypeScript Code

- **Use TypeScript** for type safety
- **Follow React best practices**
- **Use functional components** with hooks
- **Add proper error boundaries**

```typescript
// Good example
interface MessageProps {
  sender: 'user' | 'agent';
  text: string;
  backend?: string;
}

const Message: React.FC<MessageProps> = ({ sender, text, backend }) => {
  return (
    <div className={`message ${sender}`}>
      <ReactMarkdown>{text}</ReactMarkdown>
      {backend && <small>via {backend}</small>}
    </div>
  );
};
```

### Testing

#### Go Tests

- **Write unit tests** for all new functions
- **Use table-driven tests** for multiple scenarios
- **Mock external dependencies**
- **Test error conditions**

```go
func TestGetClusterInfo(t *testing.T) {
    tests := []struct {
        name        string
        clusterName string
        wantErr     bool
    }{
        {"valid cluster", "test-cluster", false},
        {"empty cluster", "", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            _, err := GetClusterInfo(context.Background(), tt.clusterName)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetClusterInfo() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

#### React Tests

- **Test component rendering**
- **Test user interactions**
- **Mock API calls**
- **Test error states**

```typescript
import { render, screen, fireEvent } from '@testing-library/react';
import { App } from './App';

test('sends message when form is submitted', async () => {
  render(<App />);
  
  const input = screen.getByPlaceholderText('Type your question...');
  const button = screen.getByText('Send');
  
  fireEvent.change(input, { target: { value: 'Test message' } });
  fireEvent.click(button);
  
  expect(screen.getByText('Test message')).toBeInTheDocument();
});
```

### Documentation

- **Update README.md** for user-facing changes
- **Add inline comments** for complex logic
- **Update API documentation** for new endpoints
- **Include examples** for new features

### Commit Messages

Use [Conventional Commits](https://www.conventionalcommits.org/) format:

```
type(scope): description

[optional body]

[optional footer]
```

Examples:
- `feat(api): add new endpoint for cluster metrics`
- `fix(web): resolve markdown rendering issue`
- `docs(readme): update installation instructions`
- `test(agent): add unit tests for MCP client`

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes
- `refactor`: Code refactoring
- `test`: Adding tests
- `chore`: Maintenance tasks

## 🔧 Development Environment

### Prerequisites

- **Go 1.22+**
- **Node.js 18+**
- **Docker** (for containerization)
- **Git**

### Local Setup

1. **Install dependencies**
   ```bash
   # Go dependencies
   go mod download
   
   # Node.js dependencies
   cd web && npm install
   ```

2. **Set up environment variables**
   ```bash
   export OPENAI_API_KEY=sk-...
   export ANTHROPIC_API_KEY=sk-ant-...
   ```

3. **Run development servers**
   ```bash
   # API server
   go run ./cmd/kaia-api
   
   # Web interface (in another terminal)
   cd web && npm run dev
   ```

### Debugging

#### Go Debugging

- **Use `log.Printf`** for debugging
- **Use Delve** for step-through debugging
- **Check Go modules** with `go mod tidy`

#### React Debugging

- **Use React DevTools** browser extension
- **Use browser developer tools**
- **Add `console.log`** statements
- **Use React error boundaries**

## 🚀 Release Process

### Versioning

We use [Semantic Versioning](https://semver.org/):

- **MAJOR**: Breaking changes
- **MINOR**: New features (backward compatible)
- **PATCH**: Bug fixes (backward compatible)

### Release Checklist

Before releasing:

- [ ] **All tests pass**
- [ ] **Documentation is updated**
- [ ] **Changelog is updated**
- [ ] **Version is bumped**
- [ ] **Docker images are built**
- [ ] **Helm charts are updated**

## 📞 Getting Help

- **GitHub Issues**: For bugs and feature requests
- **GitHub Discussions**: For questions and general discussion
- **Documentation**: Check the README and wiki
- **Code of Conduct**: Be respectful and inclusive

## 🎯 Project Goals

kaia aims to:

- **Simplify Kubernetes troubleshooting** with AI assistance
- **Provide real-time cluster insights** through MCP integration
- **Support multiple LLM backends** for flexibility
- **Maintain high code quality** and reliability
- **Foster an open and inclusive community**

Thank you for contributing to kaia! 🚀 