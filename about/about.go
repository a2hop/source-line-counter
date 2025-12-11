package about

import (
	"os"
	"path/filepath"
)

var ProgramName = "Source Line Counter"
var Version = "1.0.0"
var Description = "A tool to count lines of code in various programming languages."
var License = "MIT License"
var GitUrl = "https://github.com/a2hop/source-line-counter"

func GetBinName() string {
	// get the name dynamiacally of the current binary being run
	return filepath.Base(os.Args[0])
}
