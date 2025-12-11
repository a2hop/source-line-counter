package counters

type Counter interface {
	ShouldCount(path string) bool
	CountLines(filePath string, includeWhitespace bool) (int, error)
	Name() string
}
