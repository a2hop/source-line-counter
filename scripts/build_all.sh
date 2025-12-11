#!/bin/bash
# Build for all platforms (similar to CI)

set -e

VERSION=$(cat version)
echo "Building src-counter v${VERSION} for all platforms..."

# Create release directory
mkdir -p release

# Build matrix
declare -a builds=(
  "linux:amd64::linux-amd64"
  "linux:arm64::linux-arm64"
  "linux:arm:7:linux-armv7"
  "darwin:amd64::darwin-amd64"
  "darwin:arm64::darwin-arm64"
  "windows:amd64::windows-amd64"
  "windows:arm64::windows-arm64"
)

for build in "${builds[@]}"; do
  IFS=':' read -r GOOS GOARCH GOARM PLATFORM <<< "$build"
  
  echo ""
  echo "Building for ${PLATFORM}..."
  
  BINARY_NAME="src-counter"
  if [ "$GOOS" = "windows" ]; then
    BINARY_NAME="src-counter.exe"
  fi
  
  # Build
  if [ -n "$GOARM" ]; then
    CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH GOARM=$GOARM \
      go build -ldflags="-s -w -X github.com/a2hop/source-line-counter/about.Version=${VERSION}" \
      -o "release/${BINARY_NAME}" .
  else
    CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH \
      go build -ldflags="-s -w -X github.com/a2hop/source-line-counter/about.Version=${VERSION}" \
      -o "release/${BINARY_NAME}" .
  fi
  
  # Create archive
  cd release
  ARCHIVE_NAME="src-counter-v${VERSION}-${PLATFORM}"
  if [ "$GOOS" = "windows" ]; then
    zip "${ARCHIVE_NAME}.zip" "${BINARY_NAME}"
    rm "${BINARY_NAME}"
    echo "  Created: ${ARCHIVE_NAME}.zip"
  else
    tar czf "${ARCHIVE_NAME}.tar.gz" "${BINARY_NAME}"
    rm "${BINARY_NAME}"
    echo "  Created: ${ARCHIVE_NAME}.tar.gz"
  fi
  cd ..
done

echo ""
echo "✅ All builds complete!"
echo ""
echo "Artifacts:"
ls -lh release/

echo ""
echo "Generating checksums..."
cd release
sha256sum src-counter-v* > SHA256SUMS
echo ""
cat SHA256SUMS
cd ..

echo ""
echo "Release artifacts are in: ./release/"
