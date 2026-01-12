package prompt

// ProjectType represents the type of project (frontend/backend)
type ProjectType string

const (
	ProjectTypeFrontend ProjectType = "frontend"
	ProjectTypeBackend  ProjectType = "backend"
)

type Language string

const (
	LanguageJavaScript Language = "javascript"
	LanguageTypeScript Language = "typescript"
	LanguageGolang     Language = "golang"
	LanguagePython     Language = "python"
	LanguageJava       Language = "java"
	LanguageCSharp     Language = "csharp"
	LanguageRuby       Language = "ruby"
	LanguagePHP        Language = "php"
	LanguageRust       Language = "rust"
)

type Framework string

// Frontend frameworks
const (
	FrameworkReact         Framework = "react"
	FrameworkNext          Framework = "next"
	FrameworkTanstackStart Framework = "tanstackStart"
	FrameworkReactRouter   Framework = "reactRouter"
	FrameworkVue           Framework = "vue"
	FrameworkVanilla       Framework = "vanilla"
)

// Backend frameworks
const (
	FrameworkExpress Framework = "express"
	FrameworkFastify Framework = "fastify"
	FrameworkNest    Framework = "nest"
)

type Database string

const (
	DatabasePostgreSQL Database = "postgresql"
	DatabaseMySQL      Database = "mysql"
	DatabaseSQLite     Database = "sqlite"
	DatabaseMongoDB    Database = "mongodb"
	DatabaseRedis      Database = "redis"
)

type Bundler string

const (
	BundlerWebpack Bundler = "webpack"
	BundlerVite    Bundler = "vite"
	BundlerESBuild Bundler = "esbuild"
	BundlerRollup  Bundler = "rollup"
)

// Addon represents optional tooling
type Addon string

const (
	AddonESLint   Addon = "eslint"
	AddonPrettier Addon = "prettier"
	AddonHusky    Addon = "husky"
	AddonDocker   Addon = "docker"
	AddonGit      Addon = "git"
)

// Config holds all user selections from the prompts
type Config struct {
	GuidanceSwitch bool
	Name           string
	Type           ProjectType
	Framework      Framework
	Addons         []Addon
}
