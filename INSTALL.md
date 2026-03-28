# Installation Guide

This guide covers installing the `hailow` CLI tool and using it to install domain configurations.

## Installing the CLI

### Method 1: Automated Script (Recommended)

**Linux and macOS:**

```bash
curl -sSL https://raw.githubusercontent.com/username/hailow/main/scripts/install.sh | bash
```

This script will:
- Detect your OS and architecture
- Download the appropriate binary
- Verify the checksum
- Install to `~/.local/bin/` or `/usr/local/bin/`
- Make it executable

**Windows:**

Download the Windows binary from [GitHub Releases](https://github.com/Harvey-N-Lab/hailow/releases/latest) and add to PATH manually.

### Method 2: Manual Download

1. **Download Binary**

   Visit [GitHub Releases](https://github.com/Harvey-N-Lab/hailow/releases/latest) and download the binary for your platform:
   
   - Linux (amd64): `hailow_linux_amd64.tar.gz`
   - Linux (arm64): `hailow_linux_arm64.tar.gz`
   - macOS (Intel): `hailow_darwin_amd64.tar.gz`
   - macOS (Apple Silicon): `hailow_darwin_arm64.tar.gz`
   - Windows: `hailow_windows_amd64.zip`

2. **Extract Archive**

   ```bash
   # Linux/macOS
   tar -xzf hailow_*.tar.gz
   
   # Windows: Extract the .zip file
   ```

3. **Move to PATH**

   ```bash
   # Linux/macOS
   sudo mv hailow /usr/local/bin/
   
   # Or user bin
   mv hailow ~/.local/bin/
   
   # Windows: Move to a directory in PATH
   ```

4. **Verify Installation**

   ```bash
   hailow version
   ```

### Method 3: Build from Source

**Requirements:**
- Go 1.21 or later
- Git

**Steps:**

```bash
# Clone repository
git clone https://github.com/Harvey-N-Lab/hailow
cd hailow

# Build
make build

# Install
make install

# Or manually
go install ./cmd/hailow
```

## Post-Installation Setup

### 1. Verify Installation

```bash
hailow version
```

Expected output:
```
hailow version X.Y.Z
Commit: abc1234
Built: 2024-01-01_12:00:00
```

### 2. Initialize Configuration (Optional)

```bash
hailow config init
```

This creates `~/.hailow/config.yaml` with default settings.

### 3. List Available Domains

```bash
hailow list domains
```

## Installing Domains

### Basic Installation

**Install a single domain:**

```bash
hailow install devops-engineer --platform roo
```

**Install multiple domains:**

```bash
hailow install devops-engineer python-backend-engineer
```

**Install all domains:**

```bash
hailow install --all
```

### Installation Options

```bash
hailow install <domain>... [options]

Options:
  --platform <name>      Target platform: roo or claude (default: roo)
  --workspace <path>     Target directory (default: current directory)
  --source <url-or-path> Custom source (default: public repo)
  --all                  Install all domains
  --dry-run              Preview without making changes
  --force                Overwrite existing files
  --no-general           Skip general rules
```

### Platform-Specific Installation

**For Roo Code:**

```bash
hailow install devops-engineer --platform roo
```

Creates:
```
.agents/
  devops-researcher.md
  devops-planner.md
  devops-architect.md
  devops-implementer.md
  devops-reviewer.md
  rules/
  skills/
  commands/
  contexts/
AGENT_WORKFLOW.md
.hailow/manifest.yaml
```

**For Claude Code:**

```bash
hailow install python-backend-engineer --platform claude
```

Creates:
```
.claude/
  python-researcher.md
  python-planner.md
  python-architect.md
  python-implementer.md
  python-reviewer.md
  rules/
  skills/
  commands/
  contexts/
CLAUDE.md
.hailow/manifest.yaml
```

## Common Installation Scenarios

### Scenario 1: DevOps Team

```bash
cd /path/to/infrastructure-repo
hailow install devops-engineer --platform roo
```

### Scenario 2: Full-Stack Project

```bash
cd /path/to/fullstack-app
hailow install js-ts-software-engineer python-backend-engineer devops-engineer
```

### Scenario 3: Data Science Team

```bash
cd /path/to/ml-project
hailow install data-engineer machine-learning-engineer
```

### Scenario 4: Custom Source

```bash
# Use company's custom configurations
hailow install devops-engineer --source https://github.com/mycompany/configs
```

### Scenario 5: Local Development

```bash
# Use local configurations for testing
hailow install python-backend-engineer --source ./my-local-configs
```

## Customizing Installation

### Set Default Platform

```bash
hailow config set platform claude
```

Now installations default to Claude Code.

### Set Custom Source

```bash
hailow config set source.url https://github.com/mycompany/hailow
```

### Change Conflict Behavior

```bash
hailow config set install.on_conflict overwrite
```

Options: `skip`, `overwrite`, `prompt`

## Verifying Installation

### Check Installed Domains

```bash
hailow list installed
```

### Verify Files

**Roo Code:**
```bash
ls -la .agents/
cat .agents/devops-researcher.md
```

**Claude Code:**
```bash
ls -la .claude/
cat .claude/python-implementer.md
```

### Check Manifest

```bash
cat .hailow/manifest.yaml
```

## Troubleshooting

### Issue: command not found

**Solution:**
- Ensure binary is in PATH
- Check with: `echo $PATH`
- Add directory to PATH: `export PATH="$HOME/.local/bin:$PATH"`
- Add to shell rc file for persistence

### Issue: Permission denied

**Solution:**
```bash
chmod +x /path/to/hailow
```

### Issue: Files already exist

**Options:**

1. **Skip existing files** (default)
2. **Force overwrite:**
   ```bash
   hailow install devops-engineer --force
   ```
3. **Remove and reinstall:**
   ```bash
   hailow remove devops-engineer
   hailow install devops-engineer
   ```

### Issue: Cannot access source repository

**Check:**
- Network connectivity
- Repository URL correctness
- Authentication if private repo

**Test:**
```bash
git clone https://github.com/Harvey-N-Lab/hailow
```

### Issue: Wrong platform installed

**Solution:**
```bash
hailow remove --all
hailow install <domain> --platform <correct-platform>
```

## Uninstallation

### Remove Domains

```bash
hailow remove devops-engineer
hailow remove --all
```

### Remove CLI

```bash
# Linux/macOS
rm /usr/local/bin/hailow
rm -rf ~/.hailow/

# Windows
# Delete binary and .hailow directory manually
```

## Next Steps

After installation:

1. **Review the workflow:** Read [`AGENT_WORKFLOW.md`](AGENT_WORKFLOW.md) in your workspace
2. **Review domain rules:** Check `.agents/rules/` or `.claude/rules/`
3. **Install recommended skills:** See domain skills files
4. **Add project context:** Place project-specific docs in `contexts/`
5. **Start using agents:** Begin with the researcher agent for your first task

## Getting Help

- Run `hailow --help` for command help
- Run `hailow <command> --help` for command-specific help
- Check [GitHub Issues](https://github.com/Harvey-N-Lab/hailow/issues)
- Read [CLI_DESIGN.md](CLI_DESIGN.md) for detailed CLI documentation

## Updating the CLI

See [UPGRADE.md](UPGRADE.md) for upgrade instructions.
