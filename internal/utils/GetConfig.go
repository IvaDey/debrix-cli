package utils

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

const configFileName = ".debrix.yml"

var config *Config

func ReadConfig(atPath string) *Config {
	config = &Config{
		Pattern:  []string{"todo"},
		OutFile:  "TODO.md",
		Language: "en",
		Layout:   "plain",
		Exclude: []string{
			".git",         // Git metadata
			".idea",        // JetBrains IDE configuration
			".vscode",      // VSCode configuration
			"node_modules", // Node.js dependencies
			"dist",         // Compiled distribution files
			"build",        // Build artifacts
			"out",          // Output directory
			"bin",          // Compiled binaries
			"vendor",       // Vendored dependencies (Go)
			"third_party",  // External libraries
			"venv",         // Python virtual environment
			"__pycache__",  // Python cache files
			"target",       // Build artifacts (e.g. Rust, Java)
			"coverage",     // Code coverage reports
			"cache",
		},
		LinkTemplate: "{{filePath}}:{{lineNumber}}",
		DbPath:       ".debrix.db",
	}

	// todo: Also exclude the same content as defined in project .gitignore

	fromFile, err := os.ReadFile(filepath.Join(atPath, configFileName))
	if err == nil {
		err = yaml.Unmarshal(fromFile, config)
	}

	return config
}

func GetConfig() *Config {
	if config == nil {
		panic("Config was not initialized")
	}
	return config
}

type Config struct {
	Pattern      []string `yaml:"pattern"`
	OutFile      string   `yaml:"outFile"`
	Language     string   `yaml:"language"`
	Layout       string   `yaml:"layout"`
	Exclude      []string `yaml:"exclude"`
	LinkTemplate string   `yaml:"linkTemplate"`
	DbPath       string   `yaml:"dbPath"`
}
