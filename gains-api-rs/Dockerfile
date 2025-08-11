# Use a Rust base image for building
FROM rust:latest as build

# Set the working directory inside the container
WORKDIR /usr/src/rusty_api

# Copy the Rust project files
COPY . .

# Build the Rust application
RUN cargo build --release

# Start a new stage to create the final lightweight image
FROM fedora:38

# Set the working directory inside the container
WORKDIR /usr/src/rusty_api

# Copy the executable from the build stage to the final stage
COPY --from=build /usr/src/rusty_api/target/release/hoot_core_api .

EXPOSE 3030

# Command to run the application
CMD ["./hoot_core_api"]

