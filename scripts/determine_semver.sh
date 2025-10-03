#!/bin/bash
set -e

# This script determines the next semantic version based on ESI changelog changes
# It reads the previous ESI version from ESI_VERSION file and compares with current spec

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Source the common changelog library
source "${SCRIPT_DIR}/changelog_common.sh"

# Read previous ESI version (or empty if file doesn't exist)
PREV_ESI_VERSION=$(cat ESI_VERSION 2>/dev/null || echo "")

# Get current ESI version from spec file
if [ ! -f "esi-openapi-spec.json" ]; then
    echo "Error: esi-openapi-spec.json not found" >&2
    exit 1
fi

CURR_ESI_VERSION=$(jq -r '.info.version' esi-openapi-spec.json)

if [ -z "$CURR_ESI_VERSION" ] || [ "$CURR_ESI_VERSION" = "null" ]; then
    echo "Error: Could not extract version from spec" >&2
    exit 1
fi

# Get latest semver tag
LATEST_TAG=$(git tag -l "v*.*.*" 2>/dev/null | sort -V | tail -n 1)

if [ -z "$LATEST_TAG" ]; then
    # No semver tags yet, start at v0.0.0
    MAJOR=0
    MINOR=0
    PATCH=0
    echo "No existing semver tags found, starting at v0.0.0" >&2
else
    # Parse existing tag
    MAJOR=$(echo "$LATEST_TAG" | sed 's/v\([0-9]*\)\..*/\1/')
    MINOR=$(echo "$LATEST_TAG" | sed 's/v[0-9]*\.\([0-9]*\)\..*/\1/')
    PATCH=$(echo "$LATEST_TAG" | sed 's/v[0-9]*\.[0-9]*\.\([0-9]*\).*/\1/')
    echo "Current version: ${LATEST_TAG} (MAJOR=${MAJOR}, MINOR=${MINOR}, PATCH=${PATCH})" >&2
fi

# Check if ESI version changed
if [ "$PREV_ESI_VERSION" = "$CURR_ESI_VERSION" ]; then
    echo "No ESI version change detected (still at ${CURR_ESI_VERSION})" >&2
    echo "Manual release - incrementing PATCH version" >&2
    PATCH=$((PATCH + 1))
    NEW_VERSION="v${MAJOR}.${MINOR}.${PATCH}"
    echo "New version: ${NEW_VERSION}" >&2
    echo "$NEW_VERSION"
    exit 0
fi

echo "ESI version change detected: ${PREV_ESI_VERSION} -> ${CURR_ESI_VERSION}" >&2
echo "ESI version changed - incrementing MINOR version" >&2

# Increment MINOR version, reset PATCH
MINOR=$((MINOR + 1))
PATCH=0

NEW_VERSION="v${MAJOR}.${MINOR}.${PATCH}"
echo "New version: ${NEW_VERSION}" >&2
echo "$NEW_VERSION"
