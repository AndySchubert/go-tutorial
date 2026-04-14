#!/bin/bash
set -e

# 1. Get the latest Go version string (e.g., "go1.26.1")
LATEST_GO=$(curl -sSL "https://go.dev/VERSION?m=text" | head -n 1)
ARCH="amd64"
FILE="${LATEST_GO}.linux-${ARCH}.tar.gz"
URL="https://go.dev/dl/${FILE}"

echo "Latest Go version found: $LATEST_GO"

# 2. Download the archive
echo "Downloading $URL..."
wget -q --show-progress "$URL" -O "$FILE"

# 3. Remove old version (requires sudo)
echo "Removing old Go installation at /usr/local/go..."
sudo rm -rf /usr/local/go

# 4. Extract new version
echo "Extracting $FILE to /usr/local..."
sudo tar -C /usr/local -xzf "$FILE"

# 5. Clean up downloaded file
rm "$FILE"

# 6. Final verification
echo "Installation complete!"
/usr/local/go/bin/go version
