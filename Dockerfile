# -------- Build Stage --------
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Install git (needed for some Go modules)
RUN apk add --no-cache git

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /money-manager ./cmd/api


# -------- Runtime Stage --------
FROM alpine:3.19

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Copy binary
COPY --from=builder /money-manager /app/money-manager

# Copy migrations
COPY migrations/ /app/migrations/

# Copy entrypoint script
COPY scripts/entrypoint.sh /app/scripts/entrypoint.sh

# Make script executable
RUN chmod +x /app/scripts/entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["/app/scripts/entrypoint.sh"]
CMD ["/app/money-manager"]
