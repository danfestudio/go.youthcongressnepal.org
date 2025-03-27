# Use Alpine 3.21.3 as the base image
FROM alpine:3.21.3

# Install dependencies for Go and bash (for setup.sh)
RUN apk add go

# Set the working directory inside the container
WORKDIR /app

# Copy the rest of the application code into the container
COPY . .

# Expose the port your application will use (adjust if needed)
EXPOSE 8001

CMD ["go", "run", "main.go"]
