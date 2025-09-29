#!/bin/bash

# Script to fix the generated ESI client code to automatically handle X-Compatibility-Date
# This should be run after each code generation

# Configuration
CLIENT_DIR="esi"

echo "Fixing generated ESI client code..."

# Find all API files in the client directory
find ${CLIENT_DIR}/ -name "api_*.go" -type f | while read -r file; do
    echo "Processing $file..."
    
    # Replace the nil check with automatic default assignment using local variable
    sed -i.bak 's/if r\.xCompatibilityDate == nil {/if r.xCompatibilityDate == nil {\
		defaultCompatibilityDate := "2020-01-01"\
		r.xCompatibilityDate = \&defaultCompatibilityDate\
	}\
	\/\/ Removed original error check - now using default value\
	if false {/g' "$file"
    
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
echo "X-Compatibility-Date will now automatically default to '2020-01-01'"
echo "Test files and directories have been cleaned up and ignored for future generations"
