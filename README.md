# src-counter - Source Line Counter

A fast and efficient command-line tool to count lines of code in Go, JavaScript/TypeScript, and C/C++ projects.

## Features

- 📊 **Multi-language support**: Go, JavaScript/TypeScript, C/C++
- 🎯 **Smart counting**: Excludes comments and blank lines by default
- 📁 **Skip lists**: Built-in filtering for common directories (`.git`, `node_modules`, etc.)
- 📈 **Per-language statistics**: See breakdown by programming language
- 🔍 **Verbose mode**: Optional per-file line counts
- ⚙️ **Configurable**: Include/exclude whitespace, customize skip lists

## Installation

### Download Pre-built Binary

Download the latest release for your platform from the [releases page](https://github.com/a2hop/source-line-counter/releases).

**Linux/macOS:**
```bash
# Download and extract (replace <platform> with your platform)
tar xzf src-counter-v1.0.0-<platform>.tar.gz

# Move to system path
sudo mv src-counter /usr/local/bin/

# Verify installation
src-counter --version
```

**Windows:**
```powershell
# Extract the zip file
# Move src-counter.exe to a directory in your PATH
# Or run from the extracted location
```

### Build from Source

```bash
# Clone the repository
git clone https://github.com/a2hop/source-line-counter.git
cd source-line-counter

# Build
go build -o src-counter .

# Install (optional)
sudo mv src-counter /usr/local/bin/
```

## Usage

### Basic Usage

```bash
# Count lines in current directory
src-counter

# Count lines in specific directory
src-counter /path/to/project

# Show per-file breakdown
src-counter -v

# Include whitespace and comments
src-counter -w
```

### Examples

```bash
# Analyze a project with verbose output
src-counter -v ~/projects/myapp

# Count all lines (including comments and blanks)
src-counter -w .

# Disable skip lists
src-counter -skip none

# Use specific directory
src-counter -path ./src
```

### Options

```
  -path string         Directory to analyze (default: current directory)
  -skip string         Skip lists to use: comma-separated (default: "general")
  -v                   Show per-file line counts
  -w                   Include whitespace and comments
  -h, -help            Show help message
  --help-verbose       Show detailed help with examples
  --version            Show version information
  --list-skiplists     List all available skip lists
  --show-skiplist=name Show contents of a specific skip list
```

### Skip Lists

By default, src-counter skips common directories and files:

```bash
# See what's being skipped
src-counter --list-skiplists

# View details of a skip list
src-counter --show-skiplist=general

# Disable skipping
src-counter -skip none
```

The default "general" skip list excludes:
- Version control: `.git`, `.svn`, `.hg`
- Dependencies: `node_modules`, `vendor`
- Build artifacts: `bin`, `obj`
- IDE files: `.idea`, `.vscode`
- Minified files: `*.min.js`, `*.min.css`

## Output

### Standard Output

```
Language                Lines  Files
────────────────────────────────────
Go                        581     10
JavaScript/TypeScript     234      5
────────────────────────────────────
Total                     815     15 (code)
```

### With Verbose Flag (-v)

```
main.go: 72 lines
counters/counter.go: 6 lines
counters/go_counter.go: 34 lines
...

Language                Lines  Files
────────────────────────────────────
Go                        581     10
────────────────────────────────────
Total                     581     10 (code)
```

## Supported Languages

| Language               | Extensions                      |
|------------------------|---------------------------------|
| Go                     | `.go`                           |
| JavaScript/TypeScript  | `.js`, `.jsx`, `.ts`, `.tsx`    |
| C/C++                  | `.c`, `.cpp`, `.cc`, `.cxx`, `.h`, `.hpp`, `.hxx` |

## Development

### Building

```bash
# Build for current platform
./scripts/build.sh

# Build for all platforms
./scripts/build_all.sh
```

### Project Structure

```
.
├── main.go              # Main entry point
├── about/               # Version and program info
├── cli/                 # CLI helpers (help, skiplist printing)
├── counters/            # Language-specific line counters
├── skiplists/           # Skip list implementations
├── scripts/             # Build scripts
└── .github/workflows/   # GitHub Actions CI/CD
```

## Release Process

1. Update version in `version` file
2. Commit changes
3. Push to main branch
4. GitHub Actions will automatically:
   - Build binaries for all platforms
   - Create release with artifacts
   - Generate SHA256 checksums

## License

MIT License - see repository for details

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## Links

- Repository: https://github.com/a2hop/source-line-counter
- Issues: https://github.com/a2hop/source-line-counter/issues
