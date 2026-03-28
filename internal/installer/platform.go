package installer

// Platform represents a target agent platform
type Platform string

const (
	PlatformRoo    Platform = "roo"
	PlatformClaude Platform = "claude"
)

// PlatformMapping defines how domain files map to platform directories
type PlatformMapping struct {
	RootDir      string // .roo or .claude
	AgentsDir    string // where agent files go
	RulesDir     string // rules subdirectory
	SkillsDir    string // skills subdirectory
	CommandsDir  string // commands subdirectory
	ContextsDir  string // contexts subdirectory
	RootGuidance string // AGENT_WORKFLOW.md inside root dir
}

// GetPlatformMapping returns the directory mapping for a platform
func GetPlatformMapping(platform Platform) PlatformMapping {
	switch platform {
	case PlatformRoo:
		return PlatformMapping{
			RootDir:      ".roo",
			AgentsDir:    ".roo", // Agent files at .roo root
			RulesDir:     ".roo/rules",
			SkillsDir:    ".roo/skills",
			CommandsDir:  ".roo/commands",
			ContextsDir:  ".roo/contexts",
			RootGuidance: ".roo/AGENT_WORKFLOW.md", // Inside .roo folder
		}
	case PlatformClaude:
		return PlatformMapping{
			RootDir:      ".claude",
			AgentsDir:    ".claude/agents", // Agent files in agents subfolder
			RulesDir:     ".claude/rules",
			SkillsDir:    ".claude/skills",
			CommandsDir:  ".claude/commands",
			ContextsDir:  ".claude/contexts",
			RootGuidance: ".claude/AGENT_WORKFLOW.md", // Inside .claude folder
		}
	default:
		// Default to Roo
		return GetPlatformMapping(PlatformRoo)
	}
}

// IsValidPlatform checks if a platform name is valid
func IsValidPlatform(platform string) bool {
	return platform == string(PlatformRoo) || platform == string(PlatformClaude)
}
