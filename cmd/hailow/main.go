package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Harvey-N-Lab/hailow/internal/installer"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "hailow",
		Short: "Hailow - AI Agent Configuration Manager",
		Long: `Hailow helps you install and manage AI agent configurations
for multiple engineering domains (DevOps, Backend, Frontend, Data, ML).`,
		Version: version,
	}

	// Global flags
	rootCmd.PersistentFlags().String("config", "", "config file (default is $HOME/.hailow/config.yaml)")
	rootCmd.PersistentFlags().Bool("verbose", false, "enable verbose logging")
	rootCmd.PersistentFlags().Bool("quiet", false, "suppress non-error output")

	// Add commands
	rootCmd.AddCommand(installCmd())
	rootCmd.AddCommand(listCmd())
	rootCmd.AddCommand(removeCmd())
	rootCmd.AddCommand(upgradeCmd())
	rootCmd.AddCommand(configCmd())
	rootCmd.AddCommand(versionCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func installCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install [domain...] [directory]",
		Short: "Install domain configurations into workspace",
		Long: `Install one or more domain configurations into a workspace directory.

The directory can be specified as:
  1. A flag: --workspace /path/to/workspace
  2. The last argument: hailow install devops-engineer /path/to/workspace
  3. Default to current directory if not specified

Examples:
  hailow install devops-engineer
  hailow install devops-engineer python-backend-engineer
  hailow install devops-engineer --workspace /path/to/project
  hailow install devops-engineer /path/to/project
  hailow install --all /path/to/project`,
		Args: cobra.MinimumNArgs(0),
		RunE: runInstall,
	}

	cmd.Flags().String("platform", "roo", "target platform: roo or claude")
	cmd.Flags().String("workspace", "", "target workspace directory")
	cmd.Flags().String("source", "", "source repository or local path")
	cmd.Flags().Bool("all", false, "install all available domains")
	cmd.Flags().Bool("dry-run", false, "show what would be installed")
	cmd.Flags().Bool("force", false, "overwrite existing files")
	cmd.Flags().Bool("no-general", false, "skip installing general rules")

	return cmd
}

func runInstall(cmd *cobra.Command, args []string) error {
	platformStr, _ := cmd.Flags().GetString("platform")
	workspaceFlag, _ := cmd.Flags().GetString("workspace")
	source, _ := cmd.Flags().GetString("source")
	all, _ := cmd.Flags().GetBool("all")
	dryRun, _ := cmd.Flags().GetBool("dry-run")
	force, _ := cmd.Flags().GetBool("force")
	noGeneral, _ := cmd.Flags().GetBool("no-general")

	// Parse arguments: domains and optional directory path
	var domains []string
	var workspace string

	if all {
		// All domains, check if last arg is a directory
		if len(args) > 0 && !isDomainName(args[len(args)-1]) {
			workspace = args[len(args)-1]
		}
		domains = []string{
			"devops-engineer",
			"python-backend-engineer",
			"js-ts-software-engineer",
			"data-engineer",
			"machine-learning-engineer",
		}
	} else {
		if len(args) == 0 {
			return fmt.Errorf("specify at least one domain or use --all")
		}

		// Check if last argument is a directory path (not a domain name)
		if len(args) > 1 && !isDomainName(args[len(args)-1]) {
			// Last arg is directory
			domains = args[:len(args)-1]
			workspace = args[len(args)-1]
		} else {
			// All args are domains
			domains = args
		}
	}

	// Workspace priority: flag > argument > current directory
	if workspaceFlag != "" {
		workspace = workspaceFlag
	} else if workspace == "" {
		workspace = "."
	}

	// Convert to absolute path
	workspace, err := filepath.Abs(workspace)
	if err != nil {
		return fmt.Errorf("invalid workspace path: %w", err)
	}

	// Determine source path
	if source == "" {
		// Default to current directory (assuming we're in the hailow repo)
		source, _ = os.Getwd()
	}

	fmt.Printf("Installing domains to workspace: %s\n", workspace)
	fmt.Printf("Platform: %s\n", platformStr)
	fmt.Printf("Source: %s\n", source)

	if dryRun {
		fmt.Println("\n[DRY-RUN MODE - No changes will be made]")
	}

	fmt.Println("\nDomains to install:")
	for _, domain := range domains {
		fmt.Printf("  - %s\n", domain)
	}

	if dryRun {
		return nil
	}

	// Perform actual installation
	platform := installer.Platform(platformStr)
	inst := &installer.Installer{
		SourcePath:     source,
		WorkspacePath:  workspace,
		Platform:       platform,
		Force:          force,
		IncludeGeneral: !noGeneral,
	}

	if err := inst.Install(domains); err != nil {
		return fmt.Errorf("installation failed: %w", err)
	}

	fmt.Println("\n✓ Installation complete!")
	fmt.Printf("✓ Files installed to: %s\n", workspace)

	// Show what was created
	mapping := installer.GetPlatformMapping(platform)
	fmt.Println("\nInstalled:")
	fmt.Printf("  - %s/ (agent configurations)\n", mapping.RootDir)
	fmt.Printf("  - %s (workflow guidance)\n", mapping.RootGuidance)
	fmt.Println("  - .hailow/manifest.txt (installation manifest)")

	return nil
}

