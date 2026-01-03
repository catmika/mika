package prompt

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

var (
	guidanceSwitch bool
)

func Run() (Config, error) {
	config := Config{GuidanceSwitch: guidanceSwitch}

	initialGroup := huh.NewGroup(
		huh.NewConfirm().Title("Do you want guidance or prefer to choose stack manually?").Affirmative("Guidance").Negative("Manual").Value(&config.GuidanceSwitch),
	)

	form := huh.NewForm(initialGroup)

	err := form.Run()

	if !config.GuidanceSwitch {
		fmt.Println("So you are senior now huh?")
	} else {
		fmt.Println("We've got ourselves a newbee")
	}

	return config, err
}
