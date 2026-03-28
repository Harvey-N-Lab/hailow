# Repository Structure

## Complete Repository Tree

```
.
├── README.md                          # Main documentation
├── ARCHITECTURE.md                    # Architecture decisions
├── INSTALL.md                         # Installation guide
├── UPGRADE.md                         # Upgrade guide
├── LICENSE                            # License file
├── .gitignore                         # Git ignore patterns
├── AGENT_WORKFLOW.md                  # Root guidance for agent collaboration
│
├── src/                               # Source domain configurations
│   ├── general/                       # Cross-domain rules and guidance
│   │   └── rules/
│   │       ├── universal-engineering-practices.md
│   │       └── collaboration-protocol.md
│   │
│   ├── devops-engineer/               # DevOps domain
│   │   ├── agents/
│   │   │   ├── devops-researcher.md
│   │   │   ├── devops-planner.md
│   │   │   ├── devops-architect.md
│   │   │   ├── devops-implementer.md
│   │   │   └── devops-reviewer.md
│   │   ├── rules/
│   │   │   ├── devops-infrastructure-standards.md
│   │   │   ├── devops-security-practices.md
│   │   │   └── devops-deployment-guidelines.md
│   │   ├── skills/
│   │   │   └── devops-skills.md
│   │   ├── commands/
│   │   │   ├── devops-deploy.md
│   │   │   ├── devops-rollback.md
│   │   │   └── devops-health-check.md
│   │   └── contexts/
│   │       └── .gitkeep
│   │
│   ├── python-backend-engineer/       # Python Backend domain
│   │   ├── agents/
│   │   ├── rules/
│   │   ├── skills/
│   │   ├── commands/
│   │   └── contexts/
│   │
│   ├── js-ts-software-engineer/       # JavaScript/TypeScript domain
│   │   ├── agents/
│   │   ├── rules/
│   │   ├── skills/
│   │   ├── commands/
│   │   └── contexts/
│   │
│   ├── data-engineer/                 # Data Engineering domain
│   │   ├── agents/
│   │   ├── rules/
│   │   ├── skills/
│   │   ├── commands/
│   │   └── contexts/
│   │
│   └── machine-learning-engineer/     # ML Engineering domain
│       ├── agents/
│       ├── rules/
│       ├── skills/
│       ├── commands/
│       └── contexts/
│
├── domains/                           # Domain metadata and manifests
│   ├── domains.yaml                   # Domain registry
│   └── README.md                      # Domain documentation
│
├── cmd/                               # CLI application entry point
│   └── hailow/
│       └── main.go                    # Main CLI entry
│
├── internal/                          # Internal Go packages
│   ├── config/
│   │   ├── config.go                  # Config management
│   │   └── sources.go                 # Source management
│   ├── domain/
│   │   ├── domain.go                  # Domain definitions
│   │   ├── loader.go                  # Domain loading
│   │   └── validator.go               # Domain validation
│   ├── installer/
│   │   ├── installer.go               # Install logic
│   │   ├── manifest.go                # Install manifest
│   │   └── platform.go                # Platform mapping
│   ├── fetcher/
│   │   ├── git.go                     # Git repository fetcher
│   │   ├── local.go                   # Local path fetcher
│   │   └── fetcher.go                 # Fetcher interface
│   └── cli/
│       ├── install.go                 # Install command
│       ├── list.go                    # List command
│       ├── config.go                  # Config commands
│       └── version.go                 # Version command
│
├── docs/                              # Additional documentation
│   ├── CLI_USAGE.md                   # CLI usage examples
│   ├── DOMAIN_GUIDE.md                # Guide to creating domains
│   └── CONTRIBUTING.md                # Contribution guidelines
│
├── scripts/                           # Utility scripts
│   ├── install.sh                     # Unix installation script
│   └── validate-domains.sh            # Domain validation script
│
├── .github/
│   └── workflows/
│       ├── release.yml                # Release automation
│       ├── test.yml                   # CI tests
│       └── validate.yml               # Domain validation
│
├── go.mod                             # Go module definition
├── go.sum                             # Go dependencies
├── .goreleaser.yml                    # GoReleaser configuration
└── Makefile                           # Build automation

```

## Key Changes

### src/ Directory

All domain configurations are now organized under `src/` for better management:

- `src/general/` - Cross-domain rules
- `src/<domain>/` - Individual domain configurations

This separates domain content from project infrastructure (CLI, docs, workflows).

### Skills Approach

Skills files now use `npx skills` workflow:
- Each `skills/<domain>-skills.md` lists recommended skills to install
- Users run `npx skills` to browse and install recommended skills
- Skills are managed through the Roo Code skills system, not copied files

## Domain Loading

The CLI loads domains from the `src/` directory:

```go
// internal/domain/domain.go
func GetDomainPath(domain string) string {
    return "src/" + domain
}
```

When installing, the CLI reads from `src/<domain>/` and copies to the workspace.

## Install Process

1. **User runs**: `hailow install devops-engineer`
2. **CLI reads from**: `src/devops-engineer/`
3. **CLI installs to**: Workspace `.agents/` or `.claude/`
4. **Skills note**: Workspace gets `skills/devops-skills.md` with instructions to run `npx skills`

## Benefits

- **Cleaner root**: Domain content separated from infrastructure
- **Easier management**: All domains in one place (`src/`)
- **Better skills handling**: Uses Roo Code's native skill system
- **Scalable**: Easy to add new domains under `src/`

## Workspace Install Layout

(Unchanged from before - still installs to `.agents/` or `.claude/`)
