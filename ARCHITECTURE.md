# Architecture Decisions

## Executive Summary

This repository provides a production-ready system for managing and distributing reusable Claude Code and Roo Code configurations across multiple engineering domains. The solution includes:

1. **Domain Configuration Library**: Pre-built agent configurations, rules, skills, commands, and contexts for 5 engineering domains
2. **CLI Tool (`hailow`)**: Go-based binary for installing configurations into any workspace
3. **Multi-Platform Support**: Compatible with both Roo Code and Claude Code
4. **Flexible Sourcing**: Pull from default public repo, alternate repos, or local paths
5. **Easy Distribution**: GitHub Releases with automated builds for Linux, macOS, and Windows

## Key Architecture Decisions

### 1. CLI Language: Go

**Decision**: Use Go for the CLI binary

**Rationale**:
- Single binary distribution (no runtime dependencies)
- Excellent cross-platform support (Linux, macOS, Windows)
- Strong standard library for file operations, HTTP, and CLI building
- Fast compilation and execution
- Mature tooling ecosystem (goreleaser for releases)

**Rejected Alternative**: Python/Node.js (requires runtime, harder to distribute)

### 2. Configuration Format: YAML

**Decision**: Use YAML for CLI configuration file

**Rationale**:
- Human-readable and easy to edit
- Widely adopted in DevOps and engineering tools
- Good balance between simplicity and expressiveness
- Native Go library support

**Rejected Alternative**: TOML (less familiar to most developers), JSON (less human-friendly)

### 3. File Naming Strategy: Domain Prefixes

**Decision**: Prefix all markdown files with domain identifier

**Rationale**:
- Prevents filename collisions when multiple domains installed together
- Makes file origin immediately clear
- Enables simple merge strategy without complex conflict resolution
- Examples: `devops-researcher.md`, `python-planner.md`, `mleng-reviewer.md`

**Rejected Alternative**: Subdirectories per domain (harder to navigate, breaks agent discovery patterns)

### 4. Install Strategy: Direct File Copy with Manifest

**Decision**: Copy domain files directly to workspace with platform-specific layout, maintain install manifest

**Rationale**:
- Simple and predictable behavior
- Easy to understand and debug
- Manifest enables upgrade/removal operations
- No complex linking or symlinks needed

**Rejected Alternative**: Symlinks (platform compatibility issues), Git submodules (too complex for users)

### 5. Platform Compatibility: Mapping Layer

**Decision**: Implement internal mapping from domain structure to platform-specific layouts

**Rationale**:
- Single source of truth for domain content
- Clean separation of domain definitions from platform requirements
- Easy to extend for new platforms

**Platform Mappings**:
- **Roo Code**: `.agents/`, `.agents/skills/`, `.agents/rules/`, `.agents/commands/`, `.agents/contexts/`
- **Claude Code**: `.claude/`, `.claude/skills/`, `.claude/rules/`, `.claude/commands/`, `.claude/contexts/`

**Rejected Alternative**: Duplicate content per platform (maintenance burden)

### 6. Release Strategy: GitHub Actions + goreleaser

**Decision**: Use GitHub Actions with goreleaser for automated releases

**Rationale**:
- Automated multi-platform builds (Linux, macOS, Windows)
- Semantic versioning support
- Checksums and release notes generation
- GitHub Releases integration
- Standard industry practice

**Rejected Alternative**: Manual builds (error-prone, time-consuming)

### 7. Installation Method: Shell Script + Direct Download

**Decision**: Provide install.sh script for Unix-like systems, direct binary download for Windows

**Rationale**:
- Simple one-liner installation for most users
- Falls back to manual download if needed
- No dependency on package managers
- Easy to understand and modify

**Rejected Alternative**: Package managers only (limited reach), Docker (overkill for CLI tool)

### 8. Domain Content Organization

**Decision**: Each domain contains: agents, rules, skills, commands, contexts

**Rationale**:
- Consistent structure across all domains
- Clear separation of concerns
- Extensible for future additions
- Matches user mental model

**Structure**:
```
<domain>/
  agents/      # AI agent prompts (researcher, planner, architect, implementer, reviewer)
  rules/       # Coding standards, workflow rules, safety guidelines
  skills/      # Skill definitions and install instructions
  commands/    # Slash commands and workflows
  contexts/    # Placeholder for user-specific context
```

### 9. Configuration Sources

**Decision**: Support multiple source types with clear precedence

**Rationale**:
- Default: Public repository (zero configuration)
- Override: Custom repository (enterprise/team customization)
- Override: Local path (development and testing)

**Precedence**: CLI flag > config file > default

**Rejected Alternative**: Single source type (inflexible)

### 10. General vs Domain-Specific Rules

**Decision**: `general/` folder contains cross-domain rules, domains have specific rules

**Rationale**:
- Avoid duplication of universal rules
- Easy to update global conventions
- Domains can override or extend general rules

## Repository Structure Philosophy

The repository structure follows these principles:

1. **Clarity**: Structure immediately communicates purpose
2. **Modularity**: Each domain is self-contained
3. **Extensibility**: Easy to add new domains
4. **Maintainability**: Changes isolated to relevant sections
5. **User-Friendly**: New users can understand layout quickly

## CLI Design Philosophy

The CLI follows these principles:

1. **Simple by Default**: Common operations require minimal flags
2. **Explicit When Needed**: Advanced options available but not required
3. **Safe First**: Dry-run mode, no destructive defaults
4. **Clear Feedback**: Informative logs and error messages
5. **Composable**: Commands can be scripted and automated

## Installation Experience Goals

1. **One-liner install**: `curl -sSL https://raw.githubusercontent.com/USER/REPO/main/install.sh | bash`
2. **Quick start**: Install a domain with one command
3. **Clear documentation**: Examples for all common scenarios
4. **Zero surprises**: Explicit about what will be installed where
5. **Easy upgrade**: Simple command to update to latest version
