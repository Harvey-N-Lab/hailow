package config

import (
	"os"
	"path/filepath"
)

// Config represents the CLI configuration
type Config struct {
	Source   SourceConfig   `yaml:"source"`
	Platform string         `yaml:"platform"`
	Install  InstallConfig  `yaml:"install"`
	Paths    PathsConfig    `yaml:"paths"`
	Logging  LoggingConfig  `yaml:"logging"`
}

// SourceConfig defines the source of domain configurations
type SourceConfig struct {
	Type   string `yaml:"type"`   // "git" or "local"
	URL    string `yaml:"url"`    // for git type
	Branch string `yaml:"branch"` // for git type
	Path   string `yaml:"path"`   // for local type
}

// InstallConfig defines installation behavior
type InstallConfig struct {
	OnConflict     string `yaml:"on_conflict"`      // "skip", "overwrite", "prompt"
	IncludeGeneral bool   `yaml:"include_general"`
	Backup         bool   `yaml:"backup"`
}

// PathsConfig defines paths used by the CLI
type PathsConfig struct {
	CacheDir        string `yaml:"cache_dir"`
	CacheRetention  int    `yaml:"cache_retention"` // hours
}

// LoggingConfig defines logging configuration
type LoggingConfig struct {
	Level string `yaml:"level"` // "debug", "info", "warn", "error"
	File  string `yaml:"file"`  // optional log file
}

// DefaultConfig returns default configuration
func DefaultConfig() *Config {
	homeDir, _ := os.UserHomeDir()
	
	return &Config{
		Source: SourceConfig{
			Type:   "git",
			URL:    "https://github.com/username/agent-configs",
			Branch: "main",
		},
		Platform: "roo",
		Install: InstallConfig{
			OnConflict:     "skip",
			IncludeGeneral: true,
			Backup:         true,
		},
		Paths: PathsConfig{
			CacheDir:       filepath.Join(homeDir, ".hailow", "cache"),
			CacheRetention: 24,
		},
		Logging: LoggingConfig{
			Level: "info",
		},
	}
}

// GetConfigPath returns the path to the config file
func GetConfigPath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".hailow", "config.yaml")
}
