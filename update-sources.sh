#!/bin/bash
set -euo pipefail
cd "$(dirname "${BASH_SOURCE[0]}")"
SOURCESDIR="$(pwd)/sources"
URL="http://download.geonames.org/export/zip/US.zip"

TMPDIR="$(mktemp -d --suffix=.update-sources)"
# shellcheck disable=SC2064
trap "rm -rf $TMPDIR" EXIT
pushd "$TMPDIR"
wget "$URL" -O archive.zip
unzip archive.zip
rm archive.zip
cp "./US.txt" "$SOURCESDIR/zipcodes.txt"
cp "./readme.txt" "$SOURCESDIR/readme.txt"
