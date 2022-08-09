FROM golang:1.18.3-alpine

ENV GIN_MODE=release

WORKDIR /go/src/
COPY . .
RUN GOOS=linux go build -ldflags="-s -w"
EXPOSE $PORT
CMD ["./tecnicos_service"]
