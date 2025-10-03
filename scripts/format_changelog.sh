#!/bin/bash
set -e

# This script formats ESI changelog entries into Markdown for GitHub releases
# It takes previous and current ESI versions as arguments and formats all changes between them

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Source the common changelog library
source "${SCRIPT_DIR}/changelog_common.sh"

PREV_ESI_VERSION="${1:-}"
CURR_ESI_VERSION="${2:-}"

if [ -z "$CURR_ESI_VERSION" ]; then
    echo "Error: Current ESI version required" >&2
    echo "Usage: $0 [previous_version] current_version" >&2
    exit 1
fi

# Fetch changelog using shared function
CHANGELOG=$(fetch_changelog)

if [ -z "$CHANGELOG" ]; then
    echo "Error: Failed to fetch changelog from ESI" >&2
    exit 1
fi

# Get all dates using shared function
ALL_CHANGELOG_DATES=$(get_all_changelog_dates "$CHANGELOG")

# Get relevant dates using shared function
RELEVANT_DATES=$(get_relevant_dates "$PREV_ESI_VERSION" "$CURR_ESI_VERSION" "$ALL_CHANGELOG_DATES")

if [ -z "$RELEVANT_DATES" ]; then
    echo "No changelog entries found for this version."
    exit 0
fi

# Function to format a single change entry
format_change() {
    local method="$1"
    local path="$2"
    local description="$3"
    
    echo "- **${method} ${path}** - ${description}"
}

# Format changelog for all relevant dates
for date in $RELEVANT_DATES; do
    CHANGES=$(echo "$CHANGELOG" | jq --arg d "$date" '.changelog[$d]')
    
    if [ -z "$CHANGES" ] || [ "$CHANGES" = "null" ] || [ "$CHANGES" = "[]" ]; then
        continue
    fi
    
    echo "### ESI Changes (${date})"
    echo ""
    
    # Get counts for each change type
    BREAKING_COUNT=$(echo "$CHANGES" | jq '[.[] | select(.type == "breaking")] | length')
    NEW_COUNT=$(echo "$CHANGES" | jq '[.[] | select(.type == "new")] | length')
    CHANGED_COUNT=$(echo "$CHANGES" | jq '[.[] | select(.type == "changed")] | length')
    REMOVED_COUNT=$(echo "$CHANGES" | jq '[.[] | select(.type == "removed")] | length')
    
    # Format breaking changes
    if [ "$BREAKING_COUNT" -gt 0 ]; then
        echo "#### Breaking Changes"
        echo ""
        echo "$CHANGES" | jq -r '.[] | select(.type == "breaking") | "- **\(.method) \(.path)** - \(.description)"'
        echo ""
    fi
    
    # Format removed endpoints
    if [ "$REMOVED_COUNT" -gt 0 ]; then
        echo "#### Removed Endpoints"
        echo ""
        echo "$CHANGES" | jq -r '.[] | select(.type == "removed") | "- **\(.method) \(.path)**" + (if .description != "" then " - \(.description)" else "" end)'
        echo ""
    fi
    
    # Format new endpoints
    if [ "$NEW_COUNT" -gt 0 ]; then
        echo "#### New Endpoints"
        echo ""
        echo "$CHANGES" | jq -r '.[] | select(.type == "new") | "- **\(.method) \(.path)** - \(.description)"'
        echo ""
    fi
    
    # Format changed endpoints
    if [ "$CHANGED_COUNT" -gt 0 ]; then
        echo "#### Changed Endpoints"
        echo ""
        echo "$CHANGES" | jq -r '.[] | select(.type == "changed") | "- **\(.method) \(.path)** - \(.description)"'
        echo ""
    fi
done
