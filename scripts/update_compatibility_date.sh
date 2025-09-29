#!/bin/sh

# This script extracts the ESICompatibilityDate from esi-openapi-spec.json
# and updates the constant in APIClient.go.

set -e

SPEC_FILE="esi-openapi-spec.json"
TARGET_FILE="APIClient.go"

echo "Updating ${TARGET_FILE} with ESICompatibilityDate from ${SPEC_FILE}..."

# Extract the compatibility date
COMPATIBILITY_DATE=$(jq -r '.info.version' ${SPEC_FILE})

if [ -z "${COMPATIBILITY_DATE}" ]; then
    echo "Error: Could not extract ESICompatibilityDate from ${SPEC_FILE}"
    exit 1
fi

# Update the constant in the Go file
sed -i "s/ESICompatibilityDate = \".*\"/ESICompatibilityDate = \"${COMPATIBILITY_DATE}\"/" ${TARGET_FILE}

echo "Successfully updated ${TARGET_FILE} with ESICompatibilityDate: ${COMPATIBILITY_DATE}"
