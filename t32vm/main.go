package main

import (
	"fmt"

	"github.com/madhive/sandbox/t32vm/version"
)

func main() {
	fmt.Printf("Build: %v (%v)\r\n", version.BuildVersion, version.BuildTimestamp)
}
