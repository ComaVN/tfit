FROM golang:1.23.3-alpine AS builder
# This Dockerfile builds and runs the tfit executable.
# Usage:
#   docker build -t tfit .
#   echo 'SGVsbG8sIFdvcmxkIQo=' | docker run -i --rm tfit

WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go build -o ./build/tfit cmd/tfit/main.go


FROM scratch

COPY --from=builder /workspace/build/tfit /tfit
CMD ["/tfit"]
