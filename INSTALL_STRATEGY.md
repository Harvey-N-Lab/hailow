# Install Strategy and Compatibility Mapping

## Overview

This document defines how the CLI installs domain configurations into workspaces and maps content to different agent platforms (Roo Code and Claude Code).

## Platform Mapping

### Source Domain Structure

Each domain in the repository follows this structure:

```
<domain>/
  agents/
  rules/
  skills/
  commands/
  contexts/
```

### Target Platform Layouts

#### Roo Code (`--platform roo`)

```
<workspace>/
  .agents/                       # Root agent directory
    <domain>-researcher.md       # Agent files at root
    <domain>-planner.md
    <domain>-architect.md
    <domain>-implementer.md
    <domain>-reviewer.md
    skills/
      <domain>-skills.md
    rules/
      universal-engineering-practices.md   # General rules
      <domain>-*.md                       # Domain rules
    commands/
      <domain>-*.md
    contexts/
      .gitkeep
  AGENT_WORKFLOW.md              # Root guidance file
  .hailow/
    manifest.yaml                # Install manifest
```

#### Claude Code (`--platform claude`)

```
<workspace>/
  .claude/                       # Root agent directory
    <domain>-researcher.md       # Agent files at root
    <domain>-planner.md
    <domain>-architect.md
    <domain>-implementer.md
    <domain>-reviewer.md
    skills/
      <domain>-skills.md
    rules/
      universal-engineering-practices.md   # General rules
      <domain>-*.md                       # Domain rules
    commands/
      <domain>-*.md
    contexts/
      .gitkeep
  CLAUDE.md                      # Root guidance file
  .hailow/
    manifest.yaml                # Install manifest
```

### Mapping Rules

| Source | Roo Code Target | Claude Code Target |
|--------|----------------|-------------------|
| `<domain>/agents/*.md` | `.agents/<filename>` | `.claude/<filename>` |
| `<domain>/rules/*.md` | `.agents/rules/<filename>` | `.claude/rules/<filename>` |
| `<domain>/skills/*.md` | `.agents/skills/<filename>` | `.claude/skills/<filename>` |
| `<domain>/commands/*.md` | `.agents/commands/<filename>` | `.claude/commands/<filename>` |
| `<domain>/contexts/.gitkeep` | `.agents/contexts/.gitkeep` | `.claude/contexts/.gitkeep` |
| `general/rules/*.md` | `.agents/rules/<filename>` | `.claude/rules/<filename>` |
| `AGENT_WORKFLOW.md` | `AGENT_WORKFLOW.md` | `CLAUDE.md` |

## Installation Process

### Phase 1: Validation

1. **Validate source availability**
   - Check if source repository/path exists
   - Verify network connectivity if remote

2. **Validate target workspace**
   - Check if workspace path exists
   - Check for existing installations (read manifest)
   - Warn about potential conflicts

3. **Validate domain selection**
   - Verify requested domains exist in source
   - Check domain metadata and versions

### Phase 2: Conflict Resolution

1. **Check existing files**
   - Read `.hailow/manifest.yaml` if exists
   - Identify files that would be overwritten

2. **Apply conflict resolution strategy**
   - **Default mode**: Skip existing files, warn user
   - **Dry-run mode** (`--dry-run`): Show what would be installed, don't write
   - **Force mode** (`--force`): Overwrite existing files
   - **Merge mode** (`--merge`): Attempt to merge configurations

3. **Directory creation**
   - Create platform-specific root directory (`.agents/` or `.claude/`)
   - Create subdirectories: `skills/`, `rules/`, `commands/`, `contexts/`
   - Create `.hailow/` directory

### Phase 3: File Installation

1. **Copy general rules** (if not already present)
   - Source: `general/rules/*.md`
   - Target: `<platform-root>/rules/`

