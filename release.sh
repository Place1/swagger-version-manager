#!/bin/bash
set -eou pipefail

if ! [ -x "$(command -v github-release)" ]; then
  echo 'please install github-release using "go get github.com/aktau/github-release"'
fi

if [[ -z "$GITHUB_TOKEN" ]]; then
  echo 'please set the $GITHUB_TOKEN variable'
  exit 1
fi

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$DIR"

read -p 'GitHub Username: ' USERNAME
read -p 'Tag (e.g. v1.0.0): ' TAG

echo "building..."
make -j4

echo "creating release"
github-release release \
  --user "$USERNAME" \
  --repo swagger-version-manager \
  --tag "$TAG" \
  --name "$TAG"

echo "uploading artifacts"
ARTIFACTS=(
  "swagger-version-manager-linux-amd64"
  "swagger-version-manager-darwin-amd64"
  "swagger-version-manager-windows-amd64"
)
for ARTIFACT in "${ARTIFACTS[@]}"; do
  echo "uploading: $ARTIFACT"
  github-release upload \
    --user "$USERNAME" \
    --repo swagger-version-manager \
    --tag "$TAG" \
    --name "$ARTIFACT" \
    --file "./$ARTIFACT"
done

echo "done"
