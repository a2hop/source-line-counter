package cli

import (
	"fmt"
	"sort"

	skips "github.com/a2hop/source-line-counter/skiplists"
)

func PrintSkipLists() {
	fmt.Println("Available Skip Lists:")
	fmt.Println("────────────────────────────────")

	names := skips.SkipLists.GetSkipListNames()
	sort.Strings(names)

	for _, name := range names {
		list := skips.SkipLists.GetSkipList(name)

		// Count items
		dirs := list.GetDirs()
		files := list.GetFiles()
		patterns := list.GetPatterns()

		fmt.Printf("  %s (%dd, %df, %dp)\n",
			name, len(dirs), len(files), len(patterns))
	}

	fmt.Println("\nUse --show-skiplist=<name> for details")
}

func PrintSkipListContent(name string) {
	list := skips.SkipLists.GetSkipList(name)
	if list == nil {
		fmt.Printf("Skip list '%s' not found\n", name)
		fmt.Println("Use --list-skiplists to see available skip lists")
		return
	}

	fmt.Printf("Skip List: %s\n", name)
	fmt.Println("────────────────────────────────")

	// Print directories
	dirs := list.GetDirs()
	if len(dirs) > 0 {
		sort.Strings(dirs)
		fmt.Println("Directories:")
		for _, dir := range dirs {
			fmt.Printf("  %s/\n", dir)
		}
	}

	// Print files
	files := list.GetFiles()
	if len(files) > 0 {
		sort.Strings(files)
		fmt.Println("Files:")
		for _, file := range files {
			fmt.Printf("  %s\n", file)
		}
	}

	// Print patterns
	patterns := list.GetPatterns()
	if len(patterns) > 0 {
		sort.Strings(patterns)
		fmt.Println("Patterns:")
		for _, pattern := range patterns {
			fmt.Printf("  %s\n", pattern)
		}
	}
}

func PrintAllSkipListsContent() {
	names := skips.SkipLists.GetSkipListNames()
	sort.Strings(names)

	for i, name := range names {
		if i > 0 {
			fmt.Println()
		}
		PrintSkipListContent(name)
	}
}
