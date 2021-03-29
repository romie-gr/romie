#!/usr/bin/env bash
# Windows users can run this using Git for Windows (https://git-scm.com/download/win)

# Check dependencies
# ------------------ #
if ! curl --version &> /dev/null; then
	echo "Please install 'curl' first. Exiting ..."
	exit 1
fi

if ! go version &> /dev/null; then
	echo "Please install 'go' first. Exiting ..."
	exit 1
fi

if [[ -z "${GOPATH}" ]]; then
  echo "The env variable 'GOPATH' is not defined. Exiting ..."
	exit 1
fi
# ------------------ #

# Fetch the latest version string from upstream
URL="https://github.com/golangci/golangci-lint/releases/latest"
LATEST_VERSION=$(curl -s "$URL" | awk -F 'tag/' '{print $2}' | awk -F '"' '{ print $1 }')

if golangci-lint version --format short &> /dev/null; then
	echo "golangci-lint is already installed."
	CURRENT_VERSION="v$(golangci-lint version --format short)"
	if [ "$LATEST_VERSION" == "$CURRENT_VERSION" ]; then
		echo "golangci-lint is already updated to the latest $LATEST_VERSION"
		exit 0
	else
		echo "golangci-lint is at $CURRENT_VERSION, but the latest version is $LATEST_VERSION. Updating now ..."
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(go env GOPATH)/bin" "$LATEST_VERSION"
	fi
else
	echo "golangci-lint is not installed. Installing now ..."
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(go env GOPATH)/bin" "$LATEST_VERSION"
fi

# Verify
if golangci-lint version --format short &> /dev/null; then
	echo "golangci-lint is already installed."
	CURRENT_VERSION="$(golangci-lint version --format short)"
	if [ "$LATEST_VERSION" == "$CURRENT_VERSION" ]; then
		echo "Success!"
		exit 0
	else
		echo "Update has failed. You can still run it, but your results mights differ from the ones from CI."
		exit 1
	fi
else
	echo "golangci-lint is not installed. Exiting now ..."
	exit 1
fi
