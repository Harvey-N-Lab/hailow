package domain

// Domain represents an agent configuration domain
type Domain struct {
	ID          string   `yaml:"id"`
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Version     string   `yaml:"version"`
	Categories  []string `yaml:"categories"`
}

// DomainRegistry contains all available domains
type DomainRegistry struct {
	Version string            `yaml:"version"`
	Domains []Domain          `yaml:"domains"`
	Aliases map[string]string `yaml:"aliases"`
}

// GetAllDomains returns all available domains
func GetAllDomains() []string {
	return []string{
		"devops-engineer",
		"python-backend-engineer",
		"js-ts-software-engineer",
		"data-engineer",
		"machine-learning-engineer",
	}
}

// GetDomainPath returns the source path for a domain
func GetDomainPath(domain string) string {
	return "src/" + domain
}

// ResolveDomainAlias resolves a domain alias to its full name
func ResolveDomainAlias(alias string) string {
	aliases := map[string]string{
		"devops":           "devops-engineer",
		"ops":              "devops-engineer",
		"python":           "python-backend-engineer",
		"py":               "python-backend-engineer",
		"python-backend":   "python-backend-engineer",
		"jsts":             "js-ts-software-engineer",
		"js":               "js-ts-software-engineer",
		"ts":               "js-ts-software-engineer",
		"javascript":       "js-ts-software-engineer",
		"typescript":       "js-ts-software-engineer",
		"data":             "data-engineer",
		"de":               "data-engineer",
		"mleng":            "machine-learning-engineer",
		"ml":               "machine-learning-engineer",
		"machine-learning": "machine-learning-engineer",
	}

	if fullName, ok := aliases[alias]; ok {
		return fullName
	}
	return alias
}
