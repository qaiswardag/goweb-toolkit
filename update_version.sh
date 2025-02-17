# Read the version number from the VERSION file
VERSION=$(cat VERSION)

# Create a new tag with the desired version number
git tag $VERSION

# Push the tag to GitHub
git push origin $VERSION

# Create a new release on GitHub using the GitHub CLI
gh release create $VERSION --title "$VERSION" --notes "Release notes for $VERSION"

echo "Release version $VERSION created successfully."



# Check the version of the module in another Go project
echo "\n\nChecking the version of the GO module:"
go list -m github.com/qaiswardag/goweb-toolkit

# See the current tag
echo "\nChecking the Current tag the GO module::"
git describe --tags