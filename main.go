package main

import (
	"fmt"

	"github.com/claeusdev/launch-lever-go/launchlever"
)

func main() {
	// LaunchLever() returns a closure, 
	// when passed the name of a feature toggle returns true 
	// when the status of the toggle is on
	isEnabled := launchlever.LaunchLever("./testdata/test.json")
	dark_state_enabled := isEnabled("pfx_123")
	fmt.Println("Hello, 世界", dark_state_enabled)
}
