# Upgrade Guide

This guide covers upgrading the `hailow` CLI tool and installed domain configurations.

## Upgrading the CLI

### Method 1: Using Install Script

The same install script works for upgrades:

```bash
curl -sSL https://raw.githubusercontent.com/username/hailow/main/scripts/install.sh | bash
```

This will:
- Download the latest version
- Replace the existing binary
- Preserve your configuration

### Method 2: Manual Download

1. Download the latest release from [GitHub Releases](https://github.com/Harvey-N-Lab/hailow/releases/latest)
2. Replace the existing binary
3. Verify the new version

```bash
hailow version
```

### Method 3: Build from Source

```bash
cd hailow
git pull origin main
make build
make install
```

## Upgrading Installed Domains

### Check for Updates

```bash
# Check installed domains
hailow list installed

# Preview what would be upgraded
hailow upgrade --dry-run
```

### Upgrade All Domains

```bash
hailow upgrade
```

This will:
- Check for newer versions of all installed domains
- Download updates from the source
- Replace old files with new ones
- Update the manifest

### Upgrade Specific Domains

```bash
hailow upgrade devops-engineer python-backend-engineer
```

### Upgrade Options

```bash
hailow upgrade [domain...] [options]

Options:
  --workspace <path>    Target workspace (default: current directory)
  --source <url>        Source to upgrade from (default: from manifest)
  --dry-run             Show what would be upgraded
  --force               Force upgrade even if versions match
```

## Version Compatibility

### Semantic Versioning

The project follows semantic versioning (MAJOR.MINOR.PATCH):

- **MAJOR**: Breaking changes (may require migration)
- **MINOR**: New features (backward compatible)
- **PATCH**: Bug fixes (backward compatible)

### Compatibility Matrix

| CLI Version | Domain Version | Compatible |
|-------------|----------------|------------|
| 1.x.x | 1.x.x | ✅ |
| 1.x.x | 0.x.x | ⚠️ Upgrade recommended |
| 2.x.x | 1.x.x | ❌ Requires migration |

## Migration Guides

### Upgrading from v1 to v2 (Future)

When a major version change occurs, migration steps will be provided here.

## Backup and Safety

### Before Upgrading

Always backup important files:

```bash
# Backup workspace configurations
tar -czf backup-$(date +%Y%m%d).tar.gz .agents/ .hailow/ AGENT_WORKFLOW.md
```

### Safe Upgrade Process

1. **Check current state**
   ```bash
   hailow list installed
   ```

2. **Backup existing configuration**
   ```bash
   cp -r .agents .agents.backup
   ```

3. **Dry-run upgrade**
   ```bash
   hailow upgrade --dry-run
   ```

4. **Proceed with upgrade**
   ```bash
   hailow upgrade
   ```

5. **Verify upgrade**
   ```bash
   hailow list installed
   ls -la .agents/
   ```

6. **Test with AI assistant**
   - Verify agents load correctly
   - Test a simple workflow

7. **Remove backup if successful**
   ```bash
   rm -rf .agents.backup
   ```

## Handling Upgrade Conflicts

### File Conflicts

If files have been modified locally:

**Option 1: Keep local changes**
```bash
# Upgrade will skip modified files by default
hailow upgrade
```

**Option 2: Overwrite with new versions**
```bash
hailow upgrade --force
```

**Option 3: Manual merge**
```bash
# Backup your changes
cp .agents/devops-researcher.md .agents/devops-researcher.md.local

# Upgrade
hailow upgrade --force

# Manually merge changes
diff .agents/devops-researcher.md.local .agents/devops-researcher.md
```

### Platform Migration

**From Roo Code to Claude Code:**

```bash
# Remove Roo installation
hailow remove --all

# Install for Claude
hailow config set platform claude
hailow install --all
```

**From Claude Code to Roo Code:**

```bash
# Remove Claude installation
hailow remove --all

# Install for Roo
hailow config set platform roo
hailow install --all
```

## Upgrading Source Location

### Switch to Custom Repository

```bash
# Update configuration
hailow config set source.url https://github.com/mycompany/configs

# Upgrade from new source
hailow upgrade
```

### Switch Back to Public Repository

```bash
hailow config set source.url https://github.com/Harvey-N-Lab/hailow
hailow upgrade
```

## Rollback

### Rolling Back Domain Updates

1. **Check manifest for previous version** (future feature)
2. **Restore from backup:**
   ```bash
   rm -rf .agents
   cp -r .agents.backup .agents
   ```

### Rolling Back CLI Version

1. **Download previous version** from GitHub Releases
2. **Replace current binary**
3. **Verify version:**
   ```bash
   hailow version
   ```

## Post-Upgrade Verification

### Checklist

- [ ] CLI version updated: `hailow version`
- [ ] Domains upgraded: `hailow list installed`
- [ ] Files present: `ls .agents/` or `ls .claude/`
- [ ] Manifest updated: `cat .hailow/manifest.yaml`
- [ ] Agent files load in IDE
- [ ] Test workflow with AI assistant
- [ ] No breaking changes in your workflow

### Testing After Upgrade

1. **Open a file** referenced by agents
2. **Ask the AI assistant** a simple question
3. **Verify agent responds** appropriately
4. **Test a workflow** (e.g., researcher → planner → implementer)

## Upgrade Frequency Recommendations

### CLI Tool

- **Check monthly** for updates
- **Upgrade when:**
  - New features you need
  - Bug fixes
  - Security patches

### Domain Configurations

- **Check quarterly** for updates
- **Upgrade when:**
  - New best practices added
  - Improved prompts
  - Bug fixes in workflows
  - New skills or commands

## Staying Informed

### Release Notifications

**Watch the repository** on GitHub:
- Go to: https://github.com/Harvey-N-Lab/hailow
- Click "Watch" → "Custom" → "Releases"

**Subscribe to releases:**
- GitHub will email you for new releases

### Changelog

Check the [CHANGELOG.md](CHANGELOG.md) or GitHub Releases for:
- New features
- Bug fixes
- Breaking changes
- Migration notes

## Troubleshooting Upgrades

### Issue: Upgrade fails midway

**Solution:**
1. Check network connectivity
2. Verify source accessibility
3. Check disk space
4. Review error messages
5. Restore from backup if needed

### Issue: New version breaks workflows

**Solution:**
1. Check CHANGELOG for breaking changes
2. Review migration guide
3. Rollback if necessary
4. Report issue on GitHub

### Issue: Files not updating

**Solution:**
```bash
# Force upgrade
hailow upgrade --force

# Or reinstall
hailow remove devops-engineer
hailow install devops-engineer
```

### Issue: Manifest corruption

**Solution:**
```bash
# Backup current state
cp .hailow/manifest.yaml .hailow/manifest.yaml.backup

# Remove and reinstall
hailow remove --all
hailow install <your-domains>
```

## Best Practices

1. **Backup before upgrading** major versions
2. **Test in non-production** workspace first
3. **Read release notes** before upgrading
4. **Upgrade CLI** before domain configs
5. **Keep track** of what domains you have installed
6. **Document custom changes** to domain files
7. **Use version control** for your workspace

## Emergency Procedures

### Complete Reset

If everything breaks:

```bash
# Remove all installed domains
hailow remove --all

# Remove CLI configuration
rm -rf ~/.hailow/

# Reinstall CLI
curl -sSL https://raw.githubusercontent.com/username/hailow/main/scripts/install.sh | bash

# Reinstall domains
hailow install <your-domains>
```

## Getting Help

- Check [GitHub Issues](https://github.com/Harvey-N-Lab/hailow/issues) for known issues
- Search [Discussions](https://github.com/Harvey-N-Lab/hailow/discussions) for solutions
- Create an issue if you encounter problems
- Include CLI version and error messages

## Next Steps

After upgrading:

1. Review [CHANGELOG.md](CHANGELOG.md) for what's new
2. Check if new domains were added
3. Review updated agent prompts
4. Test your common workflows
5. Update team documentation if needed

---

**Last Updated:** 2024-01-01  
**CLI Version:** 1.0.0
