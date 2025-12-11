package skips

import (
	"path/filepath"
	"strings"
)

type skipLists struct {
	Lists map[string]ISkipList
}

func NewSkipLists() *skipLists {
	return &skipLists{
		Lists: make(map[string]ISkipList),
	}
}

func (s *skipLists) AddSkipList(name string, list ISkipList) {
	s.Lists[name] = list
}

func (s *skipLists) GetSkipList(name string) ISkipList {
	return s.Lists[name]
}

// GetSkipListNames returns all available skip list names
func (s *skipLists) GetSkipListNames() []string {
	names := make([]string, 0, len(s.Lists))
	for name := range s.Lists {
		names = append(names, name)
	}
	return names
}

// ShouldSkip checks if a path should be skipped across all skip lists
func (s *skipLists) ShouldSkip(path string) bool {
	for _, list := range s.Lists {
		if list.ShouldSkip(path) {
			return true
		}
	}
	return false
}

type ISkipList interface {
	AddDir(dir string)
	AddFile(file string)
	AddPattern(pattern string)
	ShouldSkip(path string) bool
	GetName() string
	GetDirs() []string
	GetFiles() []string
	GetPatterns() []string
}

type SkipList struct {
	Name     string
	Dirs     map[string]struct{}
	Files    map[string]struct{}
	Patterns []string
}

func (s *SkipList) AddDir(dir string) {
	s.Dirs[dir] = struct{}{}
}

func (s *SkipList) AddFile(file string) {
	s.Files[file] = struct{}{}
}

func (s *SkipList) AddPattern(pattern string) {
	s.Patterns = append(s.Patterns, pattern)
}

func (s *SkipList) GetName() string {
	return s.Name
}

func (s *SkipList) GetDirs() []string {
	dirs := make([]string, 0, len(s.Dirs))
	for dir := range s.Dirs {
		dirs = append(dirs, dir)
	}
	return dirs
}

func (s *SkipList) GetFiles() []string {
	files := make([]string, 0, len(s.Files))
	for file := range s.Files {
		files = append(files, file)
	}
	return files
}

func (s *SkipList) GetPatterns() []string {
	return s.Patterns
}

func (s *SkipList) ShouldSkip(path string) bool {
	// Check if any directory component should be skipped
	pathParts := strings.Split(filepath.ToSlash(path), "/")
	for _, part := range pathParts {
		if _, exists := s.Dirs[part]; exists {
			return true
		}
	}

	// Check if the filename should be skipped
	filename := filepath.Base(path)
	if _, exists := s.Files[filename]; exists {
		return true
	}

	// Check if the path matches any patterns
	for _, pattern := range s.Patterns {
		matched, err := filepath.Match(pattern, filename)
		if err == nil && matched {
			return true
		}
	}

	return false
}

var SkipLists = NewSkipLists()

func init() {
	SkipLists.AddSkipList("general", NewGeneralSkipList())
}
