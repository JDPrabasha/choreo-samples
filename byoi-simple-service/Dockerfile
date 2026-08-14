# This Dockerfile creates an image that contains the simple-service executable.
# The image is based on a small Alpine Linux image.
# The image uses a non-root user with a known UID/GID to run the container.
# The image has a single entrypoint, the simple-service executable.
FROM --platform=$BUILDPLATFORM golang:1.19-alpine AS builder

# Set the working directory to /app
WORKDIR /app

# Copy the Go source file into the container
COPY . .

# Cross-compile for the target platform natively (avoids running the Go
# compiler under QEMU emulation, which crashes on some host/target combos).
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o simple-service .

# Use a small Alpine Linux image as the base image for the final container
FROM alpine:latest

# Set the working directory to /app
WORKDIR /app

# Copy the executable from the builder container
COPY --from=builder /app/simple-service .

# Create a user with a known UID/GID within range 10000-20000.
# This is required by Choreo to run the container as a non-root user.
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid 10014 \
    "choreo"
# Use the above created unprivileged user
USER 10014

# Set the entrypoint to the executable
ENTRYPOINT ["./simple-service"]
