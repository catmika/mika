# -------- build stage --------
FROM golang:1.25.5-alpine AS builder

WORKDIR /app

# Install deps needed for builds
RUN apk add --no-cache git ca-certificates

COPY go.mod ./
COPY go.sum* ./
RUN go mod download

COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o server ./cmd/api

# -------- runtime stage --------
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/server /app/server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/server"]
