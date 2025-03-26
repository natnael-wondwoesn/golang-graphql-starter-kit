package tui

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*
var templatesFS embed.FS

// GenerateProject creates a new project based on the user's configuration
func GenerateProject(projectPath string, config *ProjectConfig) error {
	// Create the project directory
	if err := os.MkdirAll(projectPath, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Create base directory structure
	if err := createDirectoryStructure(projectPath); err != nil {
		return err
	}

	// Process and copy templates
	if err := processTemplates(projectPath, config); err != nil {
		return err
	}

	return nil
}

// createDirectoryStructure sets up the basic directory structure for the project
func createDirectoryStructure(projectPath string) error {
	dirs := []string{
		"cmd/server",
		"internal/auth",
		"internal/models",
		"internal/services",
		"pkg/utils",
		"graph/schema",
		"graph/resolvers",
		"db/migrations",
		"config",
	}

	for _, dir := range dirs {
		fullPath := filepath.Join(projectPath, dir)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	return nil
}

// processTemplates processes and copies all template files to the project
func processTemplates(projectPath string, config *ProjectConfig) error {
	// Templates that need to be processed with the config
	templateFiles := map[string]string{
		"templates/go.mod.tmpl":                "go.mod",
		"templates/main.go.tmpl":               "cmd/server/main.go",
		"templates/config.go.tmpl":             "config/config.go",
		"templates/config.yaml.tmpl":           "config/config.yaml",
		"templates/db_connection.go.tmpl":      "internal/models/db.go",
		"templates/user_model.go.tmpl":         "internal/models/user.go",
		"templates/user_service.go.tmpl":       "internal/services/user_service.go",
		"templates/schema.graphqls.tmpl":       "graph/schema/schema.graphqls",
		"templates/Dockerfile.tmpl":            "Dockerfile",
		"templates/docker-compose.yml.tmpl":    "docker-compose.yml",
		"templates/Makefile.tmpl":              "Makefile",
		"templates/README.md.tmpl":             "README.md",
		"templates/gqlgen.yml.tmpl":            "gqlgen.yml",
		"templates/gitignore.tmpl":             ".gitignore",
	}

	// Add conditional templates based on config
	if config.Auth == "JWT" {
		templateFiles["templates/jwt.go.tmpl"] = "internal/auth/jwt.go"
		templateFiles["templates/middleware.go.tmpl"] = "internal/auth/middleware.go"
	} else if config.Auth == "OAuth" {
		templateFiles["templates/oauth.go.tmpl"] = "internal/auth/oauth.go"
		templateFiles["templates/middleware.go.tmpl"] = "internal/auth/middleware.go"
	}

	// Process and write each template
	for tmplPath, destPath := range templateFiles {
		if err := processTemplate(tmplPath, filepath.Join(projectPath, destPath), config); err != nil {
			return fmt.Errorf("failed to process template %s: %w", tmplPath, err)
		}
	}

	return nil
}

// processTemplate processes a single template file and writes it to the destination
func processTemplate(tmplPath, destPath string, config *ProjectConfig) error {
	// Read template content
	content, err := templatesFS.ReadFile(tmplPath)
	if err != nil {
		// Skip if template doesn't exist (for conditional templates)
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	// Create template
	tmpl, err := template.New(filepath.Base(tmplPath)).Parse(string(content))
	if err != nil {
		return err
	}

	// Create the destination file
	destDir := filepath.Dir(destPath)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	file, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Execute the template
	templateData := map[string]interface{}{
		"Config":    config,
		"HasFeature": func(feature string) bool {
			for _, f := range config.Features {
				if f == feature {
					return true
				}
			}
			return false
		},
		"LowerCase": strings.ToLower,
	}

	return tmpl.Execute(file, templateData)
}