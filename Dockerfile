FROM golang:1.24 AS backend-builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o backend .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/backend /backend

EXPOSE 9000

CMD ["/backend"]
