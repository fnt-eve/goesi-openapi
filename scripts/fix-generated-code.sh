#!/bin/bash

# Script to fix the generated ESI client code to automatically handle X-Compatibility-Date
# This should be run after each code generation

# Configuration
CLIENT_DIR="esi"
SPEC_FILE="esi-openapi-spec.json"

echo "Fixing generated ESI client code..."

# Extract current ESI compatibility date from the spec file
COMPAT_DATE=$(jq -r '.info.version' ${SPEC_FILE})
if [ -z "$COMPAT_DATE" ]; then
    echo "Warning: Could not extract compatibility date from ${SPEC_FILE}, using default"
    COMPAT_DATE="2020-01-01"
fi
echo "Using compatibility date: $COMPAT_DATE"

# Add CompatibilityDate field to Configuration struct in configuration.go
CONFIG_FILE="${CLIENT_DIR}/configuration.go"
if [ -f "$CONFIG_FILE" ]; then
    echo "Adding CompatibilityDate field to Configuration struct..."
    
    # Check if field already exists
    if ! grep -q "CompatibilityDate" "$CONFIG_FILE"; then
        # Add CompatibilityDate field after HTTPClient field
        sed -i.bak '/HTTPClient[[:space:]]*\*http\.Client/a\
\tCompatibilityDate  string            `json:"compatibilityDate,omitempty"`' "$CONFIG_FILE"
        
        # Set default value in NewConfiguration function
        sed -i.bak '/Debug:[[:space:]]*false,/a\
\t\tCompatibilityDate: "'"$COMPAT_DATE"'",' "$CONFIG_FILE"
        
        rm -f "${CONFIG_FILE}.bak"
        echo "CompatibilityDate field added to Configuration struct"
    else
        echo "CompatibilityDate field already exists in Configuration struct"
    fi
fi

# Find all API files in the client directory
find ${CLIENT_DIR}/ -name "api_*.go" -type f | while read -r file; do
    echo "Processing $file..."
    
    # Replace the nil check to use the configuration's CompatibilityDate instead of hardcoded value
    sed -i.bak 's/if r\.xCompatibilityDate == nil {/if r.xCompatibilityDate == nil {\
\t\tr.xCompatibilityDate = \&a.client.cfg.CompatibilityDate\
\t}\
\t\/\/ Removed original error check - now using configuration default value\
\tif false {/g' "$file"
    
    # Clean up backup files
    rm -f "${file}.bak"
done

# Remove test directory if it was generated
if [ -d "${CLIENT_DIR}/test" ]; then
    echo "Removing generated test directory..."
    rm -rf ${CLIENT_DIR}/test
fi

# Remove any generated test files
find ${CLIENT_DIR}/ -name "*_test.go" -type f -delete 2>/dev/null || true

# Add test ignores to .openapi-generator-ignore if not already present
if ! grep -q "test/" ${CLIENT_DIR}/.openapi-generator-ignore; then
    echo "Adding test ignores to .openapi-generator-ignore..."
    echo "" >> ${CLIENT_DIR}/.openapi-generator-ignore
    echo "# Ignore test files and directories" >> ${CLIENT_DIR}/.openapi-generator-ignore
    echo "test/" >> ${CLIENT_DIR}/.openapi-generator-ignore
    echo "*_test.go" >> ${CLIENT_DIR}/.openapi-generator-ignore
fi

echo "Generated code has been fixed!"
echo "X-Compatibility-Date will now automatically use configuration value: $COMPAT_DATE"
echo "Test files and directories have been cleaned up and ignored for future generations"
