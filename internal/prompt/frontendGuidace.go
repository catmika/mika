package prompt

// RenderingMode represents how the app should render (for frontend)
type RenderingMode string

const (
	RenderingSPA    RenderingMode = "spa"
	RenderingSSR    RenderingMode = "ssr"
	RenderingStatic RenderingMode = "static"
)
