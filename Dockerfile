FROM go:1.23.3 AS builder

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o main cmd/main.go

FROM go:1.23.3-alpine AS runner

WORKDIR /app
COPY --from=builder /build/main .
ENTRYPOINT ["/main"]
