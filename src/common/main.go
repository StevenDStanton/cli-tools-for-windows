package common

import (
	"fmt"
	"os"
)

func PrintVersion(tool string, version string) {
	fmt.Printf(`%sVersion %s

Copyright 2024 The Simple Dev

Author:         Steven Stanton
License:        MIT - No Warranty
Author Github:  https//github.com/StevenDStanton
Project Github: https://github.com/StevemStanton/ltfw

Part of my CLI Tools for Windows project.`, tool, version)
	os.Exit(0)
}
