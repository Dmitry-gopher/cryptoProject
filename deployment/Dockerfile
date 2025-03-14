FROM golang:1.23.0-alpine3.10 AS builder

WORKDIR /usr/local/src

RUN apk add --no-cache add bash git gcc gettext musl-dev

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY app ./
RUN go build -o ./bin/app main.go

FROM alpine:3.10 as runner

COPY --from=builder /usr/local/src/bin/app /
COPY configs/config.yaml /config.yaml

CMD ["/app"]