#!/bin/bash

# Function to check if a command exists
command_exists() {
  command -v "$1" >/dev/null 2>&1
}

# Ensure required tools are installed
if ! command_exists git; then
  echo "Error: git is not installed."
  exit 1
fi

if ! command_exists gh; then
  echo "Error: GitHub CLI (gh) is not installed."
  exit 1
fi

# Read the version number from the file
VERSION=$(cat ./new_release/VERSION)
if [ -z "$VERSION" ]; then
  echo "Error: VERSION file is empty."
  exit 1
fi

# Check if the release notes file exists and is not empty
if [ ! -s ./new_release/RELEASE_NOTES.md ]; then
  echo "Error: RELEASE_NOTES.md file is missing or empty."
  exit 1
fi


# Create a new tag with the desired version number
git tag "$VERSION"
if [ $? -ne 0 ]; then
  echo "Error: Failed to create git tag."
  exit 1
fi

# Push the tag to GitHub
git push origin "$VERSION"
if [ $? -ne 0 ]; then
  echo "Error: Failed to push git tag to GitHub."
  exit 1
fi

# Create a GitHub release with release notes
gh release create "$VERSION" --title "$VERSION" --notes-file ./new_release/RELEASE_NOTES.md
if [ $? -eq 0 ]; then
  echo "Release version $VERSION created successfully."
else
  echo "Error: Failed to create GitHub release."
  exit 1
fi

# Check the current tag of the GO module
echo -e "\n\nChecking the current tag of the GO module:"
git describe --tags