2. **Copy domain files** (for each selected domain)
   - Agent files: `<domain>/agents/*.md` → `<platform-root>/`
   - Rules: `<domain>/rules/*.md` → `<platform-root>/rules/`
   - Skills: `<domain>/skills/*.md` → `<platform-root>/skills/`
   - Commands: `<domain>/commands/*.md` → `<platform-root>/commands/`
   - Contexts: `<domain>/contexts/.gitkeep` → `<platform-root>/contexts/.gitkeep`

3. **Copy root guidance file**
   - Source: `AGENT_WORKFLOW.md`
   - Target (Roo): `AGENT_WORKFLOW.md`
   - Target (Claude): `CLAUDE.md`

4. **Create/update manifest**
   - Record installation metadata
   - List all installed files
   - Track domain versions
   - Store source information

### Phase 4: Verification

1. **Verify file integrity**
   - Check all files were written successfully
   - Validate file permissions
   - Verify manifest consistency

2. **Report results**
   - List installed files
   - Show any warnings or errors
   - Provide next steps

## Multi-Domain Installation

### Strategy

When multiple domains are installed:

1. **File merging**
   - Agent files from different domains coexist (unique prefixes prevent collision)
   - Rules accumulate (each domain adds its own rules)
   - Skills accumulate
   - Commands accumulate
   - General rules installed once

2. **Manifest tracking**
   - Each domain listed separately in manifest
   - Files tracked per domain for selective removal/upgrade

### Example: Installing DevOps + Python

```bash
hailow install devops-engineer python-backend-engineer --platform roo
```

Result in workspace:

```
.agents/
  devops-researcher.md
  devops-planner.md
  devops-architect.md
  devops-implementer.md
  devops-reviewer.md
  python-researcher.md
  python-planner.md
  python-architect.md
  python-implementer.md
  python-reviewer.md
  rules/
    universal-engineering-practices.md     # General (installed once)
    collaboration-protocol.md              # General (installed once)
    devops-infrastructure-standards.md     # DevOps specific
    devops-security-practices.md           # DevOps specific
    devops-deployment-guidelines.md        # DevOps specific
    python-coding-standards.md             # Python specific
    python-api-design.md                   # Python specific
    python-testing-practices.md            # Python specific
  skills/
    devops-skills.md
    python-skills.md
  commands/
    devops-deploy.md
    devops-rollback.md
    devops-health-check.md
    python-test.md
    python-migrate.md
    python-api-docs.md
  contexts/
    .gitkeep
```

### Install All Domains

```bash
hailow install --all --platform roo
```

This installs all 5 domains into the workspace:
- devops-engineer
- python-backend-engineer
- js-ts-software-engineer
- data-engineer
- machine-learning-engineer

Result: 25 agent files (5 domains × 5 agents), plus all rules, skills, and commands from each domain.

## Source Resolution

### Priority Order

1. **CLI flag**: `--source <url-or-path>`
2. **Config file**: `~/.hailow/config.yaml`
3. **Default**: Public repository URL (hardcoded in CLI)

### Source Types

#### 1. Git Repository (Remote)

**Format**: `https://github.com/username/repo`

**Process**:
1. Clone repository to temporary directory
2. Verify structure (check for `domains/domains.yaml`)
3. Install from temporary clone
4. Clean up temporary directory
5. Record commit hash in manifest

**Example**:
```bash
hailow install devops-engineer --source https://github.com/mycompany/agent-configs
```

#### 2. Git Repository (SSH)

**Format**: `git@github.com:username/repo.git`

**Process**: Same as HTTPS, uses SSH authentication

**Example**:
```bash
hailow install python-backend-engineer --source git@github.com:mycompany/agent-configs.git
```

#### 3. Local Directory

**Format**: `/path/to/directory` or `./relative/path`

**Process**:
1. Verify path exists
2. Verify structure
3. Install directly from local path
4. Record absolute path in manifest

**Example**:
```bash
hailow install data-engineer --source /home/user/my-agent-configs
```

**Use cases**:
- Local development and testing
- Custom configurations
- Offline installations

## Upgrade Strategy

### Upgrade Command

```bash
hailow upgrade [domain-name]
```

