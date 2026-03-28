package fetcher

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Fetcher fetches domain configurations from a source
type Fetcher interface {
	Fetch() (localPath string, cleanup func(), err error)
}

// GetFetcher returns appropriate fetcher for the source
func GetFetcher(source string) Fetcher {
	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") || strings.HasPrefix(source, "git@") {
		return &GitFetcher{URL: source}
	}
	return &LocalFetcher{Path: source}
}

// GitFetcher fetches from a git repository
type GitFetcher struct {
	URL string
}

func (g *GitFetcher) Fetch() (string, func(), error) {
	tmpDir, err := os.MkdirTemp("", "hailow-*")
	if err != nil {
		return "", nil, fmt.Errorf("failed to create temp dir: %w", err)
	}

	cleanup := func() {
		os.RemoveAll(tmpDir)
	}

	// Clone the repository
	cmd := exec.Command("git", "clone", "--depth", "1", g.URL, tmpDir)
	output, err := cmd.CombinedOutput()
	if err != nil {
		cleanup()
		return "", nil, fmt.Errorf("failed to clone repository: %w\nOutput: %s", err, string(output))
	}

	return tmpDir, cleanup, nil
}

// LocalFetcher uses a local directory
type LocalFetcher struct {
	Path string
}

func (l *LocalFetcher) Fetch() (string, func(), error) {
	absPath, err := filepath.Abs(l.Path)
	if err != nil {
		return "", nil, fmt.Errorf("invalid path: %w", err)
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return "", nil, fmt.Errorf("path does not exist: %s", absPath)
	}

	// No cleanup needed for local paths
	return absPath, func() {}, nil
}
