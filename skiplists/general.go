package skips

type GeneralSkipList struct {
	SkipList
}

func NewGeneralSkipList() *GeneralSkipList {
	return &GeneralSkipList{
		SkipList: SkipList{
			Name: "general",
			Dirs: map[string]struct{}{
				".git":         {},
				".svn":         {},
				".hg":          {},
				"node_modules": {},
				"vendor":       {},
				"bin":          {},
				"obj":          {},
				".idea":        {},
				".vscode":      {},
			},
			Files: map[string]struct{}{
				".DS_Store":      {},
				"Thumbs.db":      {},
				".gitignore":     {},
				".gitattributes": {},
			},
			Patterns: []string{
				"*.min.js",
				"*.min.css",
			},
		},
	}
}
