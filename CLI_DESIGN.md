# CLI Design

## CLI Name

**`hailow`** - Agent Configuration Manager

## Commands Overview

```
hailow [command] [options]

Commands:
  install       Install domain configurations into workspace
  list          List available domains or platforms
  remove        Remove installed domain configurations
  upgrade       Upgrade installed domains to latest versions
  config        Manage CLI configuration
  version       Show version information
  help          Show help information

Global Flags:
  --help, -h        Show help
  --version, -v     Show version
  --config FILE     Use specific config file (default: ~/.hailow/config.yaml)
  --verbose         Enable verbose logging
  --quiet           Suppress non-error output
```

## Command: install

Install one or more domain configurations into a workspace.

### Syntax

```bash
hailow install <domain>... [options]
```

### Arguments

- `<domain>...` - One or more domain names to install. Use `--all` to install all domains.

### Options

| Flag | Description | Default |
|------|-------------|---------|
| `--platform <name>` | Target platform: `roo` or `claude` | Value from config, or `roo` |
| `--workspace <path>` | Target workspace directory | Current directory (`.`) |
| `--source <url-or-path>` | Source repository or local path | Value from config, or default repo |
| `--all` | Install all available domains | false |
| `--dry-run` | Show what would be installed without making changes | false |
| `--force` | Overwrite existing files | false |
| `--interactive` | Prompt for conflicts | false |
| `--no-general` | Skip installing general rules | false |

### Examples

**Install single domain for Roo Code:**
```bash
hailow install devops-engineer --platform roo
```

**Install multiple domains:**
```bash
hailow install devops-engineer python-backend-engineer --platform claude
```

**Install all domains:**
```bash
hailow install --all --platform roo
```

**Install from custom source:**
```bash
hailow install data-engineer --source https://github.com/mycompany/configs
```

**Install from local directory:**
```bash
hailow install mleng --source ./my-configs
```

**Install into specific workspace:**
```bash
hailow install js-ts-software-engineer --workspace ~/projects/myapp
```

**Dry-run to preview:**
```bash
hailow install devops-engineer --dry-run
```

**Force overwrite existing files:**
```bash
hailow install python-backend-engineer --force
```

### Domain Name Aliases

For convenience, the CLI supports short aliases:

| Full Name | Aliases |
|-----------|---------|
| `devops-engineer` | `devops`, `ops` |
| `python-backend-engineer` | `python`, `py`, `python-backend` |
| `js-ts-software-engineer` | `jsts`, `js`, `ts`, `javascript`, `typescript` |
| `data-engineer` | `data`, `de` |
| `machine-learning-engineer` | `mleng`, `ml`, `machine-learning` |

## Command: list

List available domains or supported platforms.

### Syntax

```bash
hailow list [type] [options]
```

### Arguments

- `type` - What to list: `domains`, `platforms`, or `installed` (default: `domains`)

### Options

| Flag | Description | Default |
|------|-------------|---------|
| `--source <url-or-path>` | Source to list domains from | Value from config, or default repo |
| `--workspace <path>` | Workspace to check installations | Current directory (`.`) |
| `--format <format>` | Output format: `table`, `json`, `yaml` | `table` |

### Examples

**List available domains:**
```bash
hailow list domains
```

**List supported platforms:**
```bash
hailow list platforms
```

**List installed domains in workspace:**
```bash
hailow list installed
```

**List domains from custom source:**
```bash
hailow list domains --source https://github.com/mycompany/configs
```

**Output as JSON:**
```bash
hailow list domains --format json
```

### Sample Output

**`hailow list domains`**
```
Available Domains:

Name                           Description                                    Version
----                           -----------                                    -------
devops-engineer                Infrastructure, CI/CD, and deployment configs  1.0.0
python-backend-engineer        Python backend development configurations      1.0.0
js-ts-software-engineer        JavaScript/TypeScript development configs      1.0.0
data-engineer                  Data pipeline and engineering configurations   1.0.0
machine-learning-engineer      ML development and deployment configurations   1.0.0

Use 'hailow install <domain>' to install a domain.
```

**`hailow list installed`**
```
Installed Domains in /home/user/myproject:

Platform: roo

Domain                         Version    Installed
------                         -------    ---------
devops-engineer                1.0.0      2026-03-28 10:30:00
python-backend-engineer        1.0.0      2026-03-28 10:30:05

Total: 2 domains (45 files)

Use 'hailow upgrade' to update to latest versions.
```

