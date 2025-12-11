package main

import (
	"flag"
	"fmt"
	"goloc/cli"
	"goloc/counters"
	skips "goloc/skiplists"
	"os"
	"path/filepath"
	"strings"
)

type LanguageStats struct {
	lines int
	files int
}

func main() {
	// Parse command line flags
	help := flag.Bool("h", false, "Show help message")
	helpLong := flag.Bool("help", false, "Show help message")
	helpVerbose := flag.Bool("help-verbose", false, "Show detailed help message")
	version := flag.Bool("version", false, "Show version information")
	listSkipLists := flag.Bool("list-skiplists", false, "List all available skip lists")
	showSkipList := flag.String("show-skiplist", "", "Show contents of a specific skip list")
	verbose := flag.Bool("v", false, "Show per-file line counts")
	includeWhitespace := flag.Bool("w", false, "Include whitespace and comments in line count")
	skipListNames := flag.String("skip", "general", "Comma-separated list of skip lists to use (default: general)")
	pathFlag := flag.String("path", ".", "The root directory to analyze")
	flag.Parse()

	// Show version if requested
	if *version {
		cli.PrintVersion()
		os.Exit(0)
	}

	// Show skip lists if requested
	if *listSkipLists {
		cli.PrintSkipLists()
		os.Exit(0)
	}

	// Show specific skip list content if requested
	if *showSkipList != "" {
		cli.PrintSkipListContent(*showSkipList)
		os.Exit(0)
	}

	// Show help if requested
	if *helpVerbose {
		cli.PrintHelp(true)
		os.Exit(0)
	}
	if *help || *helpLong {
		cli.PrintHelp(false)
		os.Exit(0)
	}

	// Get the directory to scan from -path flag or remaining args, or use current directory
	dir := *pathFlag
	if flag.NArg() > 0 {
		dir = flag.Arg(0)
	}

	// Configure skip lists
	activeSkipLists := skips.NewSkipLists()
	if *skipListNames != "" && *skipListNames != "none" {
		listNames := strings.Split(*skipListNames, ",")
		for _, name := range listNames {
			name = strings.TrimSpace(name)
			if list := skips.SkipLists.GetSkipList(name); list != nil {
				activeSkipLists.AddSkipList(name, list)
			} else {
				fmt.Fprintf(os.Stderr, "Warning: skip list '%s' not found, ignoring\n", name)
			}
		}
	}

	// Initialize all counters
	allCounters := []counters.Counter{
		&counters.GoCounter{},
		&counters.JsCounter{},
		&counters.CppCounter{},
	}

	// Track stats per language
	stats := make(map[string]*LanguageStats)
	for _, counter := range allCounters {
		stats[counter.Name()] = &LanguageStats{}
	}

	totalLines := 0
	fileCount := 0

	// Walk through the directory recursively
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip if this path is in the active skip lists
		if activeSkipLists.ShouldSkip(path) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Try each counter to see if it handles this file
		for _, counter := range allCounters {
			if counter.ShouldCount(path) {
				// Count lines in this file
				lines, err := counter.CountLines(path, *includeWhitespace)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", path, err)
					return nil
				}

				totalLines += lines
				fileCount++

				// Update language stats
				langStats := stats[counter.Name()]
				langStats.lines += lines
				langStats.files++

				if *verbose {
					fmt.Printf("%s: %d lines\n", path, lines)
				}
				break
			}
		}

		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error walking directory: %v\n", err)
		os.Exit(1)
	}

	// Print per-language totals
	countType := "code"
	if *includeWhitespace {
		countType = "total"
	}

	fmt.Printf("\n%-20s %8s %6s\n", "Language", "Lines", "Files")
	fmt.Println("────────────────────────────────────")
	for _, counter := range allCounters {
		langStats := stats[counter.Name()]
		if langStats.files > 0 {
			fmt.Printf("%-20s %8d %6d\n", counter.Name(), langStats.lines, langStats.files)
		}
	}
	fmt.Println("────────────────────────────────────")
	fmt.Printf("%-20s %8d %6d (%s)\n", "Total", totalLines, fileCount, countType)
}
