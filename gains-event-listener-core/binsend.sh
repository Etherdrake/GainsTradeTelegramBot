#!/bin/bash

# Define variables for the remote server, private key file, and username
binary_file="HootCoreListener"
destination_path="/home/ec2-user/Hoot/HootCoreListener"
remote_server="ec2-user@54.75.102.178"
private_key="/home/cybrex/.ssh/hoot.pem"

# Use SCP with the private key file to securely copy the binary file to the remote server
scp -i $private_key $binary_file $remote_server:$destination_path

# Check the exit status of the SCP command
if [ $? -eq 0 ]; then
    echo "Binary file successfully sent to the remote server."
else
    echo "Error: Failed to send the binary file to the remote server."
fi
