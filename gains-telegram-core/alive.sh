#!/bin/bash

# Function to start the Go binary
start_go_binary() {
    ./HootCore  # Replace with the path to your Go binary
}

# Infinite loop to run and restart the Go binary
while true; do
    echo "Starting Go binary..."
    start_go_binary

    # Check the exit status of the Go binary
    if [ $? -eq 0 ]; then
        echo "Go binary exited normally. Exiting script."
        exit 0
    else
        echo "Go binary exited with error. Restarting..."
        sleep 1  # Wait before restarting
    fi
done
