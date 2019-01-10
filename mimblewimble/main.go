package main

import (
	"fmt"

	"github.com/madhive/sandbox/mimblewimble/version"
)

func main() {
	fmt.Printf("Build: %v (%v)\r\n", version.BuildVersion, version.BuildTimestamp)
}
