FROM golang:1.18-bullseye as builder

LABEL maintainer="Ahmed Belhoula <ahmd.belhoula@gmail.com>"

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.mod go.sum ./
RUN go mod download

# Copy local code to the container image.
COPY . .

# Build the binary.
RUN go build -v -o rest-api

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:bullseye-slim
WORKDIR /
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/* && \
    mkdir -p /app/logs

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/rest-api .
COPY --from=builder /app/config.dev.yml .


# Expose port 8080 to the outside world
EXPOSE 8080

# Run the web service on container startup.
CMD ["./rest-api"]
