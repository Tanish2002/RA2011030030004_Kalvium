FROM golang:1.19-alpine

ENV CGO_ENABLED=1

RUN apk add --no-cache \
    # Important: required for go-sqlite3
    gcc \
    # Required for Alpine
    musl-dev

WORKDIR /app

# Build module cache
COPY go.mod .
COPY go.sum .
RUN go mod download


COPY . .

# Build the application
RUN go build -o /kalvium-expression-evaluator

EXPOSE 8080

CMD [ "/kalvium-expression-evaluator" ]
