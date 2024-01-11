FROM golang:alpine

RUN apk add build-base

WORKDIR /app

COPY . .

RUN go build -o win cmd/api/*.go

EXPOSE 1323

CMD ["./win"]
