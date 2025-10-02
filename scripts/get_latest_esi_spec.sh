#!/bin/sh

# This script fetches the latest ESI OpenAPI specification with correct versioning
# It queries the compatibility dates endpoint and downloads the spec for the latest date

set -e

DATES_ENDPOINT="https://esi.evetech.net/meta/compatibility-dates"
SPEC_ENDPOINT="https://esi.evetech.net/meta/openapi-3.0.json"
OUTPUT_FILE="esi-openapi-spec.json"

echo "Fetching latest ESI compatibility dates from ${DATES_ENDPOINT}..."

# Fetch compatibility dates
DATES_RESPONSE=$(curl -s -H "Accept: application/json" "${DATES_ENDPOINT}")

# Extract the latest date (first element in the array)
LATEST_DATE=$(echo "${DATES_RESPONSE}" | jq -r '.compatibility_dates[0]')

if [ -z "${LATEST_DATE}" ] || [ "${LATEST_DATE}" = "null" ]; then
    echo "Error: Could not extract latest compatibility date from response"
    echo "Response: ${DATES_RESPONSE}"
    exit 1
fi

echo "Latest ESI compatibility date: ${LATEST_DATE}"

# Download the spec with the compatibility date parameter
echo "Downloading ESI OpenAPI specification for date ${LATEST_DATE}..."
curl -s -L "${SPEC_ENDPOINT}?compatibility_date=${LATEST_DATE}" -o "${OUTPUT_FILE}"

# Validate the downloaded JSON
if ! jq empty "${OUTPUT_FILE}" >/dev/null 2>&1; then
    echo "Error: Downloaded specification is not valid JSON"
    exit 1
fi

echo "Successfully downloaded ESI specification to ${OUTPUT_FILE}"
echo "Specification version: ${LATEST_DATE}"

# Output the date for use by other scripts (to stdout)
echo "${LATEST_DATE}"
