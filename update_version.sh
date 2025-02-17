# Read the version number from the VERSION file
VERSION=$(cat VERSION)

# Check if release_notes.md exists and is not empty
if [ ! -s release_notes.md ]; then
  echo "Error: release_notes.md file is missing or empty."
  exit 1
fi

# Create a new tag with the desired version number
git tag $VERSION

# Push the tag to GitHub
git push origin $VERSION

# Create a GitHub release with release notes
gh release create $VERSION --title "$VERSION" --notes-file release_notes.md

if [ $? -eq 0 ]; then
  echo "Release version $VERSION created successfully."
else
  echo "Error: Failed to create GitHub release."
  exit 1
fi

# Check the current tag of the GO module
echo "\n\nChecking the Current tag of the GO module::"
git describe --tags