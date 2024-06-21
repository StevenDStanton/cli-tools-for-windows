package common

import (
	"fmt"
)

const Version = "v0.1.0"

func PrintVersion(tool string, version string) {
	fmt.Printf(`%s: Version %s

Copyright 2024 The Simple Dev

Author:         Steven Stanton
License:        MIT - No Warranty
Author Github:  https//github.com/StevenDStanton
Project Github: https://github.com/StevemStanton/ltfw

Part of my CLI Tools for Windows project.`, tool, version)
}
