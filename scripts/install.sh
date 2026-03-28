#!/bin/bash
set -e

# Hailow - AI Agent Configuration Manager Installation Script
# This script installs the hailow CLI tool

REPO="Harvey-N-Lab/hailow"
BINARY_NAME="hailow"
INSTALL_DIR="${INSTALL_DIR:-$HOME/.local/bin}"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo_info() {
    echo -e "${GREEN}==>${NC} $1"
}

echo_error() {
    echo -e "${RED}Error:${NC} $1" >&2
}

echo_warn() {
    echo -e "${YELLOW}Warning:${NC} $1"
}

# Detect OS and architecture
detect_platform() {
    local os=$(uname -s | tr '[:upper:]' '[:lower:]')
    local arch=$(uname -m)

    case "$os" in
        linux*)
            OS="linux"
            ;;
        darwin*)
            OS="darwin"
            ;;
        msys*|mingw*|cygwin*)
            OS="windows"
            ;;
        *)
            echo_error "Unsupported operating system: $os"
            exit 1
            ;;
    esac

    case "$arch" in
        x86_64|amd64)
            ARCH="amd64"
            ;;
        aarch64|arm64)
            ARCH="arm64"
            ;;
        *)
            echo_error "Unsupported architecture: $arch"
            exit 1
            ;;
    esac

    echo_info "Detected platform: $OS/$ARCH"
}

# Get latest release version
get_latest_version() {
    echo_info "Fetching latest release..."
    LATEST_VERSION=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    
    if [ -z "$LATEST_VERSION" ]; then
        echo_error "Failed to fetch latest version"
        echo_error "No releases found. Please build from source:"
        echo_error "  git clone https://github.com/$REPO"
        echo_error "  cd hailow && make build"
        exit 1
    fi
    
    echo_info "Latest version: $LATEST_VERSION"
}

# Download binary
download_binary() {
    # GoReleaser format: hailow_VERSION_OS_ARCH.tar.gz
    # Strip 'v' from version for filename
    local version_num="${LATEST_VERSION#v}"
    local filename="${BINARY_NAME}_${version_num}_${OS}_${ARCH}"
    
    if [ "$OS" = "windows" ]; then
        filename="${filename}.zip"
    else
        filename="${filename}.tar.gz"
    fi
    
    local url="https://github.com/$REPO/releases/download/$LATEST_VERSION/$filename"
    TMPDIR=$(mktemp -d)
    
    echo_info "Downloading from $url..."
    
    if ! curl -f -L -o "$TMPDIR/$filename" "$url" 2>/dev/null; then
        echo_error "Failed to download binary"
        echo_error "URL: $url"
        echo_error "Release may not exist yet. Build from source:"
        echo_error "  git clone https://github.com/$REPO && cd hailow && make build"
        rm -rf "$TMPDIR"
        exit 1
    fi
    
    echo_info "Extracting..."
    
    if [ "$OS" = "windows" ]; then
        unzip -q "$TMPDIR/$filename" -d "$TMPDIR"
    else
        tar -xzf "$TMPDIR/$filename" -C "$TMPDIR"
    fi
}

# Install binary
install_binary() {
    # Create install directory if it doesn't exist
    mkdir -p "$INSTALL_DIR"
    
    # Copy binary
    echo_info "Installing to $INSTALL_DIR/$BINARY_NAME..."
    
    if [ "$OS" = "windows" ]; then
        cp "$TMPDIR/${BINARY_NAME}.exe" "$INSTALL_DIR/"
        chmod +x "$INSTALL_DIR/${BINARY_NAME}.exe"
    else
        cp "$TMPDIR/$BINARY_NAME" "$INSTALL_DIR/"
        chmod +x "$INSTALL_DIR/$BINARY_NAME"
    fi
    
    # Cleanup
    rm -rf "$TMPDIR"
}

# Verify installation
verify_installation() {
    if command -v $BINARY_NAME &> /dev/null; then
        local version=$($BINARY_NAME version 2>/dev/null | head -n1 || echo "unknown")
        echo_info "Installation successful!"
        echo_info "Version: $version"
        return 0
    else
        echo_warn "Binary installed but not in PATH"
        echo_warn "Add $INSTALL_DIR to your PATH:"
        echo "  export PATH=\"$INSTALL_DIR:\$PATH\""
        echo ""
        echo "Add this to your shell profile (~/.bashrc, ~/.zshrc, etc.) to make it permanent"
        return 1
    fi
}

# Main installation
main() {
    echo "Hailow - AI Agent Configuration Manager - Installation Script"
    echo "=============================================================="
    echo ""
    
    detect_platform
    get_latest_version
    
    download_binary
    install_binary
    
    echo ""
    verify_installation
    
    echo ""
    echo "Next steps:"
    echo "  1. Run: $BINARY_NAME list domains"
    echo "  2. Install a domain: $BINARY_NAME install devops-engineer ./my-project"
    echo "  3. Read the docs: https://github.com/$REPO"
    echo ""
}

main "$@"
