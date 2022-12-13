#!/bin/bash
set -e

is_comment() {
	case "$1" in
	\#*)
		log_verbose "Skip: $1"
		return 0
		;;
	esac
	return 1
}

is_blank() {
	case "$1" in
	'')
		log_verbose "Skip: $1"
		return 0
		;;
	esac
	return 1
}

while IFS='=' read -r key temp || [ -n "$key" ]; do
  if is_comment "$key"; then
    continue
  fi

  if is_blank "$key"; then
    continue
  fi

  # Strip any existing quotes
  temp="${temp%[\'\"]}"
  temp="${temp#[\'\"]}"
  # Add new double quotes for interpolation
  value=$(eval echo "\"$temp\"")

  eval export "$key='$value'"

done < <(sed -nr 's/([A-Z_]+): (.*)/\1=\2/ p' .cloudbuild/config/${BRANCH_NAME}.env.yaml)