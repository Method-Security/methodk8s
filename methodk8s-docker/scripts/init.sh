#!/bin/bash

# First argument is the container-specific script
CONTAINER_SCRIPT="$1"

# Shift the arguments so $@ no longer includes the first argument (the script path)
# This leaves only the arguments intended for the container-specific script
shift

# Ensure the container-specific script is provided and exists
if [ -z "$CONTAINER_SCRIPT" ] || [ ! -f "$CONTAINER_SCRIPT" ]; then
  echo "Error: Container-specific script not found."
  exit 1
fi

# Set the path for temporary output and error files
TEMP_DATA_OUTPUT="var/data/raw_output"
STDOUT_OUTPUT="var/data/stdout.txt"
STDERR_OUTPUT="var/data/error_message.txt"
ENCODED_CONTENT="var/data/encoded_content.txt"
OUTPUT_FILE="/mnt/output/output.json"


# Cleanup previous runs
rm -f "$TEMP_DATA_OUTPUT" "$STDOUT_OUTPUT" "$STDERR_OUTPUT" "$ENCODED_CONTENT"

# Capture the start time
started_at=$(date +%s)

# Execute the container-specific script and redirect its output and error
bash "$CONTAINER_SCRIPT" "$@" > "$STDOUT_OUTPUT" 2> "$STDERR_OUTPUT"
status=$?

# Capture the end time
completed_at=$(date +%s)

# Encode the output in Base64
base64 -w 0 "$TEMP_DATA_OUTPUT" > "$ENCODED_CONTENT"


# Check for error messages
if [ -s "$STDERR_OUTPUT" ]; then
    error_message=$(<"$STDERR_OUTPUT")
else
    error_message=null
fi

# Generate JSON
if [ -s "$STDERR_OUTPUT" ]; then
    # If there's an error message, include it in the JSON
    json=$(jq -n \
                --rawfile content "$ENCODED_CONTENT" \
                --argjson started_at "$started_at" \
                --argjson completed_at "$completed_at" \
                --arg status "$status" \
                --rawfile error_message "$STDERR_OUTPUT" \
                '{
                  content: $content,
                  started_at: $started_at,
                  completed_at: $completed_at,
                  status: $status,
                  error_message: $error_message
                 }')
else
    # No error message, exclude it from the JSON
    json=$(jq -n \
                --rawfile content "$ENCODED_CONTENT" \
                --argjson started_at "$started_at" \
                --argjson completed_at "$completed_at" \
                --arg status "$status" \
                '{
                  content: $content,
                  started_at: $started_at,
                  completed_at: $completed_at,
                  status: $status
                 }')
fi

# Save JSON to file
echo "$json" > "$OUTPUT_FILE"
