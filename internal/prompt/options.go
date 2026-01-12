package prompt

import "github.com/charmbracelet/huh"

// FrameworkOption represents a framework choice with its display label
type FrameworkOption struct {
	Value Framework
	Label string
}

// AddonOption represents an addon choice with its display label
type AddonOption struct {
	Value Addon
	Label string
}

// Frontend frameworks in desired display order
var frontendFrameworks = []FrameworkOption{
	{FrameworkReact, "React"},
	{FrameworkNext, "Next.js"},
	{FrameworkTanstackStart, "TanStack Start"},
	{FrameworkReactRouter, "React Router"},
	{FrameworkVue, "Vue"},
	{FrameworkVanilla, "Vanilla JS"},
}

// Backend frameworks in desired display order
var backendFrameworks = []FrameworkOption{
	{FrameworkExpress, "Express"},
	{FrameworkFastify, "Fastify"},
	{FrameworkNest, "NestJS"},
}

// Available addons in desired display order
var addons = []AddonOption{
	{AddonESLint, "ESLint"},
	{AddonPrettier, "Prettier"},
	{AddonHusky, "Husky (git hooks)"},
	{AddonDocker, "Docker"},
	{AddonGit, "Git"},
}

// toFrameworkOptions converts framework options to huh options
func toFrameworkOptions(frameworks []FrameworkOption) []huh.Option[Framework] {
	opts := make([]huh.Option[Framework], len(frameworks))
	for i, fw := range frameworks {
		opts[i] = huh.NewOption(fw.Label, fw.Value)
	}
	return opts
}

// toAddonOptions converts addon options to huh options
func toAddonOptions(addons []AddonOption) []huh.Option[Addon] {
	opts := make([]huh.Option[Addon], len(addons))
	for i, addon := range addons {
		opts[i] = huh.NewOption(addon.Label, addon.Value)
	}
	return opts
}
