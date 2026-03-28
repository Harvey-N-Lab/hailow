# Hailow - AI Agent Configuration Manager

A production-ready repository and CLI tool for managing reusable Claude Code and Roo Code configurations across multiple engineering domains.

## Features

- 🎯 **5 Engineering Domains**: DevOps, Python Backend, JavaScript/TypeScript, Data Engineering, Machine Learning
- 🤖 **5 Agent Roles per Domain**: Researcher, Planner, Architect, Implementer, Reviewer
- 🔧 **Multi-Platform Support**: Works with both Roo Code and Claude Code
- 📦 **Easy Installation**: Simple CLI tool with one-command setup
- 🔄 **Flexible Sourcing**: Pull from public repo, custom repo, or local path
- 🚀 **Production-Ready**: Comprehensive configurations with best practices

## Quick Start

### Install the CLI

**Linux/macOS:**
```bash
curl -sSL https://raw.githubusercontent.com/username/hailow/main/scripts/install.sh | bash
```

**Or download binary directly** from [GitHub Releases](https://github.com/Harvey-N-Lab/hailow/releases/latest)

### Install a Domain

```bash
# Install single domain for Roo Code
hailow install devops-engineer --platform roo

# Install multiple domains
hailow install devops-engineer python-backend-engineer

# Install all domains
hailow install --all

# Install for Claude Code
hailow install python-backend-engineer --platform claude
```

## Available Domains

| Domain | Description | Aliases |
|--------|-------------|---------|
| `devops-engineer` | Infrastructure, CI/CD, deployment, operations | `devops`, `ops` |
| `python-backend-engineer` | Python backend development, APIs, services | `python`, `py` |
| `js-ts-software-engineer` | JavaScript/TypeScript development | `jsts`, `js`, `ts` |
| `data-engineer` | Data pipelines, ETL, data quality | `data`, `de` |
| `machine-learning-engineer` | ML models, training, deployment | `mleng`, `ml` |

## What's Included

Each domain includes:

- **Agents** (5 specialized prompts):
  - `researcher` - Investigates requirements and context
  - `planner` - Breaks down work into actionable tasks
  - `architect` - Designs technical solutions
  - `implementer` - Executes implementation
  - `reviewer` - Validates quality and correctness

- **Rules**: Domain-specific coding standards and best practices
- **Skills**: Recommended skills to install for the domain
- **Commands**: Workflow commands and shortcuts
- **Contexts**: Placeholder for user-specific project context

## Agent Workflow

All domains follow a structured 5-phase workflow:

```
Research → Plan → Architect → Implement → Review
```

See [`AGENT_WORKFLOW.md`](AGENT_WORKFLOW.md) for details.

## CLI Commands

```bash
# Install domains
hailow install <domain>...          # Install specific domains
hailow install --all                # Install all domains
hailow install --dry-run           # Preview installation

# List information
hailow list domains                 # List available domains
hailow list platforms               # List supported platforms
hailow list installed               # Show installed domains

# Manage installations
hailow remove <domain>...           # Remove domains
hailow remove --all                 # Remove all domains
hailow upgrade [domain]...          # Upgrade to latest versions

# Configuration
hailow config show                  # Show current config
hailow config set <key> <value>     # Set config value

# Version
hailow version                      # Show version info
```

## Configuration

The CLI uses `~/.hailow/config.yaml`:

```yaml
# Default source
source:
  type: git
  url: https://github.com/Harvey-N-Lab/hailow
  branch: main

# Default platform
platform: roo

# Installation behavior
install:
  on_conflict: skip        # skip, overwrite, prompt
  include_general: true
  backup: true
```

## Custom Sources

### Use Custom Repository

```bash
# Configure
hailow config set source.url https://github.com/mycompany/configs

# Or use flag
hailow install devops-engineer --source https://github.com/mycompany/configs
```

### Use Local Path

```bash
# Configure
hailow config set source.type local
hailow config set source.path /path/to/configs

# Or use flag
hailow install python-backend-engineer --source /path/to/configs
```

## Examples

### Example 1: DevOps Engineer Setup

```bash
# Install DevOps domain
hailow install devops-engineer --platform roo

# Your workspace now has:
# - .agents/devops-researcher.md
# - .agents/devops-planner.md
# - .agents/devops-architect.md
# - .agents/devops-implementer.md
# - .agents/devops-reviewer.md
# - .agents/rules/ (infrastructure, security, deployment standards)
# - .agents/skills/ (terraform-helper, kubernetes-assistant, etc.)
# - .agents/commands/ (deploy, rollback, health-check)
# - AGENT_WORKFLOW.md (workflow guidance)
```

### Example 2: Full-Stack Team Setup

```bash
# Install frontend and backend domains
hailow install js-ts-software-engineer python-backend-engineer devops-engineer

# Team now has configurations for:
# - Frontend development (React, TypeScript, etc.)
# - Backend development (Python, API design, etc.)
# - DevOps operations (deployment, infrastructure, etc.)
```

### Example 3: Data Science Team

```bash
# Install data and ML domains
hailow install data-engineer machine-learning-engineer --all

# Team has configurations for:
# - Data pipelines and ETL
# - ML model development and deployment
```

## Repository Structure

```
.
├── general/                    # Cross-domain rules
├── devops-engineer/            # DevOps domain
├── python-backend-engineer/    # Python backend domain
├── js-ts-software-engineer/    # JS/TS domain
├── data-engineer/              # Data engineering domain
├── machine-learning-engineer/  # ML engineering domain
├── domains/                    # Domain metadata
├── cmd/hailow/            # CLI source code
├── internal/                   # Internal packages
└── .github/workflows/          # CI/CD
```

See [`REPOSITORY_STRUCTURE.md`](REPOSITORY_STRUCTURE.md) for details.

## Documentation

- [**INSTALL.md**](INSTALL.md) - Detailed installation instructions
- [**UPGRADE.md**](UPGRADE.md) - Upgrade guide
- [**AGENT_WORKFLOW.md**](AGENT_WORKFLOW.md) - Agent collaboration workflow
- [**ARCHITECTURE.md**](ARCHITECTURE.md) - Architecture decisions
- [**CLI_DESIGN.md**](CLI_DESIGN.md) - CLI design and usage
- **Domain Documentation**:
  - [DevOps Engineer Rules](devops-engineer/rules/)
  - [Python Backend Rules](python-backend-engineer/rules/)
  - And more...

## Development

### Build from Source

```bash
# Clone repository
git clone https://github.com/Harvey-N-Lab/hailow
cd hailow

# Build CLI
make build

# Run tests
make test

# Install locally
make install
```

### Requirements

- Go 1.21 or later
- Git (for fetching remote configs)

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](docs/CONTRIBUTING.md) for guidelines.

### Adding a New Domain

1. Create domain directory: `<domain-name>/`
2. Add required subdirectories: `agents/`, `rules/`, `skills/`, `commands/`, `contexts/`
3. Create 5 agent files with domain-specific prompts
4. Add domain to `domains/domains.yaml`
5. Update documentation
6. Submit pull request

## Roadmap

- [ ] Shell completion (bash, zsh, fish)
- [ ] Interactive domain selection
- [ ] Domain templates for creating custom domains
- [ ] Validation of custom domains
- [ ] Diff view for upgrades
- [ ] Backup/restore functionality
- [ ] Web UI for browsing domains

## License

MIT License - see [LICENSE](LICENSE) for details.

## Support

- **Issues**: [GitHub Issues](https://github.com/Harvey-N-Lab/hailow/issues)
- **Discussions**: [GitHub Discussions](https://github.com/Harvey-N-Lab/hailow/discussions)
- **Documentation**: [Wiki](https://github.com/Harvey-N-Lab/hailow/wiki)

## Acknowledgments

Built for engineers who want to level up their AI-assisted development workflow.

---

**Made with ❤️ for the developer community**
