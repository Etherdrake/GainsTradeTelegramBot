#!/bin/bash

# Prompt the user to enter the binary file name
read -p "Enter the name of the binary file: " binary_file

# Check if the entered file exists
if [ ! -f "$binary_file" ]; then
    echo "Error: The specified binary file does not exist."
    exit 1
fi

# Prompt the user to enter the destination path on the remote server
read -p "Enter the destination path on the remote server: " destination_path

# Define variables for the remote server, private key file, and username
remote_server="server-name"
private_key="path/to/hoot.pem"

# Use SCP with the private key file to securely copy the binary file to the remote server
scp -i $private_key $binary_file $remote_server:$destination_path

# Check the exit status of the SCP command
if [ $? -eq 0 ]; then
    echo "Binary file successfully sent to the remote server."
else
    echo "Error: Failed to send the binary file to the remote server."
fi
