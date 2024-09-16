#!/bin/bash

set -eo pipefail

BOLD_WHITE='\033[1;37m'
CYAN='\033[0;36m'
GREEN='\033[0;32m'
PURPLE='\033[0;35m'
BOLD_RED='\033[1;31m'
RED='\033[0;31m'
BOLD_YELLOW='\033[1;33m'
RESET='\033[0m' # No Color

function check_command_installed {
  if ! [ -x "$(command -v $1)" ]; then
    echo -e "${BOLD_RED}ERROR${RESET}: $1 is not installed." >&2
    exit 1
  fi
}

function print_and_execute() {
  echo "+ $@" >&2
  "$@"
}
