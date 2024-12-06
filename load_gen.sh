#!/bin/bash

# Set the duration to 2 hours (in seconds)
DURATION=$((2 * 60 * 60))

# Get the start time
START_TIME=$(date +%s)

while true; do
  # Run your command here
  grpcurl --proto health.proto server:port health.HealthService/CheckHealth
  grpcurl --proto health.proto server:port health.HealthService/CheckNonHealth
  grpcurl --proto health.proto server:port list

  # Check the current time
  CURRENT_TIME=$(date +%s)
  ELAPSED_TIME=$((CURRENT_TIME - START_TIME))

  # Exit the loop if the elapsed time exceeds the duration
  if [ "$ELAPSED_TIME" -ge "$DURATION" ]; then
    echo "Task completed after 2 hours."
    break
  fi

  # Optional: Add a sleep time between iterations (e.g., 1 second)
  sleep 1
done