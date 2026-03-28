# GitHub Release Strategy

## Overview

The CLI tool (`hailow`) is distributed through GitHub Releases with automated builds using GoReleaser and GitHub Actions.

## Versioning

### Semantic Versioning

All releases follow semantic versioning: `MAJOR.MINOR.PATCH`

- **MAJOR**: Breaking changes (incompatible API changes)
- **MINOR**: New features (backward compatible)
- **PATCH**: Bug fixes (backward compatible)

**Examples:**
- `1.0.0` - Initial stable release
- `1.1.0` - Add new command or feature
- `1.1.1` - Bug fix
- `2.0.0` - Breaking change (e.g., config format change)

### Version Tags

Git tags trigger releases:
- `v1.0.0` - Full release
- `v1.1.0-beta.1` - Pre-release (beta)
- `v1.1.0-rc.1` - Release candidate

## Release Process

### Automated Release (Recommended)

1. **Create and push a version tag:**
   ```bash
   git tag -a v1.0.0 -m "Release v1.0.0"
   git push origin v1.0.0
   ```

2. **GitHub Actions automatically:**
   - Runs tests
   - Builds binaries for multiple platforms
   - Generates checksums
   - Creates GitHub Release
   - Uploads binaries and checksums
   - Generates release notes

3. **Release is published:**
   - Binaries available for download
   - Release notes auto-generated from commits
   - Checksums for verification

### Manual Release (Fallback)

If automated release fails:

1. **Build locally:**
   ```bash
   goreleaser release --clean
   ```

2. **Manually create GitHub Release:**
   - Create release from tag
   - Upload binaries from `dist/`
   - Add release notes

## Build Targets

### Supported Platforms

The CLI is built for these platforms:

| OS | Architecture | Binary Name |
|----|--------------|-------------|
| Linux | amd64 (x86_64) | `hailow_linux_amd64` |
| Linux | arm64 | `hailow_linux_arm64` |
| macOS | amd64 (Intel) | `hailow_darwin_amd64` |
| macOS | arm64 (Apple Silicon) | `hailow_darwin_arm64` |
| Windows | amd64 | `hailow_windows_amd64.exe` |

### Archive Formats

- **Linux/macOS**: `.tar.gz`
- **Windows**: `.zip`

### Checksums

SHA256 checksums are generated for all binaries and archives:
- `checksums.txt` - All checksums in one file

## GoReleaser Configuration

### `.goreleaser.yml`

```yaml
version: 2

before:
  hooks:
    - go mod tidy
    - go generate ./...

builds:
  - id: hailow
    main: ./cmd/hailow
    binary: hailow
    env:
      - CGO_ENABLED=0
    goos:
      - linux
      - darwin
      - windows
    goarch:
      - amd64
      - arm64
    ldflags:
      - -s -w
      - -X main.version={{.Version}}
      - -X main.commit={{.Commit}}
      - -X main.date={{.Date}}
      - -X main.builtBy=goreleaser

archives:
  - id: hailow
    format_overrides:
      - goos: windows
        format: zip
    files:
      - LICENSE
      - README.md
      - INSTALL.md

checksum:
  name_template: 'checksums.txt'

changelog:
  sort: asc
  filters:
    exclude:
      - '^docs:'
      - '^test:'
      - '^ci:'
      - '^chore:'
      - Merge pull request
      - Merge branch

release:
  github:
    owner: username
    name: hailow
  draft: false
  prerelease: auto
  name_template: "{{.ProjectName}} {{.Version}}"
  header: |
    ## Agent Configuration Manager {{.Version}}
    
    Install or upgrade using:
    ```bash
    curl -sSL https://raw.githubusercontent.com/username/hailow/main/install.sh | bash
    ```
  footer: |
    ## Checksums
    
    Verify downloads using `checksums.txt`
    
    ## Install
    
    See [INSTALL.md](https://github.com/Harvey-N-Lab/hailow/blob/main/INSTALL.md) for detailed installation instructions.

snapshot:
  name_template: "{{ incpatch .Version }}-next"
```

## GitHub Actions Workflow

### Release Workflow (`.github/workflows/release.yml`)

**Trigger:** Push of version tag (`v*.*.*`)

**Steps:**
1. Checkout code
2. Set up Go
3. Run tests
4. Run GoReleaser
5. Publish to GitHub Releases

**Required Secrets:**
- `GITHUB_TOKEN` (automatically provided)

### Additional Workflows

#### **Test Workflow** (`.github/workflows/test.yml`)
- Runs on every push and PR
- Runs unit tests
- Reports coverage

#### **Validation Workflow** (`.github/workflows/validate.yml`)
- Validates domain structure
- Checks for filename collisions
- Verifies domain metadata

## Installation Script Distribution

### Primary Method: Raw GitHub URL

The install script is accessible via:
```
https://raw.githubusercontent.com/username/hailow/main/scripts/install.sh
```

**One-liner installation:**
```bash
curl -sSL https://raw.githubusercontent.com/username/hailow/main/scripts/install.sh | bash
```

### Script Responsibilities

