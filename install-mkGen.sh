#!/bin/bash

# Check if Go is installed
if ! command -v go &> /dev/null
then
    echo "Go could not be found. Please install Go and try again."
    exit
fi

# Define the repository URL and script names
REPO_URL="https://raw.githubusercontent.com/YourGitHubUsername/YourRepoName/main/"
SCRIPT_NAME="mkGen.go"
EXECUTABLE_NAME="mkGen"

# Download the Go script
curl -LO "${REPO_URL}${SCRIPT_NAME}"

# Compile the Go script into a binary
go build -o $EXECUTABLE_NAME $SCRIPT_NAME

# Move the binary to /usr/local/bin
sudo mv $EXECUTABLE_NAME /usr/local/bin/

# Cleanup the downloaded script
rm $SCRIPT_NAME

echo "Installation completed. You can now use 'mkGen' to generate microservice projects."
