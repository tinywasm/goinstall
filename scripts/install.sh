#!/bin/bash
set -e

# Default version if not provided
GO_VERSION=${1:-"1.25.2"}

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64)
        GOARCH="amd64"
        ;;
    aarch64|arm64)
        GOARCH="arm64"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

FILENAME="go$GO_VERSION.$OS-$GOARCH.tar.gz"
URL="https://go.dev/dl/$FILENAME"

echo "Downloading Go $GO_VERSION for $OS-$GOARCH..."
curl -L -o /tmp/$FILENAME $URL

echo "Removing existing Go installation..."
sudo rm -rf /usr/local/go

echo "Extracting Go..."
sudo tar -C /usr/local -xzf /tmp/$FILENAME

echo "Creating symlink..."
sudo ln -sf /usr/local/go/bin/go /usr/local/bin/go

echo "Verifying installation..."
/usr/local/go/bin/go version

rm /tmp/$FILENAME
echo "Go $GO_VERSION installed successfully."
