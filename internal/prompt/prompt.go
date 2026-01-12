package prompt

import (
	"github.com/charmbracelet/huh"
)

func Run() (Config, error) {
	config := Config{}

	// Group 1: Initial questions
	initialGroup := huh.NewGroup(
		huh.NewConfirm().
			Title("Do you want guidance or prefer to choose stack manually?").
			Affirmative("Guidance").
			Negative("Manual").
			Value(&config.GuidanceSwitch),

		huh.NewSelect[ProjectType]().
			Title("What are you building?").
			Options(
				huh.NewOption("Frontend", ProjectTypeFrontend),
				huh.NewOption("Backend", ProjectTypeBackend),
			).
			Value(&config.Type),
	)

	// Group 2: Frontend-specific questions
	frontendGroup := huh.NewGroup(
		huh.NewSelect[Framework]().
			Title("Which frontend framework?").
			Options(toFrameworkOptions(frontendFrameworks)...).
			Value(&config.Framework),
	).WithHideFunc(func() bool {
		return config.Type != ProjectTypeFrontend
	})

	// Group 3: Backend-specific questions
	backendGroup := huh.NewGroup(
		huh.NewSelect[Framework]().
			Title("Which backend framework?").
			Options(toFrameworkOptions(backendFrameworks)...).
			Value(&config.Framework),
	).WithHideFunc(func() bool {
		return config.Type != ProjectTypeBackend
	})

	// Group 4: Addons (shown for both frontend and backend)
	addonsGroup := huh.NewGroup(
		huh.NewMultiSelect[Addon]().
			Title("Additional tools (optional)").
			Options(toAddonOptions(addons)...).
			Value(&config.Addons),
	)

	// Single form with all groups
	form := huh.NewForm(
		initialGroup,
		frontendGroup,
		backendGroup,
		addonsGroup,
	)

	err := form.Run()

	return config, err
}
