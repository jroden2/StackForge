#!/usr/bin/env bash

set -e

echo "Starting Mac dev setup"

if ! command -v brew >/dev/null 2>&1; then
	echo "Installing Homebrew"
	/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

	if [[ -d "/opt/homebrew/bin" ]]; then
		eval "$(/opt/homebrew/bin/brew shellen)"
	elif [[ -d "/usr/local/bin" ]]; then
		eval "$(/usr/local/bin/brew shellenv)"
	fi
else
	echo "Homebrew installed"
fi

echo "Updating..."
brew update

echo "Installing deps..."
sleep 3

echo "Installing docker"
brew install docker

echo "Installing maven"
brew install maven

echo "Installing JDK 21 (Eclipse Temurin)"
brew install --cask temurin@21

echo ""
echo "Installed versions:"
docker --version || true
mvn -version || true
java -version || true
