FROM golang:1.23.3 AS builder
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o main cmd/main.go

FROM scratch
ENV GIN_MODE=release

WORKDIR /bin
COPY --from=builder /build/main .

EXPOSE 9000
ENTRYPOINT ["/bin/main"]
