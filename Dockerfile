FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY . .

RUN go mod tidy

RUN go build -o backend .

FROM alpine:3.18

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/backend /backend

EXPOSE 9000

CMD ["/backend"]
