#!/bin/bash

# Load environment variables from .env file
while read line; do
  if echo "$line" | grep -F = &>/dev/null; then
    export "$line"
  fi
done < .env

# Check if required environment variables are set
if [ -z "$DOCKER_CONFIG_JSON" ] || [ -z "$QUAY_TOKEN" ] || [ -z "$MY_GITHUB_ORG" ] || [ -z "$GITHUB_TOKEN" ] || [ -z "$QUAY_E2E_ORGANIZATION" ]; then
  echo "Error: One or more required environment variables are not set."
  exit 1
fi

# Check if MONITORING_URL is set
if [ -z "$MONITORING_URL" ]; then
  echo "Monitoring instance is not set, running the test without pushing metrics"
fi

# Run go test command
go test -timeout 30m -test.v -test.run ^TestFeatures$
