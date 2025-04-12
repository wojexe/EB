FROM golang:1.24-alpine
EXPOSE 8080

RUN apk add --update-cache
RUN apk add --update alpine-sdk

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -v -o compiled_app

CMD ["./compiled_app"]

