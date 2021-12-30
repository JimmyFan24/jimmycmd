package app

import (
	"runtime"
	"strings"
)

type SubCommand struct {
	usage    string
	desc     string
	options  CliOptions
	commands []*SubCommand
	runFunc  RunCommandFunc
}

type RunCommandFunc func(args []string) error

func FormatBaseName(basename string) string {
	// Make case-insensitive and strip executable suffix if present
	if runtime.GOOS == "windows" {
		basename = strings.ToLower(basename)
		basename = strings.TrimSuffix(basename, ".exe")
	}

	return basename
}
