package installer

import (
	"time"
)

// Manifest tracks what has been installed in a workspace
type Manifest struct {
	Version     string          `yaml:"version"`
	Platform    string          `yaml:"platform"`
	InstalledAt time.Time       `yaml:"installed_at"`
	Source      SourceInfo      `yaml:"source"`
	Domains     []DomainInstall `yaml:"domains"`
}

// SourceInfo contains information about the source
type SourceInfo struct {
	Type   string `yaml:"type"`   // "git" or "local"
	URL    string `yaml:"url"`    // for git
	Commit string `yaml:"commit"` // for git
	Path   string `yaml:"path"`   // for local
}

// DomainInstall tracks an installed domain
type DomainInstall struct {
	Name    string   `yaml:"name"`
	Version string   `yaml:"version"`
	Files   []string `yaml:"files"`
}

// NewManifest creates a new manifest
func NewManifest(platform string) *Manifest {
	return &Manifest{
		Version:     "1.0",
		Platform:    platform,
		InstalledAt: time.Now(),
		Domains:     []DomainInstall{},
	}
}

// AddDomain adds a domain to the manifest
func (m *Manifest) AddDomain(domain DomainInstall) {
	m.Domains = append(m.Domains, domain)
}

// GetDomain retrieves a domain from the manifest
func (m *Manifest) GetDomain(name string) *DomainInstall {
	for i := range m.Domains {
		if m.Domains[i].Name == name {
			return &m.Domains[i]
		}
	}
	return nil
}

// RemoveDomain removes a domain from the manifest
func (m *Manifest) RemoveDomain(name string) bool {
	for i, domain := range m.Domains {
		if domain.Name == name {
			m.Domains = append(m.Domains[:i], m.Domains[i+1:]...)
			return true
		}
	}
	return false
}
