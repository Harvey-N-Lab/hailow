package installer

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// Installer handles domain installation
type Installer struct {
	SourcePath     string
	WorkspacePath  string
	Platform       Platform
	Force          bool
	IncludeGeneral bool
}

// Install installs domains to the workspace
func (i *Installer) Install(domains []string) error {
	// Create workspace directory if it doesn't exist
	if err := os.MkdirAll(i.WorkspacePath, 0755); err != nil {
		return fmt.Errorf("failed to create workspace: %w", err)
	}

	mapping := GetPlatformMapping(i.Platform)

	// Create platform root directory
	platformDir := filepath.Join(i.WorkspacePath, mapping.RootDir)
	if err := os.MkdirAll(platformDir, 0755); err != nil {
		return fmt.Errorf("failed to create platform directory: %w", err)
	}

	// Create subdirectories
	dirs := []string{
		filepath.Join(i.WorkspacePath, mapping.AgentsDir), // For Claude's agents/ subfolder
		filepath.Join(i.WorkspacePath, mapping.RulesDir),
		filepath.Join(i.WorkspacePath, mapping.SkillsDir),
		filepath.Join(i.WorkspacePath, mapping.CommandsDir),
		filepath.Join(i.WorkspacePath, mapping.ContextsDir),
	}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Install general rules if enabled
	if i.IncludeGeneral {
		if err := i.installGeneral(mapping); err != nil {
			return fmt.Errorf("failed to install general rules: %w", err)
		}
	}

	// Install each domain
	for _, domain := range domains {
		if err := i.installDomain(domain, mapping); err != nil {
			return fmt.Errorf("failed to install domain %s: %w", domain, err)
		}
	}

	// Copy root guidance file
	if err := i.copyRootGuidance(mapping); err != nil {
		return fmt.Errorf("failed to copy root guidance: %w", err)
	}

	// Create manifest
	if err := i.createManifest(domains); err != nil {
		return fmt.Errorf("failed to create manifest: %w", err)
	}

	return nil
}

func (i *Installer) installGeneral(mapping PlatformMapping) error {
	generalRulesPath := filepath.Join(i.SourcePath, "src/general/rules")

	// Check if general rules exist
	if _, err := os.Stat(generalRulesPath); os.IsNotExist(err) {
		return nil // No general rules, skip
	}

	entries, err := os.ReadDir(generalRulesPath)
	if err != nil {
		return err
	}

	targetDir := filepath.Join(i.WorkspacePath, mapping.RulesDir)

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
			continue
		}

		src := filepath.Join(generalRulesPath, entry.Name())
		dst := filepath.Join(targetDir, entry.Name())

		if err := i.copyFile(src, dst); err != nil {
			return fmt.Errorf("failed to copy %s: %w", entry.Name(), err)
		}
	}

	return nil
}

func (i *Installer) installDomain(domain string, mapping PlatformMapping) error {
	domainPath := filepath.Join(i.SourcePath, "src", domain)

	// Check if domain exists
	if _, err := os.Stat(domainPath); os.IsNotExist(err) {
		return fmt.Errorf("domain not found: %s", domain)
	}

	// Install agents (go to platform root)
	if err := i.copyDomainFiles(
		filepath.Join(domainPath, "agents"),
		filepath.Join(i.WorkspacePath, mapping.AgentsDir),
	); err != nil {
		return err
	}

	// Install rules
	if err := i.copyDomainFiles(
		filepath.Join(domainPath, "rules"),
		filepath.Join(i.WorkspacePath, mapping.RulesDir),
	); err != nil {
		return err
	}

	// Install skills (copy entire skill directories)
	if err := i.copySkillsDirectory(
		filepath.Join(domainPath, "skills"),
		filepath.Join(i.WorkspacePath, mapping.SkillsDir),
	); err != nil {
		return err
	}

	// Install commands
	if err := i.copyDomainFiles(
		filepath.Join(domainPath, "commands"),
		filepath.Join(i.WorkspacePath, mapping.CommandsDir),
	); err != nil {
		return err
	}

	// Install contexts placeholder
	if err := i.copyDomainFiles(
		filepath.Join(domainPath, "contexts"),
		filepath.Join(i.WorkspacePath, mapping.ContextsDir),
	); err != nil {
		return err
	}

	return nil
}

func (i *Installer) copyDomainFiles(srcDir, dstDir string) error {
	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		return nil // Directory doesn't exist, skip
	}

	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue // Skip directories for now
		}

		src := filepath.Join(srcDir, entry.Name())
		dst := filepath.Join(dstDir, entry.Name())

		if err := i.copyFile(src, dst); err != nil {
			return err
		}
	}

	return nil
}

func (i *Installer) copySkillsDirectory(srcDir, dstDir string) error {
	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		return nil
	}

	// Copy all contents recursively
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Calculate relative path
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dstDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, 0755)
		}

		return i.copyFile(path, dstPath)
	})
}

func (i *Installer) copyFile(src, dst string) error {
	// Check if destination exists
	if _, err := os.Stat(dst); err == nil && !i.Force {
		// File exists and force is not enabled, skip
		return nil
	}

	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

func (i *Installer) copyRootGuidance(mapping PlatformMapping) error {
	src := filepath.Join(i.SourcePath, "AGENT_WORKFLOW.md")
	dst := filepath.Join(i.WorkspacePath, mapping.RootGuidance)

	return i.copyFile(src, dst)
}

func (i *Installer) createManifest(domains []string) error {
	manifestDir := filepath.Join(i.WorkspacePath, ".hailow")
	if err := os.MkdirAll(manifestDir, 0755); err != nil {
		return err
	}

	manifest := NewManifest(string(i.Platform))
	manifest.InstalledAt = time.Now()
	manifest.Source = SourceInfo{
		Type: "local",
		Path: i.SourcePath,
	}

	for _, domain := range domains {
		manifest.AddDomain(DomainInstall{
			Name:    domain,
			Version: "1.0.0",
			Files:   []string{}, // Would be populated with actual file list
		})
	}

	// Write manifest (simplified - would use YAML in real implementation)
	manifestPath := filepath.Join(manifestDir, "manifest.txt")
	content := fmt.Sprintf("Platform: %s\n", manifest.Platform)
	content += fmt.Sprintf("Installed: %s\n", manifest.InstalledAt.Format(time.RFC3339))
	content += fmt.Sprintf("Domains: %v\n", domains)

	return os.WriteFile(manifestPath, []byte(content), 0644)
}