The `install.sh` script:
1. Detects OS and architecture
2. Determines latest release version
3. Downloads appropriate binary
4. Verifies checksum
5. Installs to `~/.local/bin/` or `/usr/local/bin/`
6. Makes executable
7. Verifies installation

## Release Checklist

Before creating a release:

- [ ] All tests passing
- [ ] CHANGELOG.md updated
- [ ] Version bumped in appropriate files
- [ ] Documentation updated for new features
- [ ] Migration guide created for breaking changes (if applicable)
- [ ] README examples tested
- [ ] Installation script tested on all platforms

Create release:

- [ ] Create and push version tag
- [ ] Verify GitHub Actions workflow succeeds
- [ ] Verify binaries uploaded to GitHub Release
- [ ] Verify checksums present
- [ ] Test installation script with new version
- [ ] Update documentation if needed
- [ ] Announce release (if applicable)

## Version Upgrade Path

### Backward Compatibility

Minor and patch versions maintain backward compatibility:
- Config file format remains compatible
- CLI flags remain compatible
- Manifest format remains compatible
- Domain structure remains compatible

### Breaking Changes

Major versions may introduce breaking changes. Migration guide required:

**Example: v1 → v2 Migration**
- Document config file format changes
- Provide migration script if possible
- Support gradual migration period
- Clear upgrade instructions

## Release Notes Format

Release notes are auto-generated from commit messages but should be edited for clarity:

```markdown
## v1.1.0 (2026-03-28)

### Features
- Add support for custom domain repositories
- Implement `hailow config` command
- Add shell completion for bash/zsh/fish

### Bug Fixes
- Fix file permission issues on Windows
- Resolve manifest merge conflicts
- Correct path handling on Windows

### Documentation
- Add platform-specific installation guides
- Improve troubleshooting section
- Add examples for all commands

### Internal
- Refactor installer package
- Improve test coverage to 85%
- Update dependencies

**Full Changelog**: https://github.com/Harvey-N-Lab/hailow/compare/v1.0.0...v1.1.0
```

## Binary Distribution

### Install Locations

**Linux/macOS:**
- User install: `~/.local/bin/hailow`
- System install: `/usr/local/bin/hailow`
- Package managers: TBD

**Windows:**
- User install: `%USERPROFILE%\AppData\Local\Programs\hailow\hailow.exe`
- System install: `C:\Program Files\hailow\hailow.exe`
- Package managers: TBD (Chocolatey, Scoop)

### Verification

Users can verify downloads:

```bash
# Download checksum file
curl -LO https://github.com/Harvey-N-Lab/hailow/releases/download/v1.0.0/checksums.txt

# Verify binary
sha256sum -c checksums.txt --ignore-missing
```

## Update Mechanism

### Check for Updates

CLI can check for new versions:
```bash
hailow version --check
```

Output:
```
Current version: 1.0.0
Latest version: 1.1.0

A new version is available!

To upgrade:
  curl -sSL https://raw.githubusercontent.com/username/hailow/main/scripts/install.sh | bash

Or manual download:
  https://github.com/Harvey-N-Lab/hailow/releases/latest
```

### Auto-Update (Future)

Potential self-update command:
```bash
hailow update
```

This would:
1. Check for latest version
2. Download appropriate binary
3. Replace current binary
4. Verify installation

## Rollback Strategy

If a release has critical issues:

1. **Mark release as pre-release** (hide from latest)
2. **Create hotfix release** with fix
3. **Document known issues** in release notes
4. **Provide rollback instructions:**
   ```bash
   # Install specific version
   curl -sSL https://raw.githubusercontent.com/username/hailow/main/scripts/install.sh | bash -s -- v1.0.0
   ```

## Testing Releases

### Pre-Release Testing

Before tagging a release:

1. **Build locally:**
   ```bash
   goreleaser build --snapshot --clean
   ```

2. **Test binaries:**
   ```bash
   ./dist/hailow_linux_amd64_v1/hailow version
   ./dist/hailow_linux_amd64_v1/hailow list domains
   ./dist/hailow_linux_amd64_v1/hailow install devops-engineer --dry-run
   ```

3. **Test on multiple platforms:**
   - Linux (Ubuntu, Fedora)
   - macOS (Intel, Apple Silicon)
   - Windows (10, 11)

### Beta Releases

For testing major changes:

```bash
git tag -a v2.0.0-beta.1 -m "Beta release for v2.0.0"
git push origin v2.0.0-beta.1
```

Beta releases:
- Marked as pre-release in GitHub
- Not considered "latest"
- Used for community testing
- Multiple beta versions possible: `beta.1`, `beta.2`, etc.

## Release Frequency

- **Patch releases**: As needed for critical bugs (weekly if needed)
- **Minor releases**: Monthly or when features are ready
- **Major releases**: Yearly or for significant breaking changes

## Communication

Release announcements:
- GitHub Release notes (automatic)
- Repository README badge (automatic)
- Social media/blog (manual, for major releases)
- Discussion/Discord (if applicable)

## Metrics

Track release metrics:
- Download counts per platform
- Installation success rate
- Update adoption rate
- Issue reports per version
- User feedback

Use GitHub Release API and Analytics to monitor adoption.
