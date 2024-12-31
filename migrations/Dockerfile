FROM golang:1.21

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

WORKDIR /app

ENTRYPOINT ["goose"]
CMD ["help"]
