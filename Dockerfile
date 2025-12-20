# syntax=docker/dockerfile:1

################################################################################
# Build stage
ARG GO_VERSION=1.23.6
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine AS build
WORKDIR /src

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

ARG TARGETARCH

# Build the application
RUN CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/bot ./cmd/app/main.go

################################################################################
# Final stage
FROM alpine:latest AS final

WORKDIR /app

# Install runtime dependencies
RUN apk --no-cache add \
    ca-certificates \
    tzdata

# Create a non-privileged user for security
ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser

# Copy the executable from build stage
COPY --from=build /bin/bot ./bot

# Copy locale files
COPY internal/repository/locales/ ./internal/repository/locales/

# Set ownership to appuser
RUN chown -R appuser:appuser /app

USER appuser

# Run the bot
ENTRYPOINT ["./bot"]