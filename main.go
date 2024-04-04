package main

import (
	"fmt"

	"github.com/claeusdev/launch-lever-go/launchlever"
)

func main() {
	flagJson := `[
		{
			"Name": "DarkMode",
			"Description": "Enables the dark mode theme across the application",
			"Status": "Enabled"
		},
		{
			"Name": "FeatureX",
			"Description": "Activates Feature X in the application",
			"Status": "Disabled"
		}
	]`

	getFlagByName := launchlever.LaunchLever(flagJson)
	var darkMode launchlever.Toggle
	if toggle, found := getFlagByName("DarkMode"); found {
		darkMode = toggle
	}
	fmt.Println("Hello, 世界", darkMode.Status)
}
