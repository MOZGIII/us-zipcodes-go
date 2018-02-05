#!/bin/bash
set -euo pipefail
cd "$(dirname "${BASH_SOURCE[0]}")"
./update-sources.sh
./update-generated-code.sh