## Command: remove

Remove installed domain configurations from workspace.

### Syntax

```bash
hailow remove <domain>... [options]
```

### Arguments

- `<domain>...` - One or more domain names to remove. Use `--all` to remove all.

### Options

| Flag | Description | Default |
|------|-------------|---------|
| `--workspace <path>` | Target workspace directory | Current directory (`.`) |
| `--all` | Remove all installed domains | false |
| `--keep-contexts` | Preserve contexts directory | true |
| `--dry-run` | Show what would be removed | false |

### Examples

**Remove single domain:**
```bash
hailow remove devops-engineer
```

**Remove multiple domains:**
```bash
hailow remove devops-engineer python-backend-engineer
```

**Remove all domains:**
```bash
hailow remove --all
```

**Preview removal:**
```bash
hailow remove data-engineer --dry-run
```

## Command: upgrade

Upgrade installed domains to latest versions.

### Syntax

```bash
hailow upgrade [domain]... [options]
```

### Arguments

- `[domain]...` - One or more domain names to upgrade. If omitted, upgrades all.

### Options

| Flag | Description | Default |
|------|-------------|---------|
| `--workspace <path>` | Target workspace directory | Current directory (`.`) |
| `--source <url-or-path>` | Source repository or local path | Value from manifest |
| `--dry-run` | Show what would be upgraded | false |
| `--force` | Force upgrade even if versions match | false |

### Examples

**Upgrade all domains:**
```bash
hailow upgrade
```

**Upgrade specific domain:**
```bash
hailow upgrade devops-engineer
```

**Upgrade multiple domains:**
```bash
hailow upgrade python-backend-engineer data-engineer
```

**Preview upgrade:**
```bash
hailow upgrade --dry-run
```

## Command: config

Manage CLI configuration.

### Syntax

```bash
hailow config <subcommand> [options]
```

### Subcommands

#### config show

Show current configuration.

```bash
hailow config show
```

**Output:**
```yaml
source:
  type: git
  url: https://github.com/username/agent-configs
  branch: main
platform: roo
paths:
  config_file: /home/user/.hailow/config.yaml
  cache_dir: /home/user/.hailow/cache
```

#### config set

Set configuration value.

```bash
hailow config set <key> <value>
```

**Examples:**
```bash
hailow config set platform claude
hailow config set source.url https://github.com/mycompany/configs
hailow config set source.type local
hailow config set source.path /home/user/my-configs
```

#### config get

Get configuration value.

```bash
hailow config get <key>
```

**Examples:**
```bash
hailow config get platform
hailow config get source.url
```

#### config init

Initialize configuration file with defaults.

```bash
hailow config init [options]
```

**Options:**
- `--force` - Overwrite existing config file

**Example:**
```bash
hailow config init
```

#### config reset

Reset configuration to defaults.

```bash
hailow config reset [options]
```

**Options:**
- `--confirm` - Skip confirmation prompt

**Example:**
```bash
hailow config reset --confirm
```

## Command: version

Show version information.

### Syntax

```bash
hailow version [options]
```

### Options

| Flag | Description |
|------|-------------|
| `--short` | Show only version number |
| `--check` | Check for updates |

### Examples

**Show version:**
```bash
hailow version
```

**Output:**
```
hailow version 1.0.0
Build: abc123
Built: 2026-03-28T10:00:00Z
Go: go1.21.0
OS/Arch: linux/amd64
```

**Short version:**
```bash
hailow version --short
```

**Output:**
```
1.0.0
```

## Configuration File

### Location

Default: `~/.hailow/config.yaml`

Can be overridden with `--config` flag or `AGENTCONFIG_CONFIG` environment variable.

### Format (YAML)

```yaml
# Agent Configuration Manager Config
# Version: 1.0

# Default source for domain configurations
source:
  # Source type: "git" or "local"
  type: git
  
  # For git source:
  url: https://github.com/username/agent-configs
  branch: main  # optional, defaults to "main"
  
  # For local source:
  # type: local
  # path: /path/to/configs

# Default target platform: "roo" or "claude"
platform: roo

# Installation behavior
install:
  # Default behavior for existing files: "skip", "overwrite", "prompt"
  on_conflict: skip
  
  # Always install general rules with domains
  include_general: true
  
  # Create backup before overwriting
  backup: true

# Paths
paths:
  # Cache directory for temporary downloads
  cache_dir: ~/.hailow/cache
  
  # Cache retention in hours (0 = no retention, always re-download)
  cache_retention: 24

# Logging
logging:
  # Log level: "debug", "info", "warn", "error"
  level: info
  
  # Log file (optional)
  file: ~/.hailow/hailow.log

# Update checking
updates:
  # Check for CLI updates
  check_on_run: true
  
  # Update check interval in hours
  interval: 24
```

