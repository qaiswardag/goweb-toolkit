# Read the version number from the VERSION file
VERSION=$(cat VERSION)

# Create a new tag with the desired version number
git tag $VERSION

# Push the tag to GitHub
git push origin $VERSION

# Create a new release on GitHub using the GitHub CLI
gh release create $VERSION --title "$VERSION" --notes --notes-file RELEASE_NOTES.md

echo "Release version $VERSION created successfully."


# Check the current tag of the GO module
echo "\n\nChecking the Current tag of the GO module::"
git describe --tags