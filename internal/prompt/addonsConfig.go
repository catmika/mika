package prompt

// PrettierConfig holds Prettier formatting options
type PrettierConfig struct {
	PrintWidth       int    // Line wrap at this column (default: 80)
	TabWidth         int    // Spaces per indentation level (default: 2)
	UseTabs          bool   // Use tabs instead of spaces
	Semi             bool   // Add semicolons at end of statements (default: true)
	SingleQuote      bool   // Use single quotes instead of double (default: false)
	TrailingComma    string // "none" | "es5" | "all" (default: "es5")
	BracketSpacing   bool   // Spaces between brackets in object literals (default: true)
	ArrowParens      string // "always" | "avoid" - Include parens around single arrow function params
	EndOfLine        string // "lf" | "crlf" | "cr" | "auto" (default: "lf")
	PrintSemicolons  bool   // Print semicolons at the ends of statements
}

// DefaultPrettierConfig returns popular sensible defaults
func DefaultPrettierConfig() PrettierConfig {
	return PrettierConfig{
		PrintWidth:      80,
		TabWidth:        2,
		UseTabs:         false,
		Semi:            true,
		SingleQuote:     true, // Popular in modern JS/TS projects
		TrailingComma:   "es5",
		BracketSpacing:  true,
		ArrowParens:     "avoid",
		EndOfLine:       "lf",
		PrintSemicolons: true,
	}
}

// ESLintConfig holds ESLint configuration options
type ESLintConfig struct {
	Extends     []string          // Preset configs: "eslint:recommended", "airbnb", "standard", "prettier"
	Parser      string            // "@babel/eslint-parser", "@typescript-eslint/parser"
	ParserOptions map[string]any  // ecmaVersion, sourceType, etc.
	Plugins     []string          // Additional plugins
	Env         map[string]bool   // browser, node, es2021, jest
	Rules       map[string]any    // Custom rule overrides
}

// DefaultESLintConfig returns popular config for modern JS/TS projects
func DefaultESLintConfig(isTypeScript bool) ESLintConfig {
	config := ESLintConfig{
		Extends: []string{"eslint:recommended"},
		ParserOptions: map[string]any{
			"ecmaVersion": "latest",
			"sourceType":  "module",
		},
		Env: map[string]bool{
			"browser": false, // Will be set based on project type
			"node":    false,
			"es2021":  true,
		},
		Rules: map[string]any{
			"no-unused-vars": "warn",
			"no-console":     "warn",
		},
	}

	if isTypeScript {
		config.Parser = "@typescript-eslint/parser"
		config.Extends = append(config.Extends,
			"plugin:@typescript-eslint/recommended",
		)
		config.Plugins = []string{"@typescript-eslint"}
		config.ParserOptions["project"] = "./tsconfig.json"
	}

	return config
}

// HuskyConfig holds git hooks configuration
type HuskyConfig struct {
	PreCommit  []string // Commands to run before commit
	PrePush    []string // Commands to run before push
	CommitMsg  string   // Script to validate commit message
}

// DefaultHuskyConfig returns common git hooks setup
func DefaultHuskyConfig(hasESLint, hasPrettier bool) HuskyConfig {
	config := HuskyConfig{
		PreCommit: []string{},
		PrePush:   []string{},
	}

	// Pre-commit: lint and format staged files
	if hasESLint && hasPrettier {
		config.PreCommit = append(config.PreCommit, "npx lint-staged")
	} else if hasESLint {
		config.PreCommit = append(config.PreCommit, "npm run lint")
	} else if hasPrettier {
		config.PreCommit = append(config.PreCommit, "npm run format")
	}

	// Pre-push: run tests if they exist
	config.PrePush = append(config.PrePush, "npm test")

	// Commit message validation (conventional commits)
	config.CommitMsg = "npx --no -- commitlint --edit $1"

	return config
}

// LintStagedConfig for pre-commit hooks (used with Husky)
type LintStagedConfig map[string][]string

// DefaultLintStagedConfig returns lint-staged configuration
func DefaultLintStagedConfig(hasESLint, hasPrettier bool) LintStagedConfig {
	config := LintStagedConfig{}

	jsFiles := []string{"*.js", "*.jsx", "*.ts", "*.tsx"}

	var commands []string
	if hasESLint {
		commands = append(commands, "eslint --fix")
	}
	if hasPrettier {
		commands = append(commands, "prettier --write")
	}

	for _, pattern := range jsFiles {
		config[pattern] = commands
	}

	return config
}

// DockerConfig holds Docker configuration options
type DockerConfig struct {
	BaseImage       string   // e.g., "node:20-alpine", "node:20-slim"
	WorkDir         string   // Working directory inside container (default: /app)
	Port            int      // Exposed port (default: 3000)
	NodeVersion     string   // Node version: "20", "18", "22"
	UseMultiStage   bool     // Use multi-stage build for smaller images
	InstallCmd      string   // "npm install" or "npm ci"
	BuildCmd        string   // Build command if needed
	StartCmd        string   // Start command: "npm start", "node dist/index.js"
	IncludeDevDeps  bool     // Include devDependencies in final image
}

// DefaultDockerConfig returns sensible Docker defaults
func DefaultDockerConfig(projectType ProjectType) DockerConfig {
	config := DockerConfig{
		NodeVersion:    "20",
		WorkDir:        "/app",
		Port:           3000,
		UseMultiStage:  true,
		InstallCmd:     "npm ci",
		IncludeDevDeps: false,
	}

	if projectType == ProjectTypeFrontend {
		config.BaseImage = "node:20-alpine"
		config.BuildCmd = "npm run build"
		config.StartCmd = "npm run preview" // For Vite, or serve for other bundlers
	} else {
		config.BaseImage = "node:20-alpine"
		config.StartCmd = "node dist/index.js" // Assumes compiled backend
	}

	return config
}
