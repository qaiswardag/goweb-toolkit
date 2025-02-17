# Read the version number from the file
pwd
VERSION=$(cat ./new_release/update_current_version_number)

# Check if the file exists and is not empty
if [ ! -s ./new_release/RELEASE_NOTES.md ]; then
  echo "Error: File is missing or empty."
  exit 1
fi

# Create a new tag with the desired version number
git tag $VERSION

# Push the tag to GitHub
git push origin $VERSION

# Create a GitHub release with release notes
gh release create $VERSION --title "$VERSION" --notes-file ./new_release/RELEASE_NOTES.md

if [ $? -eq 0 ]; then
  echo "Release version $VERSION created successfully."
else
  echo "Error: Failed to create GitHub release."
  exit 1
fi

# Check the current tag of the GO module
echo "\n\nChecking the Current tag of the GO module::"
git describe --tags