FROM golang:1.22-alpine AS builder

WORKDIR /build
COPY . .
RUN go mod download 
RUN go build -o ./receipt-processor

FROM gcr.io/distroless/base-debian12 

WORKDIR /app
COPY --from=builder /build/receipt-processor ./receipt-processor
CMD ["/app/receipt-processor"]