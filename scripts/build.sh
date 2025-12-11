#!/bin/bash
# Local build script to test the release build process

set -e

VERSION=$(cat version)
echo "Building src-counter v${VERSION}..."

# Create build directory
mkdir -p build

# Build for current platform
echo "Building for current platform..."
go build -ldflags="-s -w -X github.com/a2hop/source-line-counter/about.Version=${VERSION}" -o build/src-counter .

echo ""
echo "Build successful!"
echo "Binary: build/src-counter"
echo ""
echo "Testing binary:"
build/src-counter --version

echo ""
echo "To test all platforms, run:"
echo "  ./scripts/build_all.sh"
