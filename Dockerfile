FROM golang:1.25-alpine AS builder
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -o watchducker .

FROM alpine:latest
RUN apk add --no-cache tzdata
COPY --from=builder /app/watchducker /usr/local/bin/

CMD ["watchducker"]