### Minimal Configuration

```yaml
source:
  type: git
  url: https://github.com/username/agent-configs
platform: roo
```

### Environment Variables

Configuration can also be set via environment variables:

| Variable | Description | Example |
|----------|-------------|---------|
| `AGENTCONFIG_CONFIG` | Config file path | `~/.hailow/config.yaml` |
| `AGENTCONFIG_PLATFORM` | Default platform | `roo` |
| `AGENTCONFIG_SOURCE_TYPE` | Source type | `git` |
| `AGENTCONFIG_SOURCE_URL` | Source URL | `https://github.com/...` |
| `AGENTCONFIG_SOURCE_PATH` | Source path (local) | `/path/to/configs` |
| `AGENTCONFIG_CACHE_DIR` | Cache directory | `~/.hailow/cache` |
| `AGENTCONFIG_LOG_LEVEL` | Log level | `debug` |

### Configuration Precedence

1. CLI flags (highest priority)
2. Environment variables
3. Configuration file
4. Built-in defaults (lowest priority)

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Invalid arguments |
| 3 | Source not found |
| 4 | Domain not found |
| 5 | Workspace error |
| 6 | Conflict detected |
| 7 | Network error |
| 8 | Permission denied |

## Error Handling

### Example Error Messages

**Domain not found:**
```
Error: Domain 'invalid-domain' not found in source

Available domains:
  - devops-engineer
  - python-backend-engineer
  - js-ts-software-engineer
  - data-engineer
  - machine-learning-engineer

Use 'hailow list domains' to see all available domains.
```

**Conflict detected:**
```
Error: Files already exist in workspace

Conflicting files:
  - .agents/devops-researcher.md
  - .agents/devops-planner.md
  (3 more...)

Options:
  - Use --force to overwrite existing files
  - Use --interactive to resolve conflicts interactively
  - Remove existing installation with 'hailow remove'
```

**Source not accessible:**
```
Error: Cannot access source repository

URL: https://github.com/username/invalid-repo
Reason: Repository not found (404)

Please check:
  - Repository URL is correct
  - Repository is public or you have access
  - Network connection is working

Update source with 'hailow config set source.url <url>'
```

## Shell Completion

The CLI supports shell completion for bash, zsh, and fish.

### Installation

**Bash:**
```bash
hailow completion bash > /etc/bash_completion.d/hailow
```

**Zsh:**
```bash
hailow completion zsh > "${fpath[1]}/_hailow"
```

**Fish:**
```bash
hailow completion fish > ~/.config/fish/completions/hailow.fish
```

### Features

- Command completion
- Domain name completion
- Platform name completion
- Flag completion
- File path completion

## Workflow Examples

### First-Time Setup

```bash
# Install hailow
curl -sSL https://raw.githubusercontent.com/username/agent-configs/main/install.sh | bash

# Initialize config (optional, uses defaults if skipped)
hailow config init

# List available domains
hailow list domains

# Install a domain
hailow install devops-engineer --platform roo
```

### Using Custom Source

```bash
# Configure custom source repository
hailow config set source.url https://github.com/mycompany/agent-configs

# Install from custom source
hailow install python-backend-engineer
```

### Using Local Development Source

```bash
# Set local path as source
hailow config set source.type local
hailow config set source.path /home/user/my-agent-configs

# Install from local source
hailow install data-engineer
```

### Multi-Domain Workspace

```bash
# Install multiple domains at once
hailow install devops-engineer python-backend-engineer data-engineer

# Or install all domains
hailow install --all
```

### Upgrade Workflow

```bash
# Check what's installed
hailow list installed

# Preview upgrades
hailow upgrade --dry-run

# Apply upgrades
hailow upgrade
```

### Switching Platforms

```bash
# Remove Roo Code installation
hailow remove --all

# Install for Claude Code instead
hailow config set platform claude
hailow install devops-engineer
```
