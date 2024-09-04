#!/bin/bash

# Export database connection details
if [ -f .env ]; then
    # Read .env file
    export $(grep -v '^#' .env | xargs)
else
    echo ".env file not found"
    exit 1
fi


# Export any other project-specific variables
export PROJECT_ROOT=$(pwd)

# Activate Go module for this directory
export GO111MODULE=on

echo "Environment setup complete for Todo API project"

# Start a new shell with the updated environment
exec $SHELL