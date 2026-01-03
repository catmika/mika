package prompt

// ProjectType represents the type of project (frontend/backend)
type ProjectType string

const (
	ProjectTypeFrontend ProjectType = "frontend"
	ProjectTypeBackend  ProjectType = "backend"
)

// Framework represents the selected framework
type Framework string

// Frontend frameworks
const (
	FrameworkReact   Framework = "react"
	FrameworkVue     Framework = "vue"
	FrameworkVanilla Framework = "vanilla"
)

// Backend frameworks
const (
	FrameworkExpress Framework = "express"
	FrameworkFastify Framework = "fastify"
	FrameworkNest    Framework = "nest"
)

// RenderingMode represents how the app should render (for frontend)
type RenderingMode string

const (
	RenderingSPA    RenderingMode = "spa"
	RenderingSSR    RenderingMode = "ssr"
	RenderingStatic RenderingMode = "static"
)

// Addon represents optional tooling
type Addon string

const (
	AddonESLint     Addon = "eslint"
	AddonPrettier   Addon = "prettier"
	AddonHusky      Addon = "husky"
	AddonTypeScript Addon = "typescript"
)

// Config holds all user selections from the prompts
type Config struct {
	GuidanceSwitch bool
	Name           string
	Type           ProjectType
	Framework      Framework
	Rendering      RenderingMode // only for frontend
	Addons         []Addon
}