// isDomainName checks if a string is a known domain name or alias
func isDomainName(s string) bool {
	knownDomains := map[string]bool{
		"devops-engineer":           true,
		"devops":                    true,
		"ops":                       true,
		"python-backend-engineer":   true,
		"python":                    true,
		"py":                        true,
		"python-backend":            true,
		"js-ts-software-engineer":   true,
		"jsts":                      true,
		"js":                        true,
		"ts":                        true,
		"javascript":                true,
		"typescript":                true,
		"data-engineer":             true,
		"data":                      true,
		"de":                        true,
		"machine-learning-engineer": true,
		"mleng":                     true,
		"ml":                        true,
		"machine-learning":          true,
	}
	return knownDomains[s]
}

func listCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [type]",
		Short: "List available domains or installed configurations",
		Long: `List available domains, supported platforms, or installed configurations.

Types: domains, platforms, installed

Examples:
  hailow list domains
  hailow list installed`,
		Args: cobra.MaximumNArgs(1),
		RunE: runList,
	}

	cmd.Flags().String("source", "", "source repository or local path")
	cmd.Flags().String("workspace", ".", "workspace directory")
	cmd.Flags().String("format", "table", "output format: table, json, yaml")

	return cmd
}

func runList(cmd *cobra.Command, args []string) error {
	listType := "domains"
	if len(args) > 0 {
		listType = args[0]
	}

	switch listType {
	case "domains":
		fmt.Println("Available Domains:\n")
		fmt.Println("Name                           Description")
		fmt.Println("----                           -----------")
		fmt.Println("devops-engineer                Infrastructure, CI/CD, and deployment")
		fmt.Println("python-backend-engineer        Python backend development")
		fmt.Println("js-ts-software-engineer        JavaScript/TypeScript development")
		fmt.Println("data-engineer                  Data pipelines and engineering")
		fmt.Println("machine-learning-engineer      ML development and deployment")

	case "platforms":
		fmt.Println("Supported Platforms:\n")
		fmt.Println("- roo       Roo Code")
		fmt.Println("- claude    Claude Code")

	case "installed":
		workspace, _ := cmd.Flags().GetString("workspace")
		fmt.Printf("Installed domains in %s:\n\n", workspace)

		manifestPath := filepath.Join(workspace, ".hailow", "manifest.txt")
		if _, err := os.Stat(manifestPath); err == nil {
			content, _ := os.ReadFile(manifestPath)
			fmt.Println(string(content))
		} else {
			fmt.Println("No domains currently installed.")
			fmt.Println("\nUse 'hailow install <domain>' to install a domain.")
		}

	default:
		return fmt.Errorf("unknown list type: %s (use: domains, platforms, installed)", listType)
	}

	return nil
}

func removeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [domain...]",
		Short: "Remove installed domain configurations",
		Long: `Remove one or more installed domain configurations from workspace.

Examples:
  hailow remove devops-engineer
  hailow remove --all`,
		Args: cobra.MinimumNArgs(0),
		RunE: runRemove,
	}

	cmd.Flags().String("workspace", ".", "target workspace directory")
	cmd.Flags().Bool("all", false, "remove all installed domains")
	cmd.Flags().Bool("dry-run", false, "show what would be removed")

	return cmd
}

func runRemove(cmd *cobra.Command, args []string) error {
	all, _ := cmd.Flags().GetBool("all")
	if !all && len(args) == 0 {
		return fmt.Errorf("specify at least one domain or use --all")
	}

	fmt.Println("Remove functionality not yet implemented")
	return nil
}

func upgradeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade [domain...]",
		Short: "Upgrade installed domains to latest versions",
		Long: `Upgrade one or more installed domains to their latest versions.

Examples:
  hailow upgrade
  hailow upgrade devops-engineer`,
		Args: cobra.MinimumNArgs(0),
		RunE: runUpgrade,
	}

	cmd.Flags().String("workspace", ".", "target workspace directory")
	cmd.Flags().Bool("dry-run", false, "show what would be upgraded")

	return cmd
}

func runUpgrade(cmd *cobra.Command, args []string) error {
	fmt.Println("Upgrade functionality not yet implemented")
	return nil
}

func configCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage CLI configuration",
		Long:  "Manage Hailow CLI configuration",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "show",
		Short: "Show current configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Configuration:")
			fmt.Println("  platform: roo")
			fmt.Println("  source: https://github.com/Harvey-N-Lab/hailow")
			return nil
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set configuration value",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Setting %s = %s\n", args[0], args[1])
			return nil
		},
	})

	return cmd
}

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("hailow version %s\n", version)
			fmt.Printf("Commit: %s\n", commit)
			fmt.Printf("Built: %s\n", date)
		},
	}
}