**Options**:
- `hailow upgrade` - Upgrade all installed domains
- `hailow upgrade devops-engineer` - Upgrade specific domain
- `hailow upgrade --dry-run` - Show what would be upgraded

### Upgrade Process

1. **Read manifest**
   - Identify installed domains and versions
   - Identify source location

2. **Fetch latest from source**
   - Pull latest from git repository or local path
   - Check for version changes

3. **Compare versions**
   - Determine if upgrade is available
   - Show changelog if available

4. **Apply upgrade**
   - Remove old domain files
   - Install new domain files
   - Update manifest with new versions

5. **Preserve user content**
   - Never overwrite `contexts/` user files
   - Preserve user-modified files (with warning)

## Removal Strategy

### Remove Command

```bash
hailow remove [domain-name]
```

**Examples**:
- `hailow remove devops-engineer` - Remove specific domain
- `hailow remove --all` - Remove all domains

### Removal Process

1. **Read manifest**
   - Identify files belonging to domain

2. **Remove domain files**
   - Delete agent files
   - Delete domain-specific rules
   - Delete domain-specific skills
   - Delete domain-specific commands

3. **Clean up shared resources**
   - If no domains remain, remove general rules
   - If no domains remain, remove root guidance file
   - Keep contexts/ directory (user content)

4. **Update manifest**
   - Remove domain from manifest
   - If manifest empty, delete `.hailow/` directory

## Conflict Handling

### Scenario 1: File Already Exists

**Detection**: File exists at target path

**Resolution**:
- **Default**: Skip file, log warning
- **--force**: Overwrite file
- **--interactive**: Prompt user for each conflict

### Scenario 2: Different Platform Already Installed

**Detection**: Both `.agents/` and `.claude/` exist

**Resolution**:
- Error: "Multiple platforms detected. Please remove existing installation or specify --force"
- User must explicitly choose to proceed

### Scenario 3: Manifest Mismatch

**Detection**: Files exist but not tracked in manifest

**Resolution**:
- Warning: "Untracked files detected. Use --force to overwrite or --merge to attempt merge"
- List untracked files
- Require explicit flag to proceed

### Scenario 4: Version Conflict

**Detection**: Trying to install older version than currently installed

**Resolution**:
- Warning: "Newer version already installed. Use --force to downgrade"
- Show current vs. target version
- Require explicit flag to proceed

## Dry-Run Mode

### Purpose

Allow users to preview installation without making changes.

### Behavior

```bash
hailow install devops-engineer --dry-run
```

**Output**:
```
Dry-run mode: No files will be modified

Installation Plan:
  Platform: roo
  Target: /home/user/myproject
  Source: https://github.com/username/agent-configs (main branch)

Would install:
  [NEW] .agents/devops-researcher.md
  [NEW] .agents/devops-planner.md
  [NEW] .agents/devops-architect.md
  [NEW] .agents/devops-implementer.md
  [NEW] .agents/devops-reviewer.md
  [NEW] .agents/rules/universal-engineering-practices.md
  [NEW] .agents/rules/devops-infrastructure-standards.md
  [NEW] .agents/rules/devops-security-practices.md
  [NEW] .agents/rules/devops-deployment-guidelines.md
  [NEW] .agents/skills/devops-skills.md
  [NEW] .agents/commands/devops-deploy.md
  [NEW] .agents/commands/devops-rollback.md
  [NEW] .agents/commands/devops-health-check.md
  [NEW] AGENT_WORKFLOW.md
  [NEW] .hailow/manifest.yaml

Total: 15 files (12.4 KB)

To proceed, run without --dry-run flag.
```

## Safety Features

1. **Non-destructive by default**: Never overwrite without explicit flag
2. **Dry-run available**: Preview changes before applying
3. **Manifest tracking**: Always know what's installed
4. **Rollback support**: Can remove domains cleanly
5. **Validation**: Verify source and target before proceeding
6. **Clear logging**: Always show what's being done
7. **Preserve user data**: Never touch context files created by user
