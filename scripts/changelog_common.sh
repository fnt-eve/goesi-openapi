#!/bin/bash

# Common functions for working with ESI changelog

# Fetch the ESI changelog with the latest compatibility date to get new format
fetch_changelog() {
    curl -s "https://esi.evetech.net/meta/changelog?compatibility_date=2025-09-30"
}

# Get all changelog dates sorted
get_all_changelog_dates() {
    local changelog="$1"
    echo "$changelog" | jq -r '.changelog | keys[]' | sort -V
}

# Get dates between PREV (exclusive) and CURR (inclusive)
# Args: prev_date curr_date all_dates
get_relevant_dates() {
    local prev="$1"
    local curr="$2"
    local all_dates="$3"
    
    if [ -z "$prev" ]; then
        # No previous version, include only current
        echo "$curr"
        return
    fi
    
    # Filter dates: must be > prev AND <= curr
    for date in $all_dates; do
        # Skip if date <= prev
        if [ "$date" = "$prev" ] || [ "$date" \< "$prev" ]; then
            continue
        fi
        
        # Skip if date > curr  
        if [ "$date" \> "$curr" ]; then
            continue
        fi
        
        # Date is in range
        echo "$date"
    done
}
