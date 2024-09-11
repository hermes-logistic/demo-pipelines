#!/bin/bash

# Run tests and generate coverage profile
go test ./... -coverprofile=coverage.out

go tool cover -html=coverage.out -o coverage.html

# Extract the total coverage percentage
COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print substr($3, 1, length($3)-1)}')

# Check if coverage is below the threshold
THRESHOLD=10.0
if (( $(echo "$COVERAGE < $THRESHOLD" | bc -l) )); then
  echo "Test coverage ($COVERAGE%) is below threshold ($THRESHOLD%)"
  exit 1
else
  echo "Test coverage ($COVERAGE%) is sufficient."
fi
