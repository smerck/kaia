# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial project setup with Go API server and React web interface
- Multi-LLM backend support (OpenAI and Claude)
- MCP client infrastructure for EKS and CloudWatch integration
- Docker containerization for both API and web components
- Helm charts for Kubernetes deployment
- CLI tool for API interaction
- Markdown rendering in web interface
- ExternalSecrets integration for secure secret management
- Comprehensive documentation and contributing guidelines

### Changed
- N/A

### Deprecated
- N/A

### Removed
- N/A

### Fixed
- N/A

### Security
- N/A

## [0.1.0] - 2025-07-04

### Added
- **Core Architecture**: Go API server with switchable LLM backends
- **Web Interface**: React-based chat UI with markdown support
- **MCP Integration**: Infrastructure for EKS and CloudWatch data gathering
- **Containerization**: Docker support for both components
- **Kubernetes Deployment**: Helm charts with ExternalSecrets
- **CLI Tool**: Command-line interface for API interaction
- **Documentation**: Comprehensive README, contributing guidelines, and changelog
- **License**: MIT License for open source distribution

### Features
- **Multi-LLM Support**: Toggle between OpenAI and Claude backends
- **Real-time Context**: MCP integration for cluster troubleshooting
- **Modern UI**: Responsive web interface with real-time chat
- **Secure Secrets**: AWS Secrets Manager integration via ExternalSecrets
- **Production Ready**: Docker images and Kubernetes manifests

---

## Version History

- **0.1.0**: Initial release with core functionality
- **Unreleased**: Future development and improvements

## Release Notes

### Version 0.1.0
This is the initial release of kaia, providing a solid foundation for AI-powered Kubernetes troubleshooting. The project includes:

- **Backend**: Go API server with modular LLM integration
- **Frontend**: React web interface with markdown rendering
- **Infrastructure**: Docker containers and Kubernetes deployment
- **Security**: Secure secret management and best practices
- **Documentation**: Comprehensive guides for users and contributors

### Future Releases
Planned features for upcoming releases:

- **v0.2.0**: Enhanced MCP integration with real AWS services
- **v0.3.0**: Prometheus metrics and advanced analytics
- **v1.0.0**: Production-ready with enterprise features

---

For detailed information about each release, see the [GitHub releases page](https://github.com/smerck/kaia/releases). 