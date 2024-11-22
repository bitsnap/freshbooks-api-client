#!/usr/bin/env bash
set +eux

cd "$(dirname "$0")" || exit 1

echo "" > .tool-versions-updated

while IFS= read -r line; do
  name="$(echo "$line" | awk '{ print $1 }')"
  version="$(asdf list-all "$name" | \
    grep "^[0-9]" | \
    grep "[0-9a-z]-" -v | \
    grep "[0-9][a-z]" -v | \
    sort --version-sort | \
    tail -n1)"
  echo "$name $version"
  echo "$name $version" >> .tool-versions-updated
  asdf plugin add "$name" > /dev/null
done < .tool-versions

mv .tool-versions-updated .tool-versions