package cli

import (
	"fmt"

	"github.com/a2hop/source-line-counter/about"
)

func PrintHelp(verbose bool) {
	if verbose {
		printVerboseHelp()
	} else {
		printShortHelp()
	}
}

func PrintVersion() {
	fmt.Printf("%s v%s\n", about.ProgramName, about.Version)
	fmt.Printf("%s\n", about.Description)
	fmt.Printf("License: %s\n", about.License)
	fmt.Printf("Repository: %s\n", about.GitUrl)
}

func printShortHelp() {
	fmt.Printf("%s v%s - %s\n\n", about.GetBinName(), about.Version, about.ProgramName)
	helpText := `Usage: src-counter [options] [directory]

Options:
  -path string         Directory to analyze (default: current directory)
  -skip string         Skip lists to use: comma-separated (default: "general")
  -v                   Show per-file line counts
  -w                   Include whitespace and comments
  -h, -help            Show this help message
  --help-verbose       Show detailed help with examples
  --version            Show version information
  --list-skiplists     List all available skip lists
  --show-skiplist=name Show contents of a specific skip list

Examples:
  src-counter                      # Count lines in current directory
  src-counter -w                   # Include whitespace and comments
  src-counter -v /path/to/dir      # Verbose output for specific directory
  src-counter --list-skiplists     # Show available skip lists
  src-counter --show-skiplist=general  # Show contents of 'general' skip list
  src-counter --help-verbose       # Show detailed help
`
	fmt.Println(helpText)
}

func printVerboseHelp() {
	fmt.Printf("%s v%s - %s\n\n", about.GetBinName(), about.Version, about.ProgramName)
	helpText := `Usage: src-counter [options] [directory]

Description:
  Counts lines of code in Go, JavaScript/TypeScript, and C/C++ files.
  By default, counts only code lines (excluding comments and blank lines).

Options:
  -path string
        The root directory to analyze (default: current directory)
        Can also be specified as the first positional argument.

  -skip string
        Comma-separated list of skip lists to use (default: "general")
        Use "none" to disable all skip lists.
        Available skip lists:
          - general: Skips common directories like .git, node_modules, vendor, etc.

  -v    Show verbose output with per-file line counts

  -w    Include whitespace and comments in line count
        By default, only code lines are counted (excluding blank lines and comments)

  -h, -help
        Show this help message (short version)

  --help-verbose
        Show detailed help with examples

  --version
        Show version information

  --list-skiplists
        List all available skip lists with their item counts

  --show-skiplist=name
        Show the detailed contents of a specific skip list
        (directories, files, and patterns that will be skipped)

Supported File Types:
  Go:                    .go
  JavaScript/TypeScript: .js, .jsx, .ts, .tsx
  C/C++:                 .c, .cpp, .cc, .cxx, .h, .hpp, .hxx

Examples:
  # Count lines in current directory (excluding comments/blanks)
  src-counter

  # Count all lines including whitespace and comments
  src-counter -w

  # Count lines in a specific directory with verbose output
  src-counter -v /path/to/project

  # Use multiple skip lists
  src-counter -skip general,tests

  # Disable all skip lists
  src-counter -skip none

  # Count with all options
  src-counter -path ./myproject -skip general -v -w

  # Using positional argument for directory
  src-counter ./src

  # List available skip lists
  src-counter --list-skiplists

  # Show what the 'general' skip list contains
  src-counter --show-skiplist=general

Repository: %s
License: %s
`
	fmt.Printf(helpText, about.GitUrl, about.License)
}
