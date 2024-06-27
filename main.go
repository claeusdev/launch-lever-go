package main

import (
	"fmt"

	"github.com/claeusdev/launch-lever-go/launchlever"
)

func main() {
	isEnabled := launchlever.LaunchLever("./testdata/test.json")
	dark_state_enabled := isEnabled("pfx_123")
	fmt.Println("Hello, 世界", dark_state_enabled)
}
