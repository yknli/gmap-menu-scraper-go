#!/bin/bash

# Ensure necessary tools are installed
command -v wget >/dev/null 2>&1 || { echo >&2 "wget is required but it's not installed. Aborting."; exit 1; }
command -v unzip >/dev/null 2>&1 || { echo >&2 "unzip is required but it's not installed. Aborting."; exit 1; }

# download selenium
mkdir -p selenium
wget https://selenium-release.storage.googleapis.com/3.141/selenium-server-standalone-3.141.59.jar -O selenium/selenium-server-standalone.jar

# Note: If you are using macOS or Windows, you need to change the download links below to download the correct binaries for your operating system.
# https://googlechromelabs.github.io/chrome-for-testing/

# download chrome and unzip
wget https://storage.googleapis.com/chrome-for-testing-public/131.0.6778.204/linux64/chrome-linux64.zip
unzip chrome-linux64.zip

# move chrome binary into chrome directory
mv chrome-linux64 chrome

# download chromedriver and unzip
wget https://storage.googleapis.com/chrome-for-testing-public/131.0.6778.204/linux64/chromedriver-linux64.zip
unzip chromedriver-linux64.zip

# move chromedriver binary into chrome directory
mv chromedriver-linux64/chromedriver chrome

# cleanup
rm -rf chromedriver-linux64 chromedriver-linux64.zip chrome-linux64.zip