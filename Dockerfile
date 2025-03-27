# Use Alpine 3.21.3 as the base image
FROM alpine:3.21.3

# Install dependencies for Go and bash (for setup.sh)
RUN apk update && \
    apk add --no-cache \
    bash \  
    curl

# Set the working directory inside the container
WORKDIR /app

# Copy the .devcontainer/setup.sh script into the container
COPY .devcontainer/setup.sh .devcontainer/setup.sh

# Make sure the script is executable
RUN chmod +x .devcontainer/setup.sh

# Copy the rest of the application code into the container
COPY . .

# Run the setup.sh script
RUN ./devcontainer/setup.sh

# Expose the port your application will use (adjust if needed)
EXPOSE 8080

# Command to run the Go application
CMD ["go", "run", "main.go"]